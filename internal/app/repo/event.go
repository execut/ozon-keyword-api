package repo

import (
    "errors"
    "github.com/Masterminds/squirrel"
    "github.com/execut/ozon-keyword-api/internal/model"
    "github.com/jmoiron/sqlx"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type EventRepo interface {
    Lock(n uint64) ([]model.KeywordEvent, error)
    Unlock(eventIDs []uint64) error

    Add(event []model.KeywordEvent) error
    Remove(eventIDs []uint64) error
}

type EventDbRepo struct {
    db *sqlx.DB
}

func (r *EventDbRepo) Lock(n uint64) ([]model.KeywordEvent, error) {
    tx, err := r.db.Beginx()
    if err != nil {
        return nil, err
    }
    rows, err := psql.Select("id", "keyword_id", "type").
        From("keyword_events").
        Where("status is null").
        OrderBy("id").
        Limit(n).
        RunWith(r.db).Query()
    if err != nil {
        return nil, err
    }

    var events []model.KeywordEvent
    var ids = []uint64{}
    for rows.Next() {
        event := model.KeywordEvent{}
        var keywordId uint64
        err := rows.Scan(&event.ID, &keywordId, &event.Type)
        if err != nil {
            return nil, err
        }

        event.Entity = &model.Keyword{ID: keywordId}
        events = append(events, event)
        ids = append(ids, event.ID)
    }

    query := psql.
        Update("keyword_events").
        Set("status", model.Deferred).
        Where(squirrel.Eq{"id": ids}).
        RunWith(r.db)
    _, err = query.
        Exec()

    if err != nil {
        return nil, err
    }

    tx.Commit()

    return events, nil
}

func (r *EventDbRepo) Unlock(eventIDs []uint64) error {
    query := psql.
        Update("keyword_events").
        Set("status", nil).
        Where(squirrel.Eq{"id": eventIDs}).
        RunWith(r.db)
    _, err := query.
        Exec()
    if err != nil {
        return err
    }

    return nil
}

func (r *EventDbRepo) Add(events []model.KeywordEvent) error {
    tx, err := r.db.Beginx()
    if err != nil {
        return err
    }

    for _, event := range events {
        query := psql.
            Insert("keyword_events").
            Columns("keyword_id", "type").
            Suffix("RETURNING id").
            RunWith(r.db).
            Values(event.Entity.ID, event.Type)
        err = query.
            Scan(&event.ID)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    tx.Commit()

    return nil
}

func (r *EventDbRepo) Remove(eventIDs []uint64) error {
    query := psql.
        Update("keyword_events").
        Set("status", model.Processed).
        Where(squirrel.Eq{"id": eventIDs}).
        RunWith(r.db)
    _, err := query.
        Exec()
    if err != nil {
        return err
    }

    return nil
}

var ErrEventNotFound error = errors.New("Event not found")

var ErrNoMoreEvents error = errors.New("No more events for lock")

func NewEventRepo(db *sqlx.DB) EventRepo {
    return &EventDbRepo{
        db: db,
    }
}
