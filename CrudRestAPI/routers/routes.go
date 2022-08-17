package routers

import (
	"CrudRestAPI/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func LoadRoutes()  {
	r := mux.NewRouter()

	r.HandleFunc("/viewsUserData", handlers.ViewsUsers).Methods("GET")
	r.HandleFunc("/viewUserData/{ id }", handlers.ViewUser).Methods("GET")
	r.HandleFunc("/signup", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/updaters/{id}", handlers.ModifyUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.RemoveUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8085", r))
}
