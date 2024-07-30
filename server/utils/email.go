package utils

import (
	"net/smtp"
	"os"

	"github.com/Rishabhcodes65536/StockinGo/errors"
	"github.com/joho/godotenv"
)

func SendEmail(to, subject, body string) {
	err := godotenv.Load()
	errors.HandleErr(err)
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	msg := "From: "+ from + "\n"+ "To: "+ to + "\n" + "Subject: " + subject + "\n\n" + body 

	err1:= smtp.SendMail("smtp.gmail.com:587",smtp.PlainAuth("",from,password,"smtp.gmail.com"),from,[]string{to}, []byte(msg))

	errors.HandleErr(err1)
	
}