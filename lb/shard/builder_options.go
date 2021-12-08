package shard

import (
	"context"
)

type builderOption func(*builderOptions) error
type builderOptions struct {
	shardKeyFunc func(context.Context) string
}

func defaultBuilderOptions() builderOptions {
	return builderOptions{
		shardKeyFunc: func(ctx context.Context) string {
			key, ok := ctx.Value(DefaultShardKeyName).(string)
			if !ok {
				return "NOT_FOUND"
			}
			return key
		},
	}
}
