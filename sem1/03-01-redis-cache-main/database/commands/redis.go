package commands

import (
	"context"

	"github.com/go-redis/redis/v8"

	"redis-cache/database"
)

var _ Repository = (*redisRepository)(nil)

func NewCacheRepository(db *redis.Client) Repository {
	return &redisRepository{db: db}
}

type redisRepository struct {
	db *redis.Client
}

func (r redisRepository) AddCommand(ctx context.Context, command database.Command) error {
	// TODO implement me
	panic("implement me")
}

func (r redisRepository) FindByCommand(ctx context.Context, command string) (database.Command, error) {
	// TODO implement me
	panic("implement me")
}
