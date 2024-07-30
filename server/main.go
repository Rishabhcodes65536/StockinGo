package main

import (
	"net/http"

	"github.com/Rishabhcodes65536/StockinGo/cron"
	database "github.com/Rishabhcodes65536/StockinGo/database"
	"github.com/Rishabhcodes65536/StockinGo/middleware"
	"github.com/gorilla/mux"
	handlers "github.com/Rishabhcodes65536/StockinGo/handlers"
	errors "github.com/Rishabhcodes65536/StockinGo/errors"
)

func main() {
	database.Connect()

	r:=mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)


	r.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	r.HandleFunc("/favorites", middleware.AuthMiddleware(handlers.AddFavorite)).Methods("POST")
	r.HandleFunc("/favorites", middleware.AuthMiddleware(handlers.GetFavorite)).Methods("GET")

	go cron.CheckStockPrices()

	err :=http.ListenAndServe(":9752", r)

	errors.HandleErr(err)
	



}