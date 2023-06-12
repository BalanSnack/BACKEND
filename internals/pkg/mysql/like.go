package mysql

import (
	"database/sql"
	"fmt"
	"github.com/BalanSnack/BACKEND/internals/pkg"
)

type LikeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) *LikeRepository {
	return &LikeRepository{
		db: db,
	}
}

// CreateLikeComment 댓글 좋아요 생성
func (r *LikeRepository) CreateLikeComment(v *pkg.Like) error {
	stmt, err := r.db.Prepare("INSERT INTO likes(avatar_id, comment_id) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare create statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(v.AvatarID, v.CommentID)
	if err != nil {
		return fmt.Errorf("failed to execute create statement: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %v", err)
	}

	v.ID = int(id)
	return nil
}

// CreateLikeGame 게임 좋아요 생성, commentID 무시
func (r *LikeRepository) CreateLikeGame(v *pkg.Like) error {
	stmt, err := r.db.Prepare("INSERT INTO likes(avatar_id, game_id) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare create statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(v.AvatarID, v.GameID)
	if err != nil {
		return fmt.Errorf("failed to execute create statement: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %v", err)
	}

	v.ID = int(id)
	return nil
}

// DeleteByGameID 게임 좋아요 취소
func (r *LikeRepository) DeleteByGameID(gameID, avatarID int) error {
	stmt, err := r.db.Prepare("DELETE FROM likes WHERE game_id = ? AND avatar_id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare delete statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(gameID, avatarID)
	if err != nil {
		return fmt.Errorf("failed to execute delete statement: %v", err)
	}

	return nil
}

// DeleteByCommentID 댓글 좋아요 취소
func (r *LikeRepository) DeleteByCommentID(commentID, avatarID int) error {
	stmt, err := r.db.Prepare("DELETE FROM likes WHERE comment_id = ? AND avatar_id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare delete statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(commentID, avatarID)
	if err != nil {
		return fmt.Errorf("failed to execute delete statement: %v", err)
	}

	return nil
}

// GetLikeGameByGameID 특정 게임의 게임 좋아요 리스트 조회
func (r *LikeRepository) GetLikeGameByGameID(gameID int) (map[int]*pkg.Like, error) {
	stmt, err := r.db.Prepare("SELECT id, avatar_id FROM likes WHERE game_id = ? AND comment_id IS NULL")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(gameID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to execute read statement: %v", err)
	}

	likes := make(map[int]*pkg.Like)
	for rows.Next() {
		like := pkg.Like{}
		err = rows.Scan(&like.ID, &like.AvatarID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse values: %v", err)
		}
		likes[like.AvatarID] = &like
	}

	return likes, nil
}

// GetLikeCommentByGameID 특정 게임의 댓글 좋아요 리스트 조회
func (r *LikeRepository) GetLikeCommentByGameID(gameID int) (map[int][]*pkg.Like, error) {
	stmt, err := r.db.Prepare("SELECT id, avatar_id, comment_id FROM likes WHERE game_id = ? AND comment_id IS NOT NULL")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(gameID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to execute read statement: %v", err)
	}

	likes := make(map[int][]*pkg.Like)
	for rows.Next() {
		like := pkg.Like{}
		err = rows.Scan(&like.ID, &like.AvatarID, &like.CommentID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse values: %v", err)
		}
		likes[like.AvatarID] = append(likes[like.AvatarID], &like)
	}

	return likes, nil
}

// GetLikeGameByGameID, GetLikeCommentByCommentID 합체
