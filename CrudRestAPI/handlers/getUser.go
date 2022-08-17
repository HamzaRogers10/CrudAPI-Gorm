package handlers

import (
	"CrudRestAPI/database"
	"CrudRestAPI/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ViewUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	var user models.User
	database.DB.First(&user, params)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
