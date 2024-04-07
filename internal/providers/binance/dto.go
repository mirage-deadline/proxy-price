package binance

import (
	"github.com/mirage-deadline/proxy-price/internal/domain"
	"github.com/mirage-deadline/proxy-price/internal/providers"
)

func convertBinanceMessageToDomainModelMarkPrice(binanceAsset providers.AbstractMarkPriceMessage) *domain.Asset {
	return &domain.Asset{
		Symbol: binanceAsset.GetSymbol(),
		Price:  binanceAsset.GetPrice(),
	}
}
