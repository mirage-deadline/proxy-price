package providers

import "github.com/mirage-deadline/proxy-price/internal/domain"

type Provider interface {
	GetConnectURL() string
	ParseMessage(message []byte) (*domain.Asset, error)
}
