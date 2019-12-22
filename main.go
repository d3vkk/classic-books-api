package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bitbucket.org/d3vkk/classicbooks-api/devt/jsontesting"
)

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

var books Books

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	var data = jsontesting.BookData()

	err := json.Unmarshal([]byte(data), &books)
	if err != nil {
		fmt.Println("problem:", err)
	}

	// fmt.Println(books.Books[0].Author)

	// for i := 0; i < len(books.Books); i++ {
	// 	fmt.Println("ID: " + strconv.Itoa(books.Books[i].ID))
	// 	fmt.Println("Isbn: " + strconv.Itoa(books.Books[i].Isbn))
	// 	fmt.Println("Title: " + books.Books[i].Title)
	// 	fmt.Println("Author: " + books.Books[i].Author)
	// 	fmt.Println("Synopsis: " + books.Books[i].Synopsis)
	// }

	router := mux.NewRouter()

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
