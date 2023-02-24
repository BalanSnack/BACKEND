package repository

import (
	"fmt"
	"sync"
)

type avatarMemStore struct {
	sync.Mutex

	store  map[uint64]Avatar
	nextId uint64
}

func NewAvatarMemStore() *avatarMemStore {
	as := &avatarMemStore{
		store:  make(map[uint64]Avatar),
		nextId: 0}

	return as
}

func (s *avatarMemStore) Create(nickname, profile string, anonymity bool) Avatar {
	s.Lock()
	defer s.Unlock()

	avatar := Avatar{
		AvatarId:  s.nextId,
		Nickname:  nickname,
		Profile:   profile,
		Anonymity: anonymity}

	s.store[avatar.AvatarId] = avatar
	s.nextId++

	return avatar
}

func (s *avatarMemStore) Update(avatarId uint64, nickname, profile string) (Avatar, error) {
	s.Lock()
	defer s.Unlock()

	avatar, ok := s.store[avatarId]
	if !ok {
		return Avatar{}, fmt.Errorf("avatar with avatarId=%v not found", avatarId)
	}

	avatar.Nickname = nickname
	avatar.Profile = profile
	s.store[avatarId] = avatar

	return avatar, nil
}

func (s *avatarMemStore) Delete(avatarId uint64) error {
	s.Lock()
	defer s.Unlock()

	_, ok := s.store[avatarId]
	if !ok {
		return fmt.Errorf("avatar with avatarId=%v not found", avatarId)
	}

	delete(s.store, avatarId)

	return nil
}

func (s *avatarMemStore) GetByAvatarId(avatarId uint64) (Avatar, error) {
	s.Lock()
	defer s.Unlock()

	avatar, ok := s.store[avatarId]
	if !ok {
		return Avatar{}, fmt.Errorf("avatar with avatarId=%v not found", avatarId)
	}

	return avatar, nil
}
