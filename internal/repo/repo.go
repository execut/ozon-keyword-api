package repo

import (
	"context"
	"github.com/Masterminds/squirrel"

	"github.com/jmoiron/sqlx"

	"github.com/execut/ozon-keyword-api/internal/model"
)

// Repo is DAO for Keyword
type Repo interface {
	GetKeyword(ctx context.Context, keywordID uint64) (*model.Keyword, error)
	//Add(keyword *model.Keyword) (uint64, error)
	//List(limit uint64, cursor uint64) ([]model.Keyword, error)
	//Remove(keywordID uint64) (bool, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) GetKeyword(ctx context.Context, keywordID uint64) (*model.Keyword, error) {
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
