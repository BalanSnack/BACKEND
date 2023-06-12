package mysql

import (
	"database/sql"
	"fmt"
	"github.com/BalanSnack/BACKEND/internals/pkg"
	//_ "github.com/go-sql-driver/mysql"
)

type AvatarRepository struct {
	db *sql.DB
}

func NewAvatarRepository(db *sql.DB) *AvatarRepository {
	return &AvatarRepository{
		db: db,
	}
}

func (r *AvatarRepository) Create(a *pkg.Avatar) error {
	stmt, err := r.db.Prepare("INSERT INTO avatars(nick, profile) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare create statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(a.Nick, a.Profile)
	if err != nil {
		return fmt.Errorf("failed to execute create statement: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %v", err)
	}

	a.ID = int(id)
	return nil
}

func (r *AvatarRepository) Get(id int) (*pkg.Avatar, error) {
	stmt, err := r.db.Prepare("SELECT nick, profile FROM avatars WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare read statement: %v", err)
	}
	defer stmt.Close()

	var a pkg.Avatar
	err = stmt.QueryRow(id).Scan(&a.Nick, &a.Profile)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to execute read statement: %v", err)
	}

	a.ID = id
	return &a, nil
}

func (r *AvatarRepository) Update(a *pkg.Avatar) error {
	stmt, err := r.db.Prepare("UPDATE avatars SET nick = ?, profile = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare update statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Nick, a.Profile, a.ID)
	if err != nil {
		return fmt.Errorf("failed to execute update statement: %v", err)
	}

	return nil
}

func (r *AvatarRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM avatars WHERE id = ?")
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
