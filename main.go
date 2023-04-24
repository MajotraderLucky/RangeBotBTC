package main

import (
	"github.com/MajotraderLucky/MarketRepository/account"
	"github.com/MajotraderLucky/MarketRepository/connect"
)

func main() {
	connect.Init()
	account.Account()
}
