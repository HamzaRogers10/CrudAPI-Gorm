package handlers

import (
	"CrudRestAPI/database"
	"CrudRestAPI/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	database.DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}
