package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/keksShmeks1/notGetter/pkg/client"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	amountStr, flag := os.LookupEnv("NOTAMOUNT")
	if !flag {
		log.Fatal("No field named NOTAMOUNT")
		return
	}

	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		log.Fatal("No field named NOTAMOUNT")
		return
	}

	rate, err := client.GetNotRate(amount)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Print(rate)

}
