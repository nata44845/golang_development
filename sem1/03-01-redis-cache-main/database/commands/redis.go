package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"

	"redis-cache/database"
)

// Проверка на соответствие структуры интерфейсу
var _ Repository = (*redisRepository)(nil)

func NewCacheRepository(db *redis.Client, repository Repository) Repository {
	return &redisRepository{db: db, dbLayer: repository}
}

type redisRepository struct {
	db      *redis.Client
	dbLayer Repository
}

func (r redisRepository) AddCommand(ctx context.Context, command database.Command) error {
	if err := r.dbLayer.AddCommand(ctx, command); err != nil {
		return fmt.Errorf("redis AddCommand: %w", err)
	}
	return nil
}

func (r redisRepository) FindByCommand(ctx context.Context, command string) (database.Command, error) {
	var cmd database.Command
	bytes, err := r.db.Get(ctx, command).Bytes()
	if err != nil {
		if err == redis.Nil {
			byCommand, err := r.dbLayer.FindByCommand(ctx, command)
			if err != nil {
				return cmd, fmt.Errorf("dbLayer FindByCommand: %w", err)
			}
			encoded, err := json.Marshal(byCommand)
			if err != nil {
				return cmd, fmt.Errorf("json Marshal: %w", err)
			}
			//-1 время жизни ключа
			status := r.db.Set(ctx, command, encoded, 0)
			if err := status.Err(); err != nil {
				return cmd, fmt.Errorf("redis Set: %w", err)
			}
			return byCommand, nil
		}
		return cmd, fmt.Errorf("redis Get: %w", err)
	}
	if err := json.Unmarshal(bytes, &cmd); err != nil {
		return cmd, fmt.Errorf("json Unmarshal: %w", err)
	}
	return cmd, nil
}
