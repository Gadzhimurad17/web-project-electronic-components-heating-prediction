package redis

import (
	"context"
	"heating-prediction/internal/config"
	"log"

	"github.com/redis/go-redis/v9"
)

func Connect(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDatabase,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("can't connect to redis server %s", err)
	}
	
	return client
}
