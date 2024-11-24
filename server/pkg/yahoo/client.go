package yahoo

import (
	"github.com/piquette/finance-go/quote"
)

type Client struct {}

func NewClient() *Client {
    return &Client{}
}

func (c *Client) GetQuote(symbol string) (*quote.Quote, error) {
    q, err := quote.Get(symbol)
    if err != nil {
        return nil, err
    }
    return q, nil
}