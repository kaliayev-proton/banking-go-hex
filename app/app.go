package main

import (
	"net/http"
	"log"
)


func Start() {
		// Define routes
		http.HandleFunc("/greet", greet)
		http.HandleFunc("/customers", getAllCustomers)
	
		// Starting server
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
}