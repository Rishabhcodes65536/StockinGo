package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
    Database struct {
        User     string
        Password string
        DBName   string
    }
    JWTSecret string
}

var C Config

func LoadConfig() {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()
    viper.SetConfigType("yaml")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }

    err := viper.Unmarshal(&C)
    if err != nil {
        log.Fatalf("Unable to decode into struct, %v", err)
    }
	jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret != "" {
        C.JWTSecret = jwtSecret
    }
}
