package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Rishabhcodes65536/StockinGo/database"
	"github.com/Rishabhcodes65536/StockinGo/models"
	utils "github.com/Rishabhcodes65536/StockinGo/utils"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	hashedPassword,err1 := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err1!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	user.Password =string(hashedPassword)

	_,err =database.DB.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", user.Username,user.Password, user.Email)
	if err != nil {
		if err.(*sql.Error).Number == 1062  {
			http.Error(w, "Username already exists", http.StatusConflict)
		} else{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var storedUser models.User
    err = database.DB.QueryRow("SELECT id, username, password FROM users WHERE username=$1", user.Username).Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(storedUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}