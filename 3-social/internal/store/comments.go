package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID        int64  `json:"id,omitempty"`
	PostID    int64  `json:"post_id,omitempty"`
	UserID    int64  `json:"user_id,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	User      User   `json:"user,omitzero"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `
		INSERT INTO comments (post_id, user_id, content) 
		VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(ctx, query, comment.PostID, comment.UserID, comment.Content).Scan(
		&comment.ID,
		&comment.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}


func (s *CommentStore) GetByPostID(ctx context.Context, postID int64) ([]Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, u.username, u.id FROM comments AS c
		JOIN users AS u ON c.user_id = u.id
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC;
	`

	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var c Comment
		c.User = User{}
		err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.User.Username, &c.User.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}
