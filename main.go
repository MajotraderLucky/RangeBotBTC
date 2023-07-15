package main

import (
	"fmt"

	"github.com/MajotraderLucky/MarketRepository/account"
	"github.com/MajotraderLucky/MarketRepository/checkstartposition"
	"github.com/MajotraderLucky/MarketRepository/connect"
	"github.com/MajotraderLucky/MarketRepository/dataprocessing"
	"github.com/MajotraderLucky/MarketRepository/priceanalyst"
	"github.com/MajotraderLucky/MarketRepository/priceinsidethegridfibo"
)

func main() {
	connect.Init()
	priceanalyst.FiboLongBtc()
	account.Account()
	ch := make(chan string)
	go dataprocessing.DataGenerator(ch)
	dataprocessing.SetTestVarString(ch)
	chBool := make(chan bool)
	go checkstartposition.Checkstartposition(chBool)

	setStartPosition := <-chBool
	fmt.Println("Result start position in main.go -", setStartPosition)

	priceinsidethegridfibo.Hello()
	priceinsidethegridfibo.Priceingrid()
}
