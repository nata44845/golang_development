package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"redis-cache/database"
	"redis-cache/database/commands"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	mongoDB, err := mongo.Connect(
		ctx, &options.ClientOptions{
			Hosts: []string{fmt.Sprintf("%s:%d", "127.0.0.1", 27017)},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoDB.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(&redis.Options{DB: 1})
	status := redisClient.Ping(ctx)
	if err := status.Err(); err != nil {
		log.Fatal(err)
	}

	commandsRepository := commands.NewMongoDbRepository(mongoDB.Database("test-db"))

	if err := commandsRepository.AddCommand(
		ctx, database.Command{
			ID:        primitive.NewObjectID(),
			Command:   "docker ps",
			CreatedAt: time.Now(),
		}); err != nil {
		log.Fatal(err)
	}

	byCommand, err := commandsRepository.FindByCommand(ctx, "docker ps")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(byCommand)
	os.Exit(0)

	cachedRepository := commands.NewCacheRepository(redisClient)

	if err := cachedRepository.AddCommand(
		ctx, database.Command{
			ID:        primitive.NewObjectID(),
			Command:   "ls -la",
			CreatedAt: time.Now().UTC(),
		},
	); err != nil {
		return
	}

	command, err := cachedRepository.FindByCommand(ctx, "ls -la")
	if err != nil {
		log.Fatal(err)
	}

	_, err = cachedRepository.FindByCommand(ctx, "ls -lat")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(command.Command)

	_ = commandsRepository
	_ = cachedRepository
	_ = mongoDB

	<-ctx.Done()
}
