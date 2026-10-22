package repo

import (
	"context"
	"time"
)

type TokenRepo interface {
	AddToBlacklistToken(ctx context.Context, token string, login string, ttl time.Duration)
	IsTokenInBlacklist(ctx context.Context, token string) (bool, error)
	RemoveToken(ctx context.Context, token string) error
	Close() error
}
