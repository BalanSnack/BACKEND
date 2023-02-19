package service

import (
	"github.com/BalanSnack/BACKEND/repo"
	"github.com/sirupsen/logrus"
)

type AvatarService struct {
	avatarRepo repo.AvatarRepo
}

func NewAvatarService(avatarRepo repo.AvatarRepo) AvatarService {
	return AvatarService{avatarRepo: avatarRepo}
}

func (a AvatarService) GetAvatarById(id uint64) (repo.Avatar, error) {
	avatar, err := a.avatarRepo.GetAvatarById(id)
	if err != nil {
		logrus.Debug(err)
		return repo.Avatar{}, err
	}
	return avatar, nil
}
