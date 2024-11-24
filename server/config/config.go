package config

import "os"

type Config struct {
	MongoURI  string
	JWTSecret string
	Port      string
	Email     EmailConfig
}

type EmailConfig struct {
	From     string
	Password string
	Host     string
	Port     string
}

func Load() *Config {
	return &Config{
		MongoURI:  os.Getenv("MONGODB_URI"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		Port:      getEnvOrDefault("PORT", "8080"),
		Email: EmailConfig{
			From:     os.Getenv("EMAIL_FROM"),
			Password: os.Getenv("EMAIL_PASSWORD"),
			Host:     os.Getenv("SMTP_HOST"),
			Port:     os.Getenv("SMTP_PORT"),
		},
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}