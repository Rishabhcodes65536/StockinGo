package middleware

import (
	"net/http"

	"github.com/Rishabhcodes65536/StockinGo/pkg/utils"
)

func AuthMiddleware(next http.HandlerFunc) 	http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request) {
		token :=r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return 
		}

		_,err := utils.VerifyJWT(token)
		if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

		next(w,r)

	}
}