package client

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

// Name Notcoin
// url example https://api.swapzone.io/v1/exchange/get-rate?from=not&to=usdt&amount=40000
// from, to, amount
const (
	xApiKey    = "5gaCIZSRO"
	rubRateUrl = "https://www.cbr-xml-daily.ru/daily.xml"
)

func GetNotRate(amount int64) (float64, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.swapzone.io/v1/exchange/get-rate?from=not&to=usdt&amount=%d", amount)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("x-api-key", xApiKey)
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Парсинг JSON-ответа в карту
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	// Извлечение конкретного поля
	amountTo, ok := result["amountTo"].(float64)
	if !ok {
		return 0, err
	}

	rub, err := convertToRub(amountTo)
	if err != nil {
		return 0, err
	}

	return rub, nil
}

func convertToRub(usd float64) (float64, error) {
	resp, err := http.Get(rubRateUrl)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	// Парсинг XML
	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return 0, err
	}

	for _, valute := range valCurs.Valutes {
		if valute.ID == "R01235" {
			valute.Value = strings.ReplaceAll(valute.Value, ",", ".")

			value, err := strconv.ParseFloat(valute.Value, 64)
			if err != nil {
				return 0, err
			}
			return usd * value, nil
		}
	}
	err = errors.New("error finding RU rate")

	return 0, err
}
