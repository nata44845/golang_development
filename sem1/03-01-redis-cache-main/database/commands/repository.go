package commands

import (
	"context"
	// Имя модуля/имя пакета
	"redis-cache/database"
)

type Repository interface {
	AddCommand(ctx context.Context, command database.Command) error
	FindByCommand(ctx context.Context, command string) (database.Command, error)
}
