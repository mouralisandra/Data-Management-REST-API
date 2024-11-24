package main

import (
	"log"
	"net/http"
	"DocumentsService/document"
	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	document.Init()
	document.SetupRoutes(r)

	// Start server
	log.Fatal(http.ListenAndServe(":7000", r))
}
