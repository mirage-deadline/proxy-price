package binance

//easyjson:json
type markPriceMessage struct {
	EventType  string `json:"e"`
	EventTime  int64  `json:"E"`
	Symbol     string `json:"s"`
	MarkPrice  string `json:"p"`
	IndexPrice string `json:"i"`
	// Estimated Settle Price, only useful in the last hour before the settlement starts
	EstimatedSettlePrice string `json:"P"`
	FundingRate          string `json:"r"`
	NextFundingTime      int64  `json:"T"`
}

//easyjson:json
type binanceMarkPriceMessage struct {
	Stream string           `json:"stream"`
	Data   markPriceMessage `json:"data"`
}

func (e binanceMarkPriceMessage) GetSymbol() string {
	return e.Data.Symbol
}

func (e binanceMarkPriceMessage) GetPrice() string {
	return e.Data.MarkPrice
}
