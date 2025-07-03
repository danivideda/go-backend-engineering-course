package store

import (
	"context"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(context.Context) error {
	return nil
}