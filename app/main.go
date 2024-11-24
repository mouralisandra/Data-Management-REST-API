package main

import (
	"log"
	"net/http"
	"app/services/customerdata"
	"app/services/document"
	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Initialize and setup routes for both services
	data.Init()
	document.Init()

	// Setup routes for each service
	data.SetupRoutes(r)
	document.SetupRoutes(r)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
