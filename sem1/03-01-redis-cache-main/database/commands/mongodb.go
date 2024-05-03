package commands

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"redis-cache/database"
)

var _ Repository = (*mongoDbRepository)(nil)

const collection = "commands"

func NewMongoDbRepository(db *mongo.Database) Repository {
	return &mongoDbRepository{db: db}
}

type mongoDbRepository struct {
	db *mongo.Database
}

func (m mongoDbRepository) AddCommand(ctx context.Context, command database.Command) error {
	if _, err := m.db.Collection(collection).InsertOne(ctx, command); err != nil {
		return fmt.Errorf("mongo InsertOne: %w", err)
	}
	return nil
}

func (m mongoDbRepository) FindByCommand(ctx context.Context, command string) (database.Command, error) {
	var cmd database.Command
	result := m.db.Collection(collection).FindOne(ctx, bson.M{"command": command})
	if err := result.Err(); err != nil {
		return cmd, fmt.Errorf("mongo FindOne: %w", err)
	}
	if err := result.Decode(&cmd); err != nil {
		return cmd, fmt.Errorf("mongo Decode: %w", err)
	}
	return cmd, nil
}
