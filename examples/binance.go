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

	// GetTicketDaily
	tickerDailys, _ := binance.GetTickerDaily()
	fmt.Println(*tickerDailys[0])

	for i, _ := range tickerDailys {
		//		fmt.Println(i, tickerDailys[i])
		fmt.Println(tickerDailys[i].Symbol, tickerDailys[i].BidPrice)
	}

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
			fmt.Println(prices[i].Price)
		}
	}

	// GetSymbols
	exchangeInfo, err := binance.GetExchangeInfo()
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
	fmt.Println(len(exchangeInfo.Symbols))
	fmt.Printf("%+v", exchangeInfo.Symbols)
	for i, _ := range exchangeInfo.Symbols {
		fmt.Printf("%s - %s\n", exchangeInfo.Symbols[i].BaseAsset, exchangeInfo.Symbols[i].Symbol)
	}

	fmt.Println("Symbols")

	for i, _ := range exchangeInfo.Symbols {
		if strings.HasSuffix(exchangeInfo.Symbols[i].Symbol, "BTC") {
			shortSymbol := strings.TrimSuffix(exchangeInfo.Symbols[i].Symbol, "BTC")
			fmt.Println(shortSymbol)
		}
	}

	// GetAssets
	assets, err := binance.GetAllAssets()
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
	fmt.Println(len(assets))
	fmt.Printf("%+v\n", *assets[0])
	for i, _ := range assets {
		fmt.Printf("%s - %s\n", assets[i].Code, assets[i].Name)
	}

}
