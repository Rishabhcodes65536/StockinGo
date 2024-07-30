package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Rishabhcodes65536/StockinGo/models"
	utils "github.com/Rishabhcodes65536/StockinGo/utils"
	"github.com/markcheno/go-quote"
)

func GetCandleStickChart(w http.ResponseWriter, r *http.Request) {
	var stockRequest models.StockRequest

	err := json.NewDecoder(r.Body).Decode(&stockRequest)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	stock, err1 := quote.NewQuoteFromYahoo(stockRequest.Symbol, "2023-01-01","2024-01-01", quote.Daily, true)

	if err1 != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	var candlesticks []models.Candlestick 
	for i,_ := range stock.Date {
		candlestick := models.Candlestick{
			Date: stock.Date[i].Format("2006-1-02"),
			Open: stock.Open[i],
			High: stock.High[i],
			Low: stock.Low[i],
			Close: stock.Close[i],
			Volume: int(stock.Volume[i]),
		}
		candlesticks=append(candlesticks, candlestick)
	}

	utils.RespondWithJSON(w, http.StatusOK, candlesticks)
}