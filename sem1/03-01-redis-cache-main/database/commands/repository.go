package commands

import (
	"context"

	"redis-cache/database"
)

type Repository interface {
	AddCommand(ctx context.Context, command database.Command) error
	FindByCommand(ctx context.Context, command string) (database.Command, error)
}
