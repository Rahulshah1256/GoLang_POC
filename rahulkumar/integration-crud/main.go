package main

import (
	"log"
	"net/http"

	"integration-crud/api"

	"integration-crud/db"
	"integration-crud/util"

	"github.com/gorilla/mux"
)

func main() {
	config := util.LoadConfig() // Implement LoadConfig() function to load configuration from a file or environment variables

	router := mux.NewRouter()

	store, err := db.NewStore(config) // Implement NewStore() function to create a database store
	if err != nil {
		log.Fatal(err)
	}

	apiHandler := api.NewHandler(store) // Implement NewHandler() function to create an API handler

	router.HandleFunc("/items", apiHandler.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", apiHandler.GetItem).Methods("GET")
	router.HandleFunc("/items", apiHandler.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", apiHandler.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", apiHandler.DeleteItem).Methods("DELETE")

	log.Println("Server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
