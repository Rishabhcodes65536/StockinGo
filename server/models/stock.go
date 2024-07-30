package models

type Candlestick struct {
    Date   string  `json:"date"`
    Open   float64 `json:"open"`
    High   float64 `json:"high"`
    Low    float64 `json:"low"`
    Close  float64 `json:"close"`
    Volume int     `json:"volume"`
}

type StockRequest struct {
    Symbol string `json:"symbol"`
}
