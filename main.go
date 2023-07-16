package main

import (
	"fmt"

	"github.com/MajotraderLucky/MarketRepository/account"
	"github.com/MajotraderLucky/MarketRepository/checkstartposition"
	"github.com/MajotraderLucky/MarketRepository/connect"
	"github.com/MajotraderLucky/MarketRepository/dataprocessing"
	"github.com/MajotraderLucky/MarketRepository/findtheorderlevel"
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

	chBool2 := make(chan bool)

	priceinsidethegridfibo.Hello()
	go priceinsidethegridfibo.Priceingrid(chBool2)

	setPriceBetweenVar := <-chBool2
	fmt.Println("Condition price between 236 and 786 -", setPriceBetweenVar)

	setOrder := setStartPosition && setPriceBetweenVar
	fmt.Println("Can place buy order -", setOrder)
	fmt.Println("----------------------")

	findtheorderlevel.Hello()
	findtheorderlevel.FindOrderLevel()
}
