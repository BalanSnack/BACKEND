package mysql

import "gorm.io/gorm"

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

func (r *CommentRepo) Create(avatarID uint, parentID uint, content string) (Comment, error) {
	comment := Comment{
		AvatarID: avatarID,
		ParentID: parentID,
		Content:  content,
	}

	err := r.db.Create(&comment).Error

	return comment, err
}

func (r *CommentRepo) Update(id uint, content string) (affected int64, err error) {
	var comment Comment
	comment.ID = id

	tx := r.db.Model(&comment).Update("content", content)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

func (r *CommentRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Comment{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

func (r *CommentRepo) GetAllByGameID(gameID uint) ([]Comment, error) {
	var comments []Comment

	err := r.db.Where("game_id = ?", gameID).Find(&comments).Error

	return comments, err
}

func (r *CommentRepo) UpdateVoteUp(id uint) (affected int64, err error) {
	var comment Comment
	comment.ID = id

	tx := r.db.Model(comment).UpdateColumn("vote", gorm.Expr("vote + ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

func (r *CommentRepo) UpdateVoteDown(id uint) (affected int64, err error) {
	var comment Comment
	comment.ID = id

	tx := r.db.Model(comment).UpdateColumn("vote", gorm.Expr("vote - ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}
