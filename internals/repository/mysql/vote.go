package mysql

import (
	"database/sql"
	"fmt"
	"github.com/BalanSnack/BACKEND/internals/repository"
)

type VoteRepository struct {
	db *sql.DB
}

func NewVoteRepository(db *sql.DB) *VoteRepository {
	return &VoteRepository{
		db: db,
	}
}

func (r *VoteRepository) Create(v *repository.Vote) error {
	stmt, err := r.db.Prepare("INSERT INTO votes(game_id, avatar_id, pick) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare create statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(v.GameID, v.AvatarID, v.Pick)
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

func (r *VoteRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM votes WHERE id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare delete statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to execute delete statement: %v", err)
	}

	return nil
}

// GetByGameID 특정 게임의 참여 리스트 조회
func (r *VoteRepository) GetByGameID(gameID int) (map[int]*repository.Vote, error) {
	stmt, err := r.db.Prepare("SELECT id, game_id, avatar_id, pick FROM votes WHERE game_id = ?")
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

	votes := make(map[int]*repository.Vote)
	for rows.Next() {
		vote := repository.Vote{}
		err = rows.Scan(&vote.ID, &vote.GameID, &vote.AvatarID, &vote.Pick)
		if err != nil {
			return nil, fmt.Errorf("failed to parse values: %v", err)
		}
		votes[vote.AvatarID] = &vote
	}

	return votes, nil
}
