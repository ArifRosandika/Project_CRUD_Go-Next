package routes

import (
	"backend/controllers"
	"database/sql"
	"net/http"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/products", controllers.HandleIndexController(db))
	server.HandleFunc("/products/create", controllers.HandleProductCreate(db))
	server.HandleFunc("/products/edit", controllers.HandleProductEdit(db))
	server.HandleFunc("/products/delete", controllers.HandleDeleteProduct(db))
}