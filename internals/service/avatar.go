package service

import (
	"github.com/didnlie23/go-mvc/internals/repository"
	"log"
)

type AvatarService struct {
	avatarRepository repository.AvatarRepository
}

func NewAvatarService(avatarRepository repository.AvatarRepository) *AvatarService {
	return &AvatarService{
		avatarRepository: avatarRepository,
	}
}

func (s *AvatarService) GetByAvatarId(avatarId uint64) (repository.Avatar, error) {
	avatar, err := s.avatarRepository.GetByAvatarId(avatarId)
	if err != nil {
		log.Println(err)
		return repository.Avatar{}, err
	}
	return avatar, nil
}
