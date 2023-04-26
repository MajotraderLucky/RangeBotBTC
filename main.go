package main

import (
	"time"

	"github.com/MajotraderLucky/MarketRepository/account"
	"github.com/MajotraderLucky/MarketRepository/connect"
)

func main() {
	connect.Init()

	for range time.Tick(time.Second * 30) {
		account.Account()
	}
}
