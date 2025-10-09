package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	UserID    int64     `json:"user_id,omitempty"`
	Content   string    `json:"content,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	Version   int       `json:"version,omitempty"`
	CreatedAt string    `json:"created_at,omitempty"`
	UpdatedAt string    `json:"updated_at,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
	User      User      `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentCount *int `json:"comment_count,omitempty"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `
		INSERT INTO posts (content, title, user_id, tags) 
		VALUES ($1, $2, $3, $4) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags)).Scan(
		&post.ID,
		&post.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetByID(ctx context.Context, postID int64) (*Post, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `
		SELECT id, title, user_id, content, tags, created_at, updated_at, version
		FROM posts WHERE id = $1
	`

	var post Post
	err := s.db.QueryRowContext(ctx, query, postID).Scan(
		&post.ID,
		&post.Title,
		&post.UserID,
		&post.Content,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (s *PostStore) Delete(ctx context.Context, postID int64) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `DELETE FROM posts WHERE id = $1`
	if _, err := s.db.ExecContext(ctx, query, postID); err != nil {
		return err
	}

	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `
		UPDATE posts
		SET title = $1, content = $2, version = version + 1
		WHERE id = $3 AND version = $4
		RETURNING version
	`
	if err := s.db.QueryRowContext(ctx, query, post.Title, post.Content, post.ID, post.Version).Scan(&post.Version); err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *PostStore) GetUserFeed(ctx context.Context, userID int64) ([]PostWithMetadata, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `
		SELECT p.id, p.user_id, u.username, p.title, p.content, p.created_at, p.version, p.tags, 
		COUNT(c.id) AS comments_count
		FROM posts p 
		LEFT JOIN comments c ON c.post_id = p.id
		LEFT JOIN users u ON p.user_id = u.id
		LEFT JOIN followers f ON f.user_id = p.user_id
		WHERE (f.user_id = p.user_id AND f.follower_id = $1) OR p.user_id = $1
		GROUP BY p.id, u.username
		ORDER BY p.created_at DESC;
	`

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feed []PostWithMetadata
	for rows.Next() {
		var p PostWithMetadata
		err = rows.Scan(
			&p.ID, 
			&p.UserID, 
			&p.User.Username, 
			&p.Title, 
			&p.Content, 
			&p.CreatedAt, 
			&p.Version, 
			pq.Array(&p.Tags), 
			&p.CommentCount,
		)
		if err != nil {
			return nil, err
		}

		feed = append(feed, p)
	}

	return feed, nil
}
