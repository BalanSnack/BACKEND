package game

import (
	"context"

	"BACKEND/internals/pkg/agg"
)

type Repository interface {
	Create(ctx context.Context, avatar *agg.Game) (*agg.Game, error)
	Get(ctx context.Context, id int) (*agg.Game, error)
	Update(ctx context.Context, avatar *agg.Game) (*agg.Game, error)
	Delete(ctx context.Context, id int) error
}
