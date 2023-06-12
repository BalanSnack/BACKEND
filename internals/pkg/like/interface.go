package like

import (
	"context"

	"github.com/BalanSnack/BACKEND/internals/pkg/agg"
)

type GameRepository interface {
	Create(ctx context.Context, like *agg.GameLike) (*agg.GameLike, error)
	Get(ctx context.Context, id int) (*agg.GameLike, error)
	Update(ctx context.Context, like *agg.GameLike) (*agg.GameLike, error)
	Delete(ctx context.Context, id int) error
}

type CommentRepository interface {
	Create(ctx context.Context, like *agg.CommentLike) (*agg.CommentLike, error)
	Get(ctx context.Context, id int) (*agg.CommentLike, error)
	Update(ctx context.Context, like *agg.CommentLike) (*agg.CommentLike, error)
	Delete(ctx context.Context, id int) error
}
