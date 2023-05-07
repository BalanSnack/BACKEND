package mysql

import "gorm.io/gorm"

const (
	JoinGame = iota
	VoteGame
	VoteComment
)

type Activity struct {
	gorm.Model
	Type      byte
	Choice    bool
	AvatarID  uint
	CommentID uint
	GameID    uint
}

type Member struct {
	gorm.Model
	Email    string
	Provider string
	AvatarID uint
}

type Avatar struct {
	gorm.Model
	Nick    string
	Profile string
}

type Game struct {
	gorm.Model
	Title       string
	LeftOption  string
	LeftDesc    string
	RightOption string
	RightDesc   string
	LeftCount   uint
	RightCount  uint
	View        uint
	Vote        int
	AvatarID    uint
}

type Comment struct {
	gorm.Model
	ParentID uint
	Content  string
	Vote     int
	AvatarID uint
	GameID   uint
}

type Tag struct {
	gorm.Model
	Name  string
	Count uint
}

type GameTag struct {
	gorm.Model
	GameID uint
	TagID  uint
}
