package main

import (
	"backend/database"
	"backend/routes"
	"fmt"
	"net/http"
)

func main() {
	db := database.ConnectDb()

	server := http.NewServeMux()
	routes.MapRoutes(server, db)

	fmt.Println("Databasae Connected")
	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", server)
}