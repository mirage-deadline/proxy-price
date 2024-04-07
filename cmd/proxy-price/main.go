package main

import (
	"context"
	"github.com/mirage-deadline/proxy-price/internal/apps"
	"github.com/mirage-deadline/proxy-price/internal/config"
	"github.com/mirage-deadline/proxy-price/internal/pkg"
	"github.com/mirage-deadline/proxy-price/internal/producers/redis"
	"github.com/mirage-deadline/proxy-price/internal/providers/binance"
	"github.com/mirage-deadline/proxy-price/internal/proxy"
)

func main() {

	ctx := context.Background()
	cfg := config.MustGetConfig()
	logger := pkg.NewLogger()

	// Exchange provider section
	provider := binance.NewProvider(logger, cfg.Symbols, cfg.BinanceURL)

	// Producer section
	redisClient := apps.MustGetRedisClient(cfg.RedisDSN, cfg.RedisPassword)
	producer := redis.NewProducer(redisClient, logger, cfg.Topic)

	impl := proxy.NewImplementation(ctx, producer, provider, logger)
	impl.Run()
}
