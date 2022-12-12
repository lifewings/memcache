package command

//go:generate mockery --name Storager

import (
	"context"
)

type Storager interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) (string, error)
	Delete(ctx context.Context, key string) (string, error)
}
