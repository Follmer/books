package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handle endpoints
func main() {

	a := App{}
	a.Initialize("admin", "password", "books")
	a.Run(":8080")

	books = append(books, Book{ID: 1, Title: "Harry Potter and the Sorcerer's Stone", Author: "J.K. Rowling", Publisher: "Bloomsbury", PublishDate: "06/26/1997", Rating: 3, Status: false})
	books = append(books, Book{ID: 2, Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Publisher: "Allen & Unwin", PublishDate: "07/29/1954", Rating: 2, Status: false})
	books = append(books, Book{ID: 3, Title: "The Cat in the Hat", Author: "Dr. Seuss", Publisher: "Random House", PublishDate: "03/12/1957", Rating: 1, Status: true})

	router := mux.NewRouter()

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	err := http.ListenAndServe(":8005", router)

	if err != nil {
		log.Println(err)
	}
}
