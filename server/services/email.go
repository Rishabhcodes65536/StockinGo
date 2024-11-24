package services

import (
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

type EmailService struct {
    config EmailConfig
}

type EmailConfig struct {
    From     string
    Password string
    Host     string
    Port     string
}

func NewEmailService() *EmailService {
    return &EmailService{
        config: EmailConfig{
            From:     os.Getenv("EMAIL_FROM"),
            Password: os.Getenv("EMAIL_PASSWORD"),
            Host:     os.Getenv("SMTP_HOST"),
            Port:     os.Getenv("SMTP_PORT"),
        },
    }
}

func (s *EmailService) SendEmail(to, subject, body string) error {
    e := email.NewEmail()
    e.From = s.config.From
    e.To = []string{to}
    e.Subject = subject
    e.HTML = []byte(body)

    addr := s.config.Host + ":" + s.config.Port
    auth := smtp.PlainAuth("", s.config.From, s.config.Password, s.config.Host)

    return e.Send(addr, auth)
}