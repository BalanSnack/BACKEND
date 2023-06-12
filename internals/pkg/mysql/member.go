package mysql

import (
	"database/sql"
	"fmt"
	"github.com/BalanSnack/BACKEND/internals/pkg"
)

type MemberRepository struct {
	db *sql.DB
}

func NewMemberRepository(db *sql.DB) *MemberRepository {
	return &MemberRepository{
		db: db,
	}
}

func (r *MemberRepository) Create(m *pkg.Member) error {
	stmt, err := r.db.Prepare("INSERT INTO members(email, provider, avatar_id) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare create statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(m.Email, m.Provider, m.AvatarID)
	if err != nil {
		return fmt.Errorf("failed to execute create statement: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %v", err)
	}

	m.ID = int(id)
	return nil
}

func (r *MemberRepository) Get(id int) (*pkg.Member, error) {
	stmt, err := r.db.Prepare("SELECT email, provider, avatar_id FROM members WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get statement: %v", err)
	}
	defer stmt.Close()

	var m pkg.Member
	err = stmt.QueryRow(id).Scan(&m.Email, &m.Provider, &m.AvatarID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to execute read statement: %v", err)
	}

	m.ID = id
	return &m, nil
}

func (r *MemberRepository) Update(m *pkg.Member) error {
	stmt, err := r.db.Prepare("UPDATE members SET email = ?, provider = ?, avatar_id = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare update statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.Email, m.Provider, m.AvatarID, m.ID)
	if err != nil {
		return fmt.Errorf("failed to execute update statement: %v", err)
	}

	return nil
}

func (r *MemberRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM members WHERE id = ?")
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

func (r *MemberRepository) GetByEmailAndProvider(email, provider string) (*pkg.Member, error) {
	stmt, err := r.db.Prepare("SELECT id, avatar_id FROM members WHERE email = ? AND provider = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get statement: %v", err)
	}
	defer stmt.Close()

	var m pkg.Member
	err = stmt.QueryRow(email, provider).Scan(&m.ID, &m.AvatarID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to execute read statement: %v", err)
	}

	m.Email = email
	m.Provider = provider
	return &m, nil
}
