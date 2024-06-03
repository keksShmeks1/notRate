package main

import (
	"flag"
	"log"

	"github.com/keksShmeks1/notRate/pkg/app"
	"github.com/keksShmeks1/notRate/pkg/utils"
)

func main() {

	var amount int64

	var srv app.Server

	flag.Int64Var(&amount, "ch", 0, "change amount of not u have")

	flag.Parse()

	if amount > 0 {
		err := utils.ChangeEnvAmount(amount)
		if err != nil {
			log.Fatal("Can't change NOT amount")
		}
		log.Print("Not amount successfully changed")
		srv = *app.InitServer(amount)
	}

	envAmount, err := utils.GetEnvAmount()
	if err != nil {
		log.Fatal(err)
	}

	srv = *app.InitServer(envAmount)

	err = srv.GetRate()
	if err != nil {
		log.Fatal(err)
	}

}
