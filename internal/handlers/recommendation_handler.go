package handlers

import (
	"encoding/json"
	"net/http"

	"app-stock/internal/database"
	"app-stock/models"
)

func GetRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	row := database.DB.QueryRow(`
		SELECT symbol, price, change_percent, volume, source, created_at
		FROM stocks
		ORDER BY price DESC, volume DESC
		LIMIT 1
	`)

	var stock models.Stock
	err := row.Scan(
		&stock.Symbol,
		&stock.Price,
		&stock.ChangePercent,
		&stock.Volume,
		&stock.Source,
		&stock.CreatedAt,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stock)
}
