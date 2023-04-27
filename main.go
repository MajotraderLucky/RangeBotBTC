package main

import (
	"time"

	"github.com/MajotraderLucky/MarketRepository/account"
	"github.com/MajotraderLucky/MarketRepository/connect"
	"github.com/MajotraderLucky/MarketRepository/priceanalyst"
)

func main() {
	connect.Init()
	priceanalyst.Hello()

	for range time.Tick(time.Second * 30) {
		account.Account()
	}
}
