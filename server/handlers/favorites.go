// package handlers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/Rishabhcodes65536/StockinGo/database"
// 	"github.com/Rishabhcodes65536/StockinGo/models"
// 	"github.com/Rishabhcodes65536/StockinGo/pkg/utils"
// )

// func AddFavorite(w http.ResponseWriter,r *http.Request){
// 	var favorite models.Favorite
// 	err := json.NewDecoder(r.Body).Decode(&favorite)
// 	if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

// 	_,err = database.DB.Exec("INSERT INTO favorites (user_id, stock_symbol) VALUES ($1,$2)", favorite.UserID, favorite.StockSymbol)

// 	 if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }

//     utils.RespondWithJSON(w, http.StatusCreated, "Favorite added successfully")
// }

// func GetFavorite(w http.ResponseWriter, r *http.Request) {
//     userID := r.Context().Value("userID").(int)

//     rows, err := database.DB.Query("SELECT stock_symbol FROM favorites WHERE user_id = $1", userID)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     defer rows.Close()

//     var favorites []string
//     for rows.Next() {
//         var stockSymbol string
//         err := rows.Scan(&stockSymbol)
//         if err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//             return
//         }
//         favorites = append(favorites, stockSymbol)
//     }

//     utils.RespondWithJSON(w, http.StatusOK, favorites)
// }