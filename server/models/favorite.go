package models

type Favorite struct {
    ID          int    `json:"id"`
    UserID      int    `json:"user_id"`
    StockSymbol string `json:"stock_symbol"`
}
