package repository

import (
	"fmt"
	"sync"
	"time"
)

type gameMemStore struct {
	sync.Mutex

	store  map[uint64]Game
	nextId uint64
}

func NewGameMemStore() *gameMemStore {
	return &gameMemStore{
		store:  make(map[uint64]Game),
		nextId: 0,
	}
}

func (s *gameMemStore) GetByGameId(gameId uint64) (Game, error) {
	s.Lock()
	defer s.Unlock()

	game, ok := s.store[gameId]
	if !ok {
		return Game{}, fmt.Errorf("failed to find the game with GameId %v", gameId)
	}

	return game, nil
}

func (s *gameMemStore) Create(game Game) (Game, error) {
	s.Lock()
	defer s.Unlock()

	game.GameId = s.nextId      // 식별자 초기화
	game.Update = time.Now()    // 생성 시간 초기화
	s.store[game.GameId] = game // insert

	s.nextId++

	return game, nil
}

func (s *gameMemStore) Update(game Game) (Game, error) {
	s.Lock()
	defer s.Unlock()

	temp, ok := s.store[game.GameId]
	if ok { // 식별자로 DB에 데이터가 존재하는지 확인
		if temp.AvatarId == game.AvatarId {
			game.Update = time.Now()
			s.store[game.GameId] = game
			return game, nil
		} else {
			return Game{}, fmt.Errorf("expected avatarId %v, but got %v", temp.AvatarId, game.AvatarId)
		}
	} else {
		return Game{}, fmt.Errorf("failed to find the game with GameId %v", game.GameId)
	}
}

func (s *gameMemStore) Delete(avatarId, gameId uint64) error {
	s.Lock()
	defer s.Unlock()

	temp, ok := s.store[gameId]
	if ok {
		if temp.AvatarId == avatarId {
			delete(s.store, gameId)
			return nil
		} else {
			return fmt.Errorf("expected avatarId %v, but got %v", temp.AvatarId, avatarId)
		}
	} else {
		return fmt.Errorf("failed to find the game with GameId %v", gameId)
	}
}

func (s *gameMemStore) GetByTagId(tagId uint64) []Game {
	result := make([]Game, 1)

	for _, v := range s.store {
		if v.TagId == tagId {
			result = append(result, v)
		}
	}

	return result
}

func (s *gameMemStore) GetAll() []Game {
	result := make([]Game, 0, len(s.store))

	for _, v := range s.store {
		result = append(result, v)
	}

	return result
}
