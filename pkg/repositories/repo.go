package repositories

import (
	"database/sql"
	"errors"
	"time"
)

type Repo struct {
	timeout time.Duration
	db      *sql.DB
	DBError DBError
}

type DBError struct {
	ErrNotFound     error
	ErrDuplicateKey error
	ErrInvalidInput error
}

// var (
// 	ErrNotFound     = errors.New("record not found")
// 	ErrDuplicateKey = errors.New("duplicate key violation")
// 	ErrInvalidInput = errors.New("invalid input")
// )

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		timeout: time.Second * 3,
		db:      db,
		DBError: DBError{ErrNotFound: errors.New("record not found"),
			ErrDuplicateKey: errors.New("duplicate key violation"),
			ErrInvalidInput: errors.New("invalid input"),
		},
	}
}

func (r *Repo) IsUniqueConstraintError(err error) bool {
	// SQLite error message contains "UNIQUE constraint failed"
	return err != nil && (r.Contains(err.Error(), "UNIQUE constraint failed") ||
		r.Contains(err.Error(), "unique constraint"))
}

func (r *Repo) Contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && r.ContainsAt(s, substr, 0))
}

func (r *Repo) ContainsAt(s, substr string, start int) bool {
	for i := start; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
