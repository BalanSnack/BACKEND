package repository

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  sql.NullTime
}

type Member struct {
	ID       int
	Email    string
	Provider string
	AvatarID int
}

type Avatar struct {
	ID      int
	Nick    string
	Profile string
}

type Game struct {
	ID          int
	Title       string
	LeftOption  string
	RightOption string
	LeftDesc    string
	RightDesc   string
	AvatarID    int
}

// Vote Pick: true(오른쪽) or false(왼쪽)
type Vote struct {
	ID       int  `json:"id"`
	GameID   int  `json:"game_id"`
	AvatarID int  `json:"avatar_id"`
	Pick     bool `json:"pick"`
}

type Like struct {
	ID        int `json:"id"`
	AvatarID  int `json:"avatar_id"`
	GameID    int `json:"game_id"`
	CommentID int `json:"comment_id"`
}

type Comment struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	GameID   int    `json:"game_id"`
	AvatarID int    `json:"avatar_id"`
	Content  string `json:"content"`
	Deleted  bool   `json:"deleted"`
}
