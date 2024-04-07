package binance

import (
	"github.com/mailru/easyjson"
	"github.com/mirage-deadline/proxy-price/internal/domain"
	"log/slog"
	"strings"
)

type Provider struct {
	symbols      []string
	baseURL      string
	logger       *slog.Logger
	markInterval string
}

func NewProvider(logger *slog.Logger, symbols []string, url string) *Provider {
	const op = "BinanceProvider"
	logger = logger.With("op", op)

	return &Provider{
		baseURL:      url,
		logger:       logger,
		symbols:      symbols,
		markInterval: "@1s",
	}
}

func (p *Provider) GetConnectURL() string {
	builder := strings.Builder{}
	builder.WriteString(p.baseURL)

	for idx, symbol := range p.symbols {
		builder.WriteString(strings.ToLower(symbol))
		builder.WriteString("@markPrice")
		builder.WriteString(p.markInterval)

		if idx != len(p.symbols)-1 {
			builder.WriteString("/")
		}
	}
	return builder.String()
}

func (p *Provider) ParseMessage(message []byte) (*domain.Asset, error) {
	var providerAsset binanceMarkPriceMessage
	err := easyjson.Unmarshal(message, &providerAsset)
	if err != nil {
		p.logger.Error("error unmarshalling message", slog.Any("error", err))
		return nil, err
	}
	return convertBinanceMessageToDomainModelMarkPrice(providerAsset), nil
}
