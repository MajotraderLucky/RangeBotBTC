package main

import (
	"github.com/MajotraderLucky/MarketRepository/connect"
)

func main() {
	connect.Init()
	connect.GetApi()
}
