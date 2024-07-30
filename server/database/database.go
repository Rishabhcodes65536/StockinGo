package database

import (
	"database/sql"
	"log"
	errors "github.com/Rishabhcodes65536/StockinGo/errors"
)

var DB *sql.DB



func Connect(){
	connStr := "user=Rishabh password=:Secret dbname=hello sllmode=disable"
	var err error
	DB,err := sql.Open("postgres",connStr)
	errors.HandleErr(err)
	err= DB.Ping()
	errors.HandleErr(err)
	log.Println("Database Connected Succesfully!!")
}