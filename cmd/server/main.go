package main

import (
	"log"
	"net/http"

	"app-stock/internal/database"
	"app-stock/internal/handlers"
)

func main() {

	// 1️⃣ DB
	database.Connect()

	// 2️⃣ Rutas
	http.HandleFunc("/api/stocks", handlers.GetStocks)
	http.HandleFunc("/api/stocks/fetch", handlers.FetchAndSaveStock)
	http.HandleFunc("/api/market-status", handlers.GetMarketStatus)

	http.HandleFunc("/api/recommendation", handlers.GetRecommendation)
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API OK"))
	})

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
