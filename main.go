package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"bitbucket.org/d3vkk/classicbooks-api/devt/jsondata"
)

//The structs or classes below map the JSON data

// Books struct
type Books struct {
	Books []Book `json:"classicbooks"`
}

// Book struct
type Book struct {
	ID       int    `json:"id"`
	Isbn     int    `json:"isbn"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Synopsis string `json:"synopsis"`
}

// Note that: books.Books = []Book
var books Books

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books.Books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for j := 0; j < len(books.Books); j++ {
		if strconv.Itoa(books.Books[j].ID) == params["id"] {
			json.NewEncoder(w).Encode(books.Books[j])
			return
		}
	}
	json.NewEncoder(w).Encode(books.Books)
}

func main() {
	//Load Book Data
	var data = jsondata.BookData()

	//Parse JSON
	err := json.Unmarshal([]byte(data), &books)
	//Handle any errors
	if err != nil {
		fmt.Println("problem:", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
