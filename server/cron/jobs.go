package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/Rishabhcodes65536/StockinGo/database"
	"github.com/Rishabhcodes65536/StockinGo/errors"
	"github.com/Rishabhcodes65536/StockinGo/utils"
	"github.com/markcheno/go-quote"
)

func CheckStockPrices() {
	for {
		rows, err := database.DB.Query("SELECT user_id, stock_symbol FROM favorites")
		errors.HandleErr(err)
		defer rows.Close()

		favorites := make(map[int][]string)

		for rows.Next(){
			var userId int 
			var stockSymbol string 
			err := rows.Scan(&userId, &stockSymbol)
			errors.HandleErr(err)

			favorites[userId]= append(favorites[userId], stockSymbol)
		}

		for userIds,symbols := range favorites{
			for _, symbol :=range symbols{
				stock, err :=quote.NewQuoteFromYahoo(symbol,"2023-01-01","2024-01-01",quote.Daily, true)
				errors.HandleErr(err)

				if len(stock.Close)>1 {
					change := (stock.Close[len(stock.Close)-1] - stock.Close[len(stock.Close)-2]) / stock.Close[len(stock.Close)-2]
                    if change > 0.05 || change < -0.05 {  // essentially mai absolute 5% aur badhkar tak allow kar raha hu
                        // Fetch user email
                        var email string
                        err := database.DB.QueryRow("SELECT email FROM users WHERE id = $1", userIds).Scan(&email)
                        if err != nil {
                            log.Println("Error fetching user email:", err)
                            continue
                        }

                        // Send notification email
                        utils.SendEmail(email, "Stock Price Alert", "The stock "+symbol+" has changed significantly by "+fmt.Sprintf("%.2f", change*100)+"%.")
                    }
				}
			}
		}

		time.Sleep(24* time.Hour)
	}
}