package handlers

import (
	"CrudRestAPI/database"
	"CrudRestAPI/models"
	"encoding/json"
	"net/http"
)

func ViewsUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}
