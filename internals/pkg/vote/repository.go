package vote

import (
	"context"
	"gorm.io/gorm"

	"BACKEND/internals/pkg/agg"
)

var _ Repository = (*repository)(nil)

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(ctx context.Context, vote *agg.Vote) (*agg.Vote, error) {
	return nil, nil
}

func (r *repository) Get(ctx context.Context, id int) (*agg.Vote, error) {
	return nil, nil
}

func (r *repository) Update(ctx context.Context, vote *agg.Vote) (*agg.Vote, error) {
	return nil, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	return nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
