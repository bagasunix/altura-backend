package cache

import (
	"context"
	"fmt"

	"github.com/phuslu/log"
	"github.com/redis/go-redis/v9"

	"altura-property/internal/infrastructure/config"
)

func NewCache(ctx context.Context, cfg *config.Config, logger *log.Logger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Cache.Addr(),
		Password:     cfg.Cache.Password,
		DB:           cfg.Cache.DB,
		PoolSize:     cfg.Cache.PoolSize,
		DialTimeout:  cfg.Cache.DialTimeout,
		ReadTimeout:  cfg.Cache.ReadTimeout,
		WriteTimeout: cfg.Cache.WriteTimeout,
	})

	ctxPing, cancel := context.WithTimeout(ctx, cfg.Cache.DialTimeout)
	defer cancel()

	if err := client.Ping(ctxPing).Err(); err != nil {
		client.Close()
		logger.Error().Err(err).
			Str("host", cfg.Cache.Host).
			Int("port", cfg.Cache.Port).
			Int("db", cfg.Cache.DB).
			Msg("failed to ping redis")
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	logger.Info().
		Str("host", cfg.Cache.Host).
		Int("port", cfg.Cache.Port).
		Int("db", cfg.Cache.DB).
		Int("pool_size", cfg.Cache.PoolSize).
		Msg("redis connected")

	return client, nil
}
