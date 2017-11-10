go-binance
==========

go-binance is an implementation of the Binance API (public and private) in Golang.

Based off of https://github.com/toorop/go-bittrex/

This library is more of a framework for some bots I use so it is expected that a lot of things don't work but pull requests are excepted.

## Import
	import "github.com/jyap808/go-binance"
	
## Usage
~~~ go
package main

import (
    "fmt"
    "github.com/jyap808/go-binance"
)

const (
    API_KEY    = "YOUR_API_KEY"
    API_SECRET = "YOUR_API_SECRET"
)

func main() {
    // Binance client
    binance := binance.New(API_KEY, API_SECRET)

    // Get tickers
    tickers, err := binance.GetTickers()
    fmt.Println(err, tickers)
}
~~~	

See ["Examples" folder for more... examples](https://github.com/jyap808/go-binance/blob/master/examples/binance.go)

## Stay tuned
[Follow me on Twitter](https://twitter.com/jyap)

