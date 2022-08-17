package handlers

import (
	"CrudRestAPI/database"
	"CrudRestAPI/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ModifyUser(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	database.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}
