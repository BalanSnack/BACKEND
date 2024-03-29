package service

import (
	"BACKEND/internals/pkg"
	"BACKEND/internals/pkg/mysql"
	"log"
)

type AvatarService struct {
	avatarRepository *mysql.AvatarRepository
}

func NewAvatarService(avatarRepository *mysql.AvatarRepository) *AvatarService {
	return &AvatarService{
		avatarRepository: avatarRepository,
	}
}

func (s *AvatarService) GetByAvatarId(avatarID int) (*pkg.Avatar, error) {
	avatar, err := s.avatarRepository.Get(avatarID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return avatar, nil
}
