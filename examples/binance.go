package main

import (
	"fmt"
	"strings"

	"github.com/jyap808/go-binance"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func main() {
	// Binance client
	binance := binance.New(API_KEY, API_SECRET)

	// GetTradePairs
	prices, _ := binance.GetPrices()
	fmt.Println(len(prices))
	fmt.Println(*prices[0])

	for i, _ := range prices {
		if prices[i].Symbol == "ETHBTC" {
			fmt.Println(prices[i].Price)
		}
		if strings.HasSuffix(prices[i].Symbol, "BTC") {
			shortSymbol := strings.TrimSuffix(prices[i].Symbol, "BTC")
			fmt.Println(shortSymbol)
		}

	}

}
