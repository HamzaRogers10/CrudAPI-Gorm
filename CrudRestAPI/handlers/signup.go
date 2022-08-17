package handlers

import (
	"CrudRestAPI/database"
	"CrudRestAPI/models"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}
	database.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
