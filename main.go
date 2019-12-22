package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"./jsontesting"
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

func main() {
	var books Books
	var data = jsontesting.BookData()

	err := json.Unmarshal([]byte(data), &books)
	if err != nil {
		fmt.Println("problem:", err)
	}

	fmt.Println(books.Books[0].Author)

	for i := 0; i < len(books.Books); i++ {
		fmt.Println("ID: " + strconv.Itoa(books.Books[i].ID))
		fmt.Println("Isbn: " + strconv.Itoa(books.Books[i].Isbn))
		fmt.Println("Title: " + books.Books[i].Title)
		fmt.Println("Author: " + books.Books[i].Author)
		fmt.Println("Synopsis: " + books.Books[i].Synopsis)
	}

}
