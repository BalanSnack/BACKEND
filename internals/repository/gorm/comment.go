package gorm

import "gorm.io/gorm"

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

// Create 댓글을 생성한다.
func (r *CommentRepo) Create(avatarID uint, parentID uint, gameID uint, content string) (Comment, error) {
	comment := Comment{
		AvatarID: avatarID,
		ParentID: parentID,
		GameID:   gameID,
		Content:  content,
	}

	err := r.db.Create(&comment).Error

	return comment, err
}

// Update 댓글 내용을 수정한다.
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

// Delete 댓글을 삭제한다.
func (r *CommentRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Comment{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// GetByID 댓글 정보를 조회한다.
func (r *CommentRepo) GetByID(id uint) (Comment, error) {
	var comment Comment

	err := r.db.First(&comment, id).Error

	return comment, err
}

// GetAllByGameID 게임에 등록된 댓글들을 모두 조회한다.
func (r *CommentRepo) GetAllByGameID(gameID uint) ([]Comment, error) {
	var comments []Comment

	err := r.db.Where("game_id = ?", gameID).Find(&comments).Error

	return comments, err
}

// UpdateVoteUp 좋아요를 기록한다.
func (r *CommentRepo) UpdateVoteUp(id uint) (affected int64, err error) {
	var comment Comment
	comment.ID = id

	tx := r.db.Model(&comment).UpdateColumn("vote", gorm.Expr("vote + ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// UpdateVoteDown 싫어요를 기록한다.
func (r *CommentRepo) UpdateVoteDown(id uint) (affected int64, err error) {
	var comment Comment
	comment.ID = id

	tx := r.db.Model(&comment).UpdateColumn("vote", gorm.Expr("vote - ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}
