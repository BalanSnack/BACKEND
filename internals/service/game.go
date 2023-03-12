package service

import (
	"github.com/BalanSnack/BACKEND/internals/entity/req"
	"github.com/BalanSnack/BACKEND/internals/repository"
)

type GameService struct {
	gameRepository repository.GameRepository
}

func NewGameService(gameRepository repository.GameRepository) *GameService {
	return &GameService{
		gameRepository: gameRepository,
	}
}

func (s *GameService) GetByGameId(gameId uint64) (repository.Game, error) {
	game, err := s.gameRepository.GetByGameId(gameId)
	if err != nil {
		return repository.Game{}, err
	}
	return game, nil
}

func (s *GameService) Create(avatarId uint64, req req.CreateGameRequest) (repository.Game, error) {
	temp := repository.Game{
		AvatarId:  avatarId,
		Title:     req.Title,
		TagId:     req.TagId,
		Left:      req.Left,
		Right:     req.Right,
		LeftDesc:  req.LeftDesc,
		RightDesc: req.RightDesc,
	}

	game, err := s.gameRepository.Create(temp)
	if err != nil {
		return repository.Game{}, err
	}

	return game, nil
}

func (s *GameService) Update(avatarId uint64, req req.UpdateGameRequest) (repository.Game, error) {
	temp := repository.Game{
		GameId:    req.GameId,
		AvatarId:  avatarId,
		Title:     req.Title,
		TagId:     req.TagId,
		Left:      req.Left,
		Right:     req.Right,
		LeftDesc:  req.LeftDesc,
		RightDesc: req.RightDesc,
	}

	game, err := s.gameRepository.Update(temp)
	if err != nil {
		return repository.Game{}, err
	}

	return game, nil
}

func (s *GameService) Delete(avatarId, gameId uint64) error {
	err := s.gameRepository.Delete(avatarId, gameId)

	return err
}

func (s *GameService) GetByTagId(tagId uint64) []repository.Game {
	games := s.gameRepository.GetByTagId(tagId)

	return games
}

func (s *GameService) GetAll() []repository.Game {
	games := s.gameRepository.GetAll()

	return games
}
