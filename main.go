package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Data struct (Model)
type Data struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init datas var as a slice Data struct
var datas []Data

// Get all datas
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datas)
}

// Get single data
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through datas and find one with the id from the params
	for _, item := range datas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Data{})
}

// Add new data
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	data.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	datas = append(datas, data)
	json.NewEncoder(w).Encode(data)
}

// Update data
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range datas {
		if item.ID == params["id"] {
			datas = append(datas[:index], datas[index+1:]...)
			var data Data
			_ = json.NewDecoder(r.Body).Decode(&data)
			data.ID = params["id"]
			datas = append(datas, data)
			json.NewEncoder(w).Encode(data)
			return
		}
	}
}

// Delete data
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range datas {
		if item.ID == params["id"] {
			datas = append(datas[:index], datas[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(datas)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	datas = append(datas, Data{ID: "1", Isbn: "438227", Title: "Data One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	datas = append(datas, Data{ID: "2", Isbn: "454555", Title: "Data Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/datas", getBooks).Methods("GET")
	r.HandleFunc("/datas/{id}", getBook).Methods("GET")
	r.HandleFunc("/datas", createBook).Methods("POST")
	r.HandleFunc("/datas/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/datas/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Request sample
// {
// 	"isbn":"4545454",
// 	"title":"Data Three",
// 	"author":{"firstname":"Harry","lastname":"White"}
// }
