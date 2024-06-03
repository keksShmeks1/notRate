package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const envPath = "/Users/keksShmeks/fun/notGetter/.env"

func ChangeEnvAmount(amount int64) error {
	strAmount := fmt.Sprint(amount)
	godotenv.Write(map[string]string{"NOTAMOUNT": strAmount}, envPath)
	return nil
}

func GetEnvAmount() (int64, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return 0, err
	}

	amountStr, f := os.LookupEnv("NOTAMOUNT")
	if !f {
		err := errors.New("can't get NOTAMOUNT from env")
		return 0, err
	}

	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return amount, nil
}
