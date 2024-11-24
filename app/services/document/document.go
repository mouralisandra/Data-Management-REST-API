package document

import (
	"github.com/gorilla/mux"
)


// Function to initialize the Document service (optional)
func Init() {
	// Pre-populating with some data
	documents = append(documents, Document{ID: "1", Title: "Doc1", Author: "Sandra", Content: "Content of document 1"})
	documents = append(documents, Document{ID: "2", Title: "Doc2", Author: "Nada", Content: "Content of document 2"})
}

// Function to setup routes for the Document service
func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/documents", GetDocuments).Methods("GET")
	r.HandleFunc("/documents/{id}", GetDocument).Methods("GET")
	r.HandleFunc("/documents", CreateDocument).Methods("POST")
	r.HandleFunc("/documents/{id}", UpdateDocument).Methods("PUT")
	r.HandleFunc("/documents/{id}", DeleteDocument).Methods("DELETE")
}
