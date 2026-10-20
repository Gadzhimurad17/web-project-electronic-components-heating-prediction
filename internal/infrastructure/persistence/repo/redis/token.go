package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenRepo struct {
	client *redis.Client
}

func NewTokenRepo(client *redis.Client) *TokenRepo {
	return &TokenRepo{
		client: client,
	}
}

const blacklistPrefix = "blacklist:"

func (r *TokenRepo) AddToBlacklistToken(ctx context.Context, token string, login string, ttl time.Duration) error {
	return r.client.Set(ctx, blacklistPrefix+token, login, ttl).Err()
}

func (r *TokenRepo) IsTokenInBlacklist(ctx context.Context, token string) (bool, error) {
	n, err := r.client.Exists(ctx, token).Result()
	if err != nil {
		return false, err
	}

	return n > 0, nil
}

func (r *TokenRepo) RemoveToken(ctx context.Context, token string) error {
	return r.client.Del(ctx, token).Err()
}

func (r *TokenRepo) Close() error {
	return r.client.Close()
}
