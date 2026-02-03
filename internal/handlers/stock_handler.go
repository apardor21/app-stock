package handlers

import (
	"encoding/json"
	"net/http"

	"app-stock/internal/database"
	"app-stock/models"
)

func GetStocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := database.DB.Query("SELECT symbol, price, change_percent, volume, source, created_at FROM stocks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var stocks []models.Stock

	for rows.Next() {
		var s models.Stock
		rows.Scan(&s.Symbol, &s.Price, &s.ChangePercent, &s.Volume, &s.Source, &s.CreatedAt)
		stocks = append(stocks, s)
	}

	json.NewEncoder(w).Encode(stocks)
}

func FetchAndSaveStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "symbol is required",
		})
		return
	}

	// lógica de fetch + insert...

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Stock actualizado correctamente",
		"symbol":  symbol,
	})
}
