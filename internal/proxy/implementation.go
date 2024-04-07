package proxy

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/mirage-deadline/proxy-price/internal/producers"
	"github.com/mirage-deadline/proxy-price/internal/providers"
	"log/slog"
)

type Implementation struct {
	ctx      context.Context
	producer producers.Producer
	provider providers.Provider
	logger   *slog.Logger
}

func NewImplementation(ctx context.Context, producer producers.Producer, provider providers.Provider, logger *slog.Logger) *Implementation {
	const op = "ProxyImplementation"

	logger = logger.With("op", op)

	return &Implementation{
		ctx:      ctx,
		producer: producer,
		provider: provider,
		logger:   logger,
	}
}

func (i *Implementation) Run() {

	var dialer websocket.Dialer
	client, _, err := dialer.Dial(i.provider.GetConnectURL(), nil)
	if err != nil {
		i.logger.Error("error dialing", err)
		return
	}
	defer client.Close()

	if err != nil {
		i.logger.Error("error dialing", slog.Any("error", err))
		return
	}

	for {
		_, message, err := client.ReadMessage()
		if err != nil {
			i.logger.Error("error reading message", slog.Any("error", err))
			return
		}
		asset, err := i.provider.ParseMessage(message)
		if err != nil {
			i.logger.Warn("error parsing message", slog.Any("error", err))
			continue
		}

		jsonAsset, err := asset.MarshalJSON()
		if err != nil {
			i.logger.Warn("error marshalling message", slog.Any("error", err))
			continue
		}

		err = i.producer.PublishMessage(i.ctx, jsonAsset)
		if err != nil {
			i.logger.Warn("error publishing message", slog.Any("error", err))
			continue
		}
	}
}
