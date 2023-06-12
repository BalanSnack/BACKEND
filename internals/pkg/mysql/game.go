package mysql

import (
	"database/sql"
	"fmt"
	"github.com/BalanSnack/BACKEND/internals/pkg"
)

type GameRepository struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) *GameRepository {
	return &GameRepository{
		db: db,
	}
}

func (r *GameRepository) Create(g *pkg.Game) error {
	stmt, err := r.db.Prepare("INSERT INTO games(title, left_option, right_option, left_desc, right_desc, avatar_id) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare create statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(g.Title, g.LeftOption, g.RightOption, g.LeftDesc, g.RightDesc, g.AvatarID)
	if err != nil {
		return fmt.Errorf("failed to execute create statement: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %v", err)
	}

	g.ID = int(id)
	return nil
}

func (r *GameRepository) Get(id int) (*pkg.Game, error) {
	stmt, err := r.db.Prepare("SELECT title, left_option, right_option, left_desc, right_desc, avatar_id FROM games WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get statement: %v", err)
	}
	defer stmt.Close()

	var g pkg.Game
	err = stmt.QueryRow(id).Scan(&g.Title, &g.LeftOption, &g.RightOption, &g.LeftDesc, &g.RightDesc, &g.AvatarID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to execute read statement: %v", err)
	}

	g.ID = id
	return &g, nil
}

func (r *GameRepository) Update(g *pkg.Game) error {
	stmt, err := r.db.Prepare("UPDATE games SET title = ?, left_option = ?, right_option = ?, left_desc = ?, right_desc = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare update statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(g.Title, g.LeftOption, g.RightOption, g.LeftDesc, g.RightDesc, g.ID)
	if err != nil {
		return fmt.Errorf("failed to execute update statement: %v", err)
	}

	return nil
}

func (r *GameRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM games WHERE id = ?")
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

// GetNextRandomGame 유저가 참여하지 않는 게임 중 랜덤 ID 조회, 중복 조회를 피하기 위해 현재 조회 중인 gameID 제외
func (r *GameRepository) GetNextRandomGame(avatarID, gameID int) (int, error) {
	stmt, err := r.db.Prepare("SELECT id FROM games WHERE id NOT IN (SELECT game_id FROM votes WHERE avatar_id = ?) AND id != ? ORDER BY RAND() LIMIT 1")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare get statement: %v", err)
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(avatarID, gameID).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to execute read statement: %v", err)
	}

	return id, nil
}

// GetNextRecentGame 유저가 참여하지 않은 최신 게임 ID 조회, 중복 조회를 피하기 위해 현재 조회 중인 gameID 제외
func (r *GameRepository) GetNextRecentGame(avatarID, gameID int) (int, error) {
	stmt, err := r.db.Prepare("SELECT id FROM games WHERE id NOT IN (SELECT game_id FROM votes WHERE avatar_id = ?) AND id != ? ORDER BY id DESC LIMIT 1")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare get statement: %v", err)
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(avatarID, gameID).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to execute read statement: %v", err)
	}

	return id, nil
}
