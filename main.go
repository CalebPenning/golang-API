package main

import (
	"encoding/json"
	"log"	
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book Struct

type book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *author `json:"author"`
}

// Author struct

type author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init Books variable as a slice Book struct
var books []book

// Get All Books

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get one single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&book{})
}

// Create book instance
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // NOT SAFE METHOD FOR GENERATING RANDOM ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update existing book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete existing book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	// Create Mock Data
	books = append(books, 
		book{ID: "1", Isbn: "aaa-1111", Title: "Tester Book", Author: &author{Firstname: "John", Lastname: "Doe"}})

	books = append(books, 
		book{ID: "2", Isbn: "aaa-1112", Title: "Tester Book: The Sequel", Author: &author{Firstname: "John", Lastname: "Toe"}})

	// Route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}