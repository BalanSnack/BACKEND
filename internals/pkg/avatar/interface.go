package avatar

import (
	"context"

	"github.com/BalanSnack/BACKEND/internals/pkg/agg"
)

type Repository interface {
	Create(ctx context.Context, avatar *agg.Avatar) (*agg.Avatar, error)
	Get(ctx context.Context, id int) (*agg.Avatar, error)
	Update(ctx context.Context, avatar *agg.Avatar) (*agg.Avatar, error)
	Delete(ctx context.Context, id int) error
}
