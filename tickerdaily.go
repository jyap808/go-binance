package binance

type TickerDaily struct {
	Symbol             string  `json:"symbol"`
	PriceChange        string  `json:"priceChange"`
	PriceChangePercent string  `json:"priceChangePercent"`
	WeightedAvgPrice   string  `json:"weightedAvgPrice"`
	PrevClosePrice     string  `json:"prevClosePrice"`
	LastPrice          string  `json:"lastPrice"`
	BidPrice           float64 `json:"bidPrice,string"`
	AskPrice           string  `json:"askPrice"`
	OpenPrice          string  `json:"openPrice"`
	HighPrice          string  `json:"highPrice"`
	LowPrice           string  `json:"lowPrice"`
	Volume             string  `json:"volume"`
	QuoteVolume        float64 `json:"quoteVolume,string"`
	OpenTime           int64   `json:"openTime"`
	CloseTime          int64   `json:"closeTime"`
	FirstID            int     `json:"firstId"`
	LastID             int     `json:"lastId"`
	Count              int     `json:"count"`
}
