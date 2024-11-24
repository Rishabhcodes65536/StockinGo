package repository

import (
	"context"
	"fmt"

	"github.com/Rishabhcodes65536/StockinGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StockRepository struct {
	db *mongo.Collection
}

func NewStockRepository(db *mongo.Collection) *StockRepository {
	return &StockRepository{db: db}
}

// FindBySymbol finds a stock by its symbol.
func (r *StockRepository) FindBySymbol(symbol string) (*models.Stock, error) {
	var stock models.Stock
	filter := bson.M{"symbol": symbol}
	err := r.db.FindOne(context.Background(), filter).Decode(&stock)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("stock with symbol %s not found", symbol)
		}
		return nil, fmt.Errorf("failed to find stock: %w", err)
	}
	return &stock, nil
}

// GetFavoritesByUser retrieves the list of favorite stocks for a user.
func (r *StockRepository) GetFavoritesByUser(userEmail string) ([]models.Stock, error) {
	var stocks []models.Stock
	filter := bson.M{"favorites": userEmail}
	cursor, err := r.db.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch favorites: %w", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var stock models.Stock
		if err := cursor.Decode(&stock); err != nil {
			return nil, fmt.Errorf("failed to decode stock: %w", err)
		}
		stocks = append(stocks, stock)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}
	return stocks, nil
}

// AddFavorite adds a stock to the user's list of favorites.
func (r *StockRepository) AddFavorite(ctx context.Context, userEmail string, stock models.Stock) error {
	filter := bson.M{"symbol": stock.Symbol}
	update := bson.M{"$addToSet": bson.M{"favorites": userEmail}}
	_, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to add favorite: %w", err)
	}
	return nil
}

// RemoveFavorite removes a stock from the user's list of favorites.
func (r *StockRepository) RemoveFavorite(ctx context.Context, userEmail, stockID string) error {
	filter := bson.M{"_id": stockID}
	update := bson.M{"$pull": bson.M{"favorites": userEmail}}
	_, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to remove favorite: %w", err)
	}
	return nil
}
