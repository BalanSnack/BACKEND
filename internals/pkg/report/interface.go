package report

import (
	"context"

	"github.com/BalanSnack/BACKEND/internals/pkg/agg"
)

type GameRepository interface {
	Create(ctx context.Context, report *agg.GameReport) (*agg.GameReport, error)
	Get(ctx context.Context, id int) (*agg.GameReport, error)
	Update(ctx context.Context, report *agg.GameReport) (*agg.GameReport, error)
	Delete(ctx context.Context, id int) error
}

type CommentRepository interface {
	Create(ctx context.Context, comment *agg.CommentReport) (*agg.CommentReport, error)
	Get(ctx context.Context, id int) (*agg.CommentReport, error)
	Update(ctx context.Context, comment *agg.CommentReport) (*agg.CommentReport, error)
	Delete(ctx context.Context, id int) error
}
