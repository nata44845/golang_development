package users

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"gitlab.com/robotomize/gb-golang/homework/03-01-umanager/internal/database"
)

func New(userDB *pgx.Conn, timeout time.Duration) *Repository {
	return &Repository{userDB: userDB, timeout: timeout}
}

type Repository struct {
	userDB  *pgx.Conn
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateUserReq) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	create_time := time.Now()

	query := `
		INSERT INTO users (id, username, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE
		SET username = $2, password = $3, updated_at = $5
	`
	if _, err := r.userDB.Exec(ctx, query, req.ID, req.Username, req.Password, create_time, create_time); err != nil {
		return u, fmt.Errorf("postgres Exec: %w", err)
	}

	return u, nil
}

func (r *Repository) FindByID(ctx context.Context, userID uuid.UUID) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	// implement me
	return u, nil
}

func (r *Repository) FindByUsername(ctx context.Context, username string) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	// implement me
	return u, nil
}
