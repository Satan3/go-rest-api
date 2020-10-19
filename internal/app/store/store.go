package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

// Close ...
func (s *Store) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}
	return nil
}

func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{store: s}
	}
	return s.userRepository
}
