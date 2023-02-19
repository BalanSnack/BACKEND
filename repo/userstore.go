package repo

import (
	"fmt"
	"sync"
)

type UserMemStore struct {
	sync.Mutex

	store  map[uint64]User
	nextId uint64
}

func NewUserMemStore() *UserMemStore {
	us := &UserMemStore{
		store:  make(map[uint64]User),
		nextId: 0,
	}
	return us
}

func (s *UserMemStore) Create(avatarId uint64, email, provider string) User {
	s.Lock()
	defer s.Unlock()

	user := User{
		Id:       s.nextId,
		AvatarId: avatarId,
		Email:    email,
		Provider: provider}

	s.store[user.Id] = user
	s.nextId++

	return user
}

func (s *UserMemStore) Delete(id uint64) error {
	s.Lock()
	defer s.Unlock()

	_, ok := s.store[id]
	if !ok {
		return fmt.Errorf("user with id=%v not found", id)
	}

	delete(s.store, id)

	return nil
}

func (s *UserMemStore) GetUserByEmailAndProvider(email, provider string) (User, error) {
	s.Lock()
	defer s.Unlock()

	for _, user := range s.store {
		if user.Email == email && user.Provider == provider {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("user with email=%v, provider=%v not found", email, provider)
}
