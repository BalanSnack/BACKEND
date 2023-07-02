package vote

import (
	"BACKEND/internals/pkg/agg"
	"context"
)

type Repository interface {
	Create(ctx context.Context, vote *agg.Vote) (*agg.Vote, error)
	Get(ctx context.Context, id int) (*agg.Vote, error)
	Update(ctx context.Context, vote *agg.Vote) (*agg.Vote, error)
	Delete(ctx context.Context, id int) error
}
