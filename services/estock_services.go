package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"app-stock/internal/database"
)

type AlphaVantageResponse struct {
	GlobalQuote map[string]string `json:"Global Quote"`
}

func FetchAndSaveStock(symbol string) error {

	url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=W2ITEFIS1SAR401P"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data AlphaVantageResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	quote := data.GlobalQuote
	if quote == nil {
		return errors.New("no data from API")
	}

	price, _ := strconv.ParseFloat(quote["05. price"], 64)
	changePercent, _ := strconv.ParseFloat(
		quote["10. change percent"][:len(quote["10. change percent"])-1], 64,
	)
	volume, _ := strconv.ParseInt(quote["06. volume"], 10, 64)

	_, err = database.DB.Exec(
		`INSERT INTO stocks (symbol, price, change_percent, volume, source)
		 VALUES (?, ?, ?, ?, ?)`,
		symbol, price, changePercent, volume, "AlphaVantage",
	)

	return err
}
