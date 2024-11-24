package services

import (
	"context"
	"fmt"

	"github.com/Rishabhcodes65536/StockinGo/internal/repository"
	"github.com/Rishabhcodes65536/StockinGo/models"
)

type StockService struct {
	stockRepo repository.StockRepository
}

func NewStockService(stockRepo repository.StockRepository) *StockService {
	return &StockService{stockRepo: stockRepo}
}

// SearchStock searches for a stock by symbol and returns its details.
func (s *StockService) SearchStock(symbol string) (*models.Stock, error) {
	stock, err := s.stockRepo.FindBySymbol(symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to find stock: %w", err)
	}
	return stock, nil
}

// GetFavorites returns all favorite stocks.
func (s *StockService) GetFavorites(userEmail string) ([]models.Stock, error) {
	stocks, err := s.stockRepo.GetFavoritesByUser(userEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch favorites: %w", err)
	}
	return stocks, nil
}

// AddFavorite adds a stock to the user's favorites.
func (s *StockService) AddFavorite(userEmail string, symbol string) error {
	stock, err := s.stockRepo.FindBySymbol(symbol)
	if err != nil {
		return fmt.Errorf("failed to find stock for symbol %s: %w", symbol, err)
	}
	err = s.stockRepo.AddFavorite(context.Background(), userEmail, *stock)
	if err != nil {
		return fmt.Errorf("failed to add stock to favorites: %w", err)
	}
	return nil
}

// RemoveFavorite removes a stock from the user's favorites.
func (s *StockService) RemoveFavorite(userEmail string, symbol string) error {
	stock, err := s.stockRepo.FindBySymbol(symbol)
	if err != nil {
		return fmt.Errorf("failed to find stock for symbol %s: %w", symbol, err)
	}
	err = s.stockRepo.RemoveFavorite(context.Background(), userEmail, stock.ID.Hex())
	if err != nil {
		return fmt.Errorf("failed to remove stock from favorites: %w", err)
	}
	return nil
}
