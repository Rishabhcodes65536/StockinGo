package repository

import (
	"context"
	"fmt"

	"github.com/Rishabhcodes65536/StockinGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlertRepository struct {
	db *mongo.Collection
}

func NewAlertRepository(db *mongo.Collection) *AlertRepository {
	return &AlertRepository{db: db}
}

// Create creates a new alert in the database.
func (r *AlertRepository) Create(alert *models.Alert) error {
	_, err := r.db.InsertOne(context.Background(), alert)
	if err != nil {
		return fmt.Errorf("failed to create alert: %w", err)
	}
	return nil
}

// Update updates an existing alert.
func (r *AlertRepository) Update(alertID string, alert *models.Alert) error {
	filter := bson.M{"_id": alertID}
	update := bson.M{"$set": alert}
	_, err := r.db.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update alert: %w", err)
	}
	return nil
}

// Delete removes an alert from the database.
func (r *AlertRepository) Delete(alertID string) error {
	filter := bson.M{"_id": alertID}
	_, err := r.db.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete alert: %w", err)
	}
	return nil
}

// GetAll retrieves all alerts from the database.
func (r *AlertRepository) GetAll() ([]models.Alert, error) {
	var alerts []models.Alert
	cursor, err := r.db.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch alerts: %w", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var alert models.Alert
		if err := cursor.Decode(&alert); err != nil {
			return nil, fmt.Errorf("failed to decode alert: %w", err)
		}
		alerts = append(alerts, alert)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}
	return alerts, nil
}
