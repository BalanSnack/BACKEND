package like

import (
	"context"
	"gorm.io/gorm"

	"BACKEND/internals/pkg/agg"
)

var _ GameRepository = (*gameRepository)(nil)
var _ CommentRepository = (*commentRepository)(nil)

type gameRepository struct {
	db *gorm.DB
}

type commentRepository struct {
	db *gorm.DB
}

func (r *gameRepository) Create(ctx context.Context, like *agg.GameLike) (*agg.GameLike, error) {
	return nil, nil
}

func (r *gameRepository) Get(ctx context.Context, id int) (*agg.GameLike, error) {
	return nil, nil
}

func (r *gameRepository) Update(ctx context.Context, like *agg.GameLike) (*agg.GameLike, error) {
	return nil, nil
}

func (r *gameRepository) Delete(ctx context.Context, id int) error {
	return nil
}

func (r *commentRepository) Create(ctx context.Context, like *agg.CommentLike) (*agg.CommentLike, error) {
	return nil, nil
}

func (r *commentRepository) Get(ctx context.Context, id int) (*agg.CommentLike, error) {
	return nil, nil
}

func (r *commentRepository) Update(ctx context.Context, like *agg.CommentLike) (*agg.CommentLike, error) {
	return nil, nil
}

func (r *commentRepository) Delete(ctx context.Context, id int) error {
	return nil
}

func NewGameRepository(db *gorm.DB) GameRepository {
	return &gameRepository{
		db: db,
	}
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}
