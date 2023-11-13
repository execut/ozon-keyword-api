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
    //Add(keyword *model.Keyword) (uint64, error)
    //List(limit uint64, cursor uint64) ([]model.Keyword, error)
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
        Where(squirrel.Eq{"id": keywordID}).
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
