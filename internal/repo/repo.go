package repo

import (
    "context"
    "errors"
    "github.com/Masterminds/squirrel"

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
    db        *sqlx.DB
    batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
    return &repo{db: db, batchSize: batchSize}
}

func (r *repo) Get(ctx context.Context, keywordID uint64) (*model.Keyword, error) {
    row := squirrel.
        Select("id, name").
        From("keywords").
        Where("id=$1", keywordID).
        Where("!removed").
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
        Delete("keywords").
        Where("id=$1", keywordID).
        Where("!removed").
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
    result, _ := squirrel.Insert("keywords").
        Columns("name").
        Values(keyword.Name).
        Suffix(" RETURNS id").
        RunWith(r.db).
        Exec()
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    keyword.ID = uint64(id)

    return keyword.ID, nil
}

func (r *repo) List(ctx context.Context, limit uint64, cursor uint64) ([]model.Keyword, error) {
    query := squirrel.Select("id", "name").
        From("keywords").
        Where("!removed").
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
        list.Scan(&keyword.ID, &keyword.Name)
        result = append(result, keyword)
    }

    return result, nil
}
