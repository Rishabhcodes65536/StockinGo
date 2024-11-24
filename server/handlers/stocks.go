package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Rishabhcodes65536/StockinGo/services"
	"github.com/Rishabhcodes65536/StockinGo/models"
)

// SearchStock handler
func SearchStock(stockService services.StockService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Query().Get("symbol")
		stock, err := stockService.SearchStock(symbol)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stock)
	}
}

// GetFavorites handler
func GetFavorites(stockService services.StockService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		favorites, err := stockService.GetFavorites()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(favorites)
	}
}

// AddFavorite handler
func AddFavorite(stockService services.StockService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Query().Get("symbol")
		err := stockService.AddFavorite(symbol)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Stock added to favorites")
	}
}

// RemoveFavorite handler
func RemoveFavorite(stockService services.StockService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Query().Get("symbol")
		err := stockService.RemoveFavorite(symbol)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Stock removed from favorites")
	}
}
