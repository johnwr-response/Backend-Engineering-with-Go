package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type Follower struct {
	UserID     int64 `json:"user_id"`
	FollowerID int64 `json:"follower_id"`
	CreatedAt  int64 `json:"created_at"`
}

type FollowerStore struct {
	db *sql.DB
}

func (s *FollowerStore) Follow(ctx context.Context, followerID, userID int64) error {
	query := `
      INSERT INTO followers(user_id, follower_id) VALUES ($1, $2)
    `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	if err != nil {

		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
			return ErrUniqueConstraintConflict
		}

		// None of the below code from the course works. Replaced by the above code

		//log.Println("Follower Store Error:", err)
		//if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
		//	return ErrUniqueConstraintConflict
		//}
		//var pqErr *pq.Error
		//if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		//if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
		//return fmt.Errorf("duplicate key value violates unique constraint \"followers_user_id_fkey\"")
		//}
	}
	return nil
}

func (s *FollowerStore) Unfollow(ctx context.Context, followerID, userID int64) error {
	query := `
      DELETE FROM followers WHERE user_id = $1 AND follower_id = $2
    `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	return err
}
