package repo

import (
    "context"

    "github.com/jmoiron/sqlx"

    "github.com/execut/ozon-keyword-api/internal/model"
)

// Repo is DAO for Keyword
type Repo interface {
    DescribeKeyword(ctx context.Context, ozonID uint64) (*model.Keyword, error)
}

type repo struct {
    db        *sqlx.DB
    batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
    return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeKeyword(ctx context.Context, ozonID uint64) (*model.Keyword, error) {
    return nil, nil
}
