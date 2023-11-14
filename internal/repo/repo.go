package repo

import (
    "context"
    "errors"
    "github.com/Masterminds/squirrel"
    "time"

    "github.com/jmoiron/sqlx"

    "github.com/execut/ozon-keyword-api/internal/model"
)

var ErrKeywordNotFound = errors.New("keyword not found")

// Repo is DAO for Keyword
type Repo interface {
    Get(ctx context.Context, keywordID uint64) (*model.Keyword, error)
    Add(ctx context.Context, keyword *model.Keyword) (uint64, error)
    List(ctx context.Context, limit uint64, cursor uint64) ([]model.Keyword, error)
    Remove(ctx context.Context, keywordID uint64) error
}

type repo struct {
    db *sqlx.DB
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB) Repo {
    return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, keywordID uint64) (*model.Keyword, error) {
    row := squirrel.
        Select("id, name").
        From("keywords").
        Where("id=$1", keywordID).
        Where("NOT removed").
        RunWith(r.db).
        QueryRowContext(ctx)
    keyword := &model.Keyword{}
    err := row.Scan(&keyword.ID, &keyword.Name)
    if err != nil {
        return nil, err
    }

    return keyword, nil
}

func (r *repo) Remove(ctx context.Context, keywordID uint64) error {
    result, err := squirrel.
        Update("keywords").
        Set("removed", true).
        Where(squirrel.Eq{"id": keywordID}).
        Where("NOT removed").
        PlaceholderFormat(squirrel.Dollar).
        RunWith(r.db).
        ExecContext(ctx)

    if err != nil {
        return err
    }

    affected, _ := result.RowsAffected()
    if affected != int64(1) {
        return ErrKeywordNotFound
    }

    return nil
}

func (r *repo) Add(ctx context.Context, keyword *model.Keyword) (uint64, error) {
    err := squirrel.Insert("keywords").
        Columns("name", "removed", "created").
        Values(keyword.Name, false, time.Now()).
        PlaceholderFormat(squirrel.Dollar).
        Suffix("RETURNING id").
        RunWith(r.db).
        Scan(&keyword.ID)
    if err != nil {
        return 0, err
    }

    return keyword.ID, nil
}

func (r *repo) List(ctx context.Context, limit uint64, cursor uint64) ([]model.Keyword, error) {
    query := squirrel.Select("id", "name").
        From("keywords").
        Where("NOT removed").
        Limit(limit).
        Offset(cursor).
        RunWith(r.db)
    list, err := query.
        QueryContext(ctx)
    if err != nil {
        return nil, err
    }
    var result = []model.Keyword{}
    for list.Next() {
        keyword := model.Keyword{}
        err := list.Scan(&keyword.ID, &keyword.Name)
        if err != nil {
            return nil, err
        }
        result = append(result, keyword)
    }

    return result, nil
}
