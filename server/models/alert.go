package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AlertType string

const (
	SignificantChange AlertType = "significant_change"
	Weekly            AlertType = "weekly"
	MarketStatus      AlertType = "market_status"
)

// type Alert struct {
// 	ID        string    `bson:"_id,omitempty"`
// 	UserID    string    `bson:"user_id"`
// 	Type      AlertType `bson:"type"`
// 	Symbol    string    `bson:"symbol"`
// 	Threshold float64   `bson:"threshold,omitempty"`
// 	CreatedAt time.Time `bson:"created_at"`
// }

type Alert struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	TargetPrice float64             `bson:"targetPrice"`
	UserEmail   string              `bson:"userEmail"`
	UpdatedAt   time.Time           `bson:"updatedAt"`
}

