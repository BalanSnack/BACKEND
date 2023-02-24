package repository

import (
	"fmt"
	"sync"
)

type userMemStore struct {
	sync.Mutex

	store  map[uint64]User
	nextId uint64
}

func NewUserMemStore() *userMemStore {
	us := &userMemStore{
		store:  make(map[uint64]User),
		nextId: 0,
	}
	return us
}

func (s *userMemStore) Create(avatarId uint64, email, provider string) User {
	s.Lock()
	defer s.Unlock()

	user := User{
		UserId:   s.nextId,
		AvatarId: avatarId,
		Email:    email,
		Provider: provider}

	s.store[user.UserId] = user
	s.nextId++

	return user
}

func (s *userMemStore) Delete(userId uint64) error {
	s.Lock()
	defer s.Unlock()

	_, ok := s.store[userId]
	if !ok {
		return fmt.Errorf("user with userId=%v not found", userId)
	}

	delete(s.store, userId)

	return nil
}

func (s *userMemStore) GetByEmailAndProvider(email, provider string) (User, error) {
	s.Lock()
	defer s.Unlock()

	for _, user := range s.store {
		if user.Email == email && user.Provider == provider {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("user with email=%v, provider=%v not found", email, provider)
}
