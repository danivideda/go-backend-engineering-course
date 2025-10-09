package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Follower struct {
	UserID     int64  `json:"user_id,omitempty"`
	FollowerID int64  `json:"follower_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

type FollowerStore struct {
	db *sql.DB
}

func (s *FollowerStore) Follow(ctx context.Context, followerID, userID int64) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `
		INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)
	`
	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}
	}
	return err
}

func (s *FollowerStore) Unfollow(ctx context.Context, followerID, userID int64) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `
		DELETE FROM followers 
		WHERE user_id = $1 AND follower_id = $2
	`
	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	switch err {
	case sql.ErrNoRows:
		return ErrNotFound
	default:
		return err
	}
}
