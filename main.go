package main

import (
	routes "CRUD/route"
	"fmt"
	"log"
	"net/http"
)

// Person represents a person document in MongoDB

func main() {
	router := routes.Index()
	fmt.Println("Server started with port 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
