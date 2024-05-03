package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Command struct {
	ID        primitive.ObjectID `json:"id" bson:"id"`
	Command   string             `json:"command" bson:"command"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
