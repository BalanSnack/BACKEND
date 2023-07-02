package mysql

import (
	"BACKEND/internals/pkg"
	"database/sql"
	"fmt"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) Create(c *pkg.Comment) error {
	stmt, err := r.db.Prepare("INSERT INTO comments(parent_id, game_id, avatar_id, content, deleted) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare create statement: %c", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.ParentID, c.GameID, c.AvatarID, c.Content, c.Deleted)
	if err != nil {
		return fmt.Errorf("failed to execute create statement: %c", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %c", err)
	}

	c.ID = int(id)
	return nil
}

// Delete 완전 삭제 X, 삭제 기록
func (r *CommentRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("UPDATE comments SET deleted = 1 WHERE id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare update statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to execute update statement: %v", err)
	}

	return nil
}

func (r *CommentRepository) Update(id int, content string) error {
	stmt, err := r.db.Prepare("UPDATE comments SET content = ? WHERE id = ?")
	if err != nil {
		fmt.Errorf("failed to prepare get statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(content, id)
	if err != nil {
		return fmt.Errorf("failed to execute update statement: %v", err)
	}

	return nil
}

// GetByGameID 특정 게임의 댓글 리스트 조회
func (r *CommentRepository) GetByGameID(gameID int) ([]*pkg.Comment, error) {
	stmt, err := r.db.Prepare("SELECT id, parent_id, avatar_id, content, deleted FROM comments WHERE game_id = ? ORDER BY id")
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

	var comments []*pkg.Comment
	for rows.Next() {
		comment := pkg.Comment{}
		err = rows.Scan(&comment.ID, &comment.ParentID, &comment.AvatarID, &comment.Content, &comment.Deleted)
		if err != nil {
			return nil, fmt.Errorf("failed to parse values: %v", err)
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}
