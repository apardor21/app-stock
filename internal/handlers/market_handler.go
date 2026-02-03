package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Market struct {
	MarketType       string `json:"market_type"`
	Region           string `json:"region"`
	PrimaryExchanges string `json:"primary_exchanges"`
	CurrentStatus    string `json:"current_status"`
	LocalOpen        string `json:"local_open"`
	LocalClose       string `json:"local_close"`
}

type MarketStatusResponse struct {
	Markets []Market `json:"markets"`
}

func GetMarketStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := "https://www.alphavantage.co/query?function=MARKET_STATUS&apikey=demo"

	resp, err := http.Get(url)
	if err != nil || resp == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Alpha Vantage unavailable",
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to read response",
		})
		return
	}

	// 🔒 Rate limit or text response
	if strings.Contains(string(body), "rate limit") ||
		strings.Contains(string(body), "premium") {

		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "API rate limit exceeded",
		})
		return
	}

	var response MarketStatusResponse
	if err := json.Unmarshal(body, &response); err != nil {

		// 👇 RESPUESTA SEGURA SI ALPHA MANDA BASURA
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]Market{})
		return
	}

	// 👇 SI markets ES NIL, DEVOLVEMOS ARRAY VACÍO
	if response.Markets == nil {
		response.Markets = []Market{}
	}

	json.NewEncoder(w).Encode(response.Markets)
}
