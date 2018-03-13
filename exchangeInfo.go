package binance

// Result from: GET /api/v3/exchangeInfo

// ExchangeInfo exchange info
type ExchangeInfo struct {
	Symbols []Symbol `json:"symbols"`
}

// Symbol market symbol
type Symbol struct {
	Symbol             string              `json:"symbol"`
	Status             string              `json:"status"`
	BaseAsset          string              `json:"baseAsset"`
	BaseAssetPrecision int                 `json:"baseAssetPrecision"`
	QuoteAsset         string              `json:"quoteAsset"`
	QuotePrecision     int                 `json:"quotePrecision"`
	OrderTypes         []string            `json:"orderTypes"`
	IcebergAllowed     bool                `json:"icebergAllowed"`
	Filters            []map[string]string `json:"filters"`
}

//type ExchangeInfo struct {
//	ExchangeFilters []string     `json:"ExchangeFilters"`
//	RateLimits      []RateLimit  `json:"rateLimits"`
//	ServerTime      int64        `json:"serverTime"`
//	Symbols         []SymbolInfo `json:"symbols"`
//	TimeZone        string       `json:"timezone"`
//}

//type SymbolInfo struct {
//	Symbol             string         `json:"symbol"`
//	BaseAsset          string         `json:"baseAsset"`
//	QuotePrecision     int64          `json:"quotePrecision"`
//	BaseAssetPrecision int64          `json:"baseAssetPrecision"`
//	Status             string         `json:"status"`
//	OrderTypes         []string       `json:"orderTypes"`
//	Filters            []SymbolFilter `json:"filters"`
//	QuoteAsset         string         `json:"quoteAsset"`
//	IceBergAllowed     bool           `json:"icebergAllowed"`
//}

//type SymbolFilter struct {
//	Type        string  `json:"filterType"`
//	MinPrice    float64 `json:"minPrice,string"`
//	MaxPrice    float64 `json:"maxPrice,string"`
//	TickSize    float64 `json:"tickSize,string"`
//	StepSize    float64 `json:"stepSize,string"`
//	MinQty      float64 `json:"minQty,string"`
//	MaxQty      float64 `json:"maxQty,string"`
//	MinNotional float64 `json:"minNotional,string"`
//}

//type RateLimit struct {
//	Limit         int64  `json:"limit"`
//	Interval      string `json:"interval"`
//	RateLimitType string `json:"rateLimitType"`
//}

/*
{"timezone":"UTC",
"serverTime":1520893381431,
"rateLimits":[{"rateLimitType":"REQUESTS","interval":"MINUTE","limit":1200},{"rateLimitType":"ORDERS","interval":"SECOND","limit":10},{"rateLimitType":"ORDERS","interval":"DAY","limit":100000}],
"exchangeFilters":[],
"symbols":[{"symbol":"ETHBTC","status":"TRADING","baseAsset":"ETH","baseAssetPrecision":8,"quoteAsset":"BTC","quotePrecision":8,"orderTypes":["LIMIT","LIMIT_MAKER","MARKET","STOP_LOSS_LIMIT","TAKE_PROFIT_LIMIT"],"icebergAllowed":true,"filters":[{"filterType":"PRICE_FILTER","minPrice":"0.00000100","maxPrice":"100000.00000000","tickSize":"0.00000100"},{"filterType":"LOT_SIZE","minQty":"0.00100000","maxQty":"100000.00000000","stepSize":"0.00100000"},{"filterType":"MIN_NOTIONAL","minNotional":"0.00100000"}]},
           {"symbol":"LTCBTC","status":"TRADING","baseAsset":"LTC","baseAssetPrecision":8,"quoteAsset":"BTC","quotePrecision":8,"orderTypes":["LIMIT","LIMIT_MAKER","MARKET","STOP_LOSS_LIMIT","TAKE_PROFIT_LIMIT"],"icebergAllowed":true,"filters":[{"filterType":"PRICE_FILTER","minPrice":"0.00000100","maxPrice":"100000.00000000","tickSize":"0.00000100"},{"filterType":"LOT_SIZE","minQty":"0.01000000","maxQty":"100000.00000000","stepSize":"0.01000000"},{"filterType":"MIN_NOTIONAL","minNotional":"0.00100000"}]},{"symbol":"BNBBTC","status":"TRADING","baseAsset":"BNB","baseAssetPrecision":8,"quoteAsset":"BTC","quotePrecision":8,"orderTypes":["LIMIT","LIMIT_MAKER","MARKET","STOP_LOSS_LIMIT","TAKE_PROFIT_LIMIT"],"icebergAllowed":true,"filters":[{"filterType":"PRICE_FILTER","minPrice":"0.00000010","maxPrice":"100000.00000000","tickSize":"0.00000010"}
*/
