package comment

import (
	"context"

	"BACKEND/internals/pkg/agg"
)

type Repository interface {
	Create(ctx context.Context, comment *agg.Comment) (*agg.Comment, error)
	Get(ctx context.Context, id int) (*agg.Comment, error)
	Update(ctx context.Context, comment *agg.Comment) (*agg.Comment, error)
	Delete(ctx context.Context, id int) error
}
