package repo

import (
	"fmt"
	"sync"
)

type AvatarMemStore struct {
	sync.Mutex

	store  map[uint64]Avatar
	nextId uint64
}

func NewAvatarMemStore() *AvatarMemStore {
	as := &AvatarMemStore{
		store:  make(map[uint64]Avatar),
		nextId: 0}

	return as
}

func (s *AvatarMemStore) Create(nickname, profile string, anonymity bool) Avatar {
	s.Lock()
	defer s.Unlock()

	avatar := Avatar{
		Id:        s.nextId,
		Nickname:  nickname,
		Profile:   profile,
		Anonymity: anonymity}

	s.store[avatar.Id] = avatar
	s.nextId++

	return avatar
}

func (s *AvatarMemStore) Update(id uint64, nickname, profile string) (Avatar, error) {
	s.Lock()
	defer s.Unlock()

	avatar, ok := s.store[id]
	if !ok {
		return Avatar{}, fmt.Errorf("avatar with id=%v not found", id)
	}

	avatar.Nickname = nickname
	avatar.Profile = profile
	s.store[id] = avatar

	return avatar, nil
}

func (s *AvatarMemStore) Delete(id uint64) error {
	s.Lock()
	defer s.Unlock()

	_, ok := s.store[id]
	if !ok {
		return fmt.Errorf("avatar with id=%v not found", id)
	}

	delete(s.store, id)

	return nil
}

func (s *AvatarMemStore) GetAvatarById(id uint64) (Avatar, error) {
	s.Lock()
	defer s.Unlock()

	avatar, ok := s.store[id]
	if !ok {
		return Avatar{}, fmt.Errorf("avatar with id=%v not found", id)
	}

	return avatar, nil
}
