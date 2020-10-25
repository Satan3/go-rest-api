package sqlstore

import (
	"database/sql"
	"github.com/Satan3/go-rest-api/internal/app/store"
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{store: s}
	}
	return s.userRepository
}
