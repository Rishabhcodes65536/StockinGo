package models

type Stock struct {
	Symbol string  `bson:"symbol" json:"symbol"`
	Name   string  `bson:"name" json:"name"`
	Price  float64 `bson:"price" json:"price"`
	UserID string  `bson:"user_id" json:"user_id"`
}
