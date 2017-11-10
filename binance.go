// Package Binance is an implementation of the Binance API in Golang.
package binance

import (
	"encoding/json"
	"errors"
)

const (
	API_BASE                   = "https://www.binance.com/api/" // Binance API endpoint
	API_VERSION                = "v3"                           // Binance API version
	DEFAULT_HTTPCLIENT_TIMEOUT = 30                             // HTTP client timeout
)

// New return a instanciate binance struct
func New(apiKey, apiSecret string) *Binance {
	client := NewClient(apiKey, apiSecret)
	return &Binance{client}
}

// handleErr gets JSON response from Binance API en deal with error
func handleErr(r jsonResponse) error {
	if !r.Success {
		return errors.New(r.Message)
	}
	return nil
}

// binance represent a binance client
type Binance struct {
	client *client
}

// GetPrices Returns all prices data
func (b *Binance) GetPrices() (prices []*Prices, err error) {
	r, err := b.client.do("GET", "v1/ticker/allPrices", "", false)
	if err != nil {
		return
	}

	prices = make([]*Prices, 0)
	//	var response jsonResponse
	if err = json.Unmarshal(r, &prices); err != nil {
		return
	}
	if err != nil {
		return
	}

	return
}

// GetTradePairs Returns all trade pair data
//func (b *Binance) GetTradePairs() (pairs []Pair, err error) {
//	r, err := b.client.do("GET", "GetTradePairs", "", false)
//	if err != nil {
//		return
//	}
//	var response jsonResponse
//	if err = json.Unmarshal(r, &response); err != nil {
//		return
//	}
//	if err = handleErr(response); err != nil {
//		return
//	}
//	err = json.Unmarshal(response.Result, &pairs)
//	return
//}
