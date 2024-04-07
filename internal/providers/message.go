package providers

type AbstractMarkPriceMessage interface {
	GetSymbol() string
	GetPrice() string
}
