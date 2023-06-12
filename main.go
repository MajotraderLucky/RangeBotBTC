package main

import (
	"github.com/MajotraderLucky/MarketRepository/account"
	"github.com/MajotraderLucky/MarketRepository/connect"
	"github.com/MajotraderLucky/MarketRepository/dataprocessing"
	"github.com/MajotraderLucky/MarketRepository/priceanalyst"
)

func main() {
	connect.Init()
	priceanalyst.FiboLongBtc()
	account.Account()
	dataprocessing.TestHello()
	dataprocessing.DataGenerator()
}
