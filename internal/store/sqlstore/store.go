package sqlstore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // ...
	"github.com/zamirka/rest-demo/internal/store"
)

// Sto re ,,,
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

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
