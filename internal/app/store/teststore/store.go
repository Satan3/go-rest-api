package teststore

import (
	"github.com/Satan3/go-rest-api/internal/app/model"
	"github.com/Satan3/go-rest-api/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[string]*model.User),
		}
	}

	return s.userRepository
}
