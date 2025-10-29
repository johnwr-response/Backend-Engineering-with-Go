package store

import (
	"context"
	"database/sql"
	"encoding/gob"
	"errors"
	"log"

	"time"

	_ "github.com/go-pg/pg/v10"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `pg:",array"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  []Comment `json:"comments"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {

	//goland:noinspection ALL
	query := `
		INSERT INTO posts(content, title, user_id, tags) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at, updated_at
	`
	gob.Register(post.Tags)
	err := s.db.QueryRowContext(
		ctx, query, post.Content, post.Title, post.UserID, nil,
		// Arrays are not currently working. Fix later by testing different ORMs, like bun or gorm
		//ctx, query, post.Content, post.Title, post.UserID, pg.Array(post.Tags),
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		log.Printf("Tags type: %T", post.Tags)
		log.Printf("Tags are : %v", post.Tags)
		return err
	}

	return nil
}
func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `
-- 		SELECT id, user_id, title, content, created_at, updated_at, tags
		SELECT id, user_id, title, content, created_at, updated_at
		FROM posts 
		WHERE id = $1
	`

	var post Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
		// Arrays are not currently working. Fix later by testing different ORMs, like bun or gorm
		// OR is this just a matter of wrong library? Should it have been from https://github.com/lib/pq pq.Array instead of pg.Array?
		//pg.Array(&post.Tags),
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}
