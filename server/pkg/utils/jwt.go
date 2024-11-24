package utils

import (
	"errors"
	"time"

	"github.com/Rishabhcodes65536/StockinGo/config"
	"github.com/Rishabhcodes65536/StockinGo/models"
	"github.com/dgrijalva/jwt-go"
)


var jwtKey = []byte(config.C.JWTSecret)

type Claims struct {
	Userid   int    `json:"userid`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (string, error) {
	expirationTime := time.Now().Add(24*time.Hour)
	claims := &Claims{
		Userid:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString,err := token.SignedString(jwtKey)
	if err != nil {
        return "", err
    }
    return tokenString, nil
}

func VerifyJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{},error){
		return jwtKey,nil
	})
	if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, errors.New("invalid token")
    }
    return claims, nil
}