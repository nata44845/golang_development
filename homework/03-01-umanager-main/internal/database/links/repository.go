package links

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/robotomize/gb-golang/homework/03-01-umanager/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collection = "links"

func New(db *mongo.Database, timeout time.Duration) *Repository {
	return &Repository{db: db, timeout: timeout}
}

type Repository struct {
	db      *mongo.Database
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateReq) (database.Link, error) {
	var l database.Link
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	create_time := time.Now()

	l = database.Link{
		ID:        req.ID,
		Title:     req.Title,
		URL:       req.URL,
		Images:    req.Images,
		Tags:      req.Tags,
		UserID:    req.UserID,
		CreatedAt: create_time,
		UpdatedAt: create_time,
	}

	if _, err := r.db.Collection(collection).InsertOne(ctx, l); err != nil {
		return l, fmt.Errorf("mongo InsertOne: %w", err)
	}

	return l, nil
}

func (r *Repository) FindByUserAndURL(ctx context.Context, link, userID string) (database.Link, error) {
	var l database.Link
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	result := r.db.Collection(collection).FindOne(ctx, bson.M{"url": link, "user_id": userID})
	if err := result.Err(); err != nil {
		return l, fmt.Errorf("mongo FindOne: %w", err)
	}

	if err := result.Decode(&l); err != nil {
		return l, fmt.Errorf("mongo Decode: %w", err)
	}

	return l, nil
}

func (r *Repository) FindByCriteria(ctx context.Context, criteria Criteria) ([]database.Link, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var links []database.Link

	filter := bson.M{}
	opts := options.Find()
	if criteria.Limit != nil {
		opts.SetLimit(*criteria.Limit)
	}
	if criteria.Offset != nil {
		opts.SetSkip(*criteria.Offset)
	}
	if criteria.UserID != nil {
		filter["user_id"] = *criteria.UserID
	}
	if len(criteria.Tags) > 0 {
		tagsCriteria := make([]interface{}, 0, len(criteria.Tags))
		for _, tag := range criteria.Tags {
			tagsCriteria = append(tagsCriteria, tag)
		}

		filter["tags"] = bson.M{"$in": tagsCriteria}
	}

	cursor, err := r.db.Collection(collection).Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("mongo Find: %w", err)
	}

	for cursor.Next(ctx) {
		var l database.Link
		if err := cursor.Decode(&l); err != nil {
			return nil, fmt.Errorf("mongo Decode: %w", err)
		}
		links = append(links, l)
	}

	return links, nil
}
