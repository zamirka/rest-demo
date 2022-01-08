package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // ...
)

// Sto re ,,,
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
	db, err := sql.Open("sqlite3", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM sqlite_master LIMIT 1;")
	if err != nil {
		return err
	}
	defer rows.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
