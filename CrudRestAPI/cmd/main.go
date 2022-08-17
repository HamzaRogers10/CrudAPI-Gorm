package main

import (
	"CrudRestAPI/database"
	"CrudRestAPI/routers"
	"fmt"
)

func main() {
	fmt.Println("Welcome Crud API ")
	fmt.Println("routers is working ")

	database.Migration()
	routers.LoadRoutes()

}
 