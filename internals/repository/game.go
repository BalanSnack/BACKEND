package repository

import "time"

type Game struct {
	GameId    uint64
	AvatarId  uint64
	Title     string
	TagId     uint64
	Left      string
	Right     string
	LeftDesc  string
	RightDesc string
	Update    time.Time
	ViewCount int
}

type GameRepository interface {
	GetByGameId(gameId uint64) (Game, error)
	Create(game Game) (Game, error)
	Update(game Game) (Game, error)
	Delete(avatarId, gameId uint64) error
	GetByTagId(tagId uint64) []Game
	GetAll() []Game
}
