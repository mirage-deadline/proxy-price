package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type Producer struct {
	client  *redis.Client
	logger  *slog.Logger
	channel string
}

func NewProducer(client *redis.Client, logger *slog.Logger, channel string) *Producer {
	const op = "RedisProducer"

	logger = logger.With("op", op)

	return &Producer{
		client:  client,
		logger:  logger,
		channel: channel,
	}
}

func (p *Producer) PublishMessage(ctx context.Context, message []byte) error {
	err := p.client.Publish(ctx, p.channel, message).Err()
	if err != nil {
		p.logger.Error("error publishing message", slog.Any("error", err))
		return err
	}
	return nil
}
