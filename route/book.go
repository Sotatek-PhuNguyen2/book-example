package routes

import (
	"CRUD/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Index() http.Handler {
	routes := mux.NewRouter()
	routes.HandleFunc("/api/book", controller.Index).Methods("GET")
	routes.HandleFunc("/api/book/create", controller.CreateBook).Methods("POST")
	routes.HandleFunc("/api/book/{id}", controller.EditBook).Methods("PUT")
	routes.HandleFunc("/api/book/{id}", controller.DeleteBook).Methods("DELETE")
	routes.HandleFunc("/api/book/{id}", controller.Show).Methods("GET")
	return routes
}
