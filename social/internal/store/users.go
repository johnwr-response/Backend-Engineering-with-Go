package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(_ context.Context) error {
	return nil
}
