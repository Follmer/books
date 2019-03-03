package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Book struct {
	ID 			string	`json:id,omitempty`
	Title		string	`json:title,omitempty`
	Author		string	`json:author,omitempty`
	Publisher	string	`json:publisher,omitempty`
	PublishDate	string	`json:publishdate,omitempty`
	Rating		int		`json:rating,omitempty`
	Status		bool	`json:status,omitempty`
}

var books []Book

// Handle endpoints
func main() {

	books = append(books, Book{ID: "1", Title: "Harry Potter and the Sorcerer's Stone", Author: "J.K. Rowling", Publisher: "Bloomsbury", PublishDate: "06/26/1997", Rating: 3, Status: false})
	books = append(books, Book{ID: "2", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Publisher: "Allen & Unwin", PublishDate: "07/29/1954", Rating: 2, Status: false})
	books = append(books, Book{ID: "3", Title: "The Cat in the Hat", Author: "Dr. Seuss", Publisher: "Random House", PublishDate: "03/12/1957", Rating: 1, Status: true})

	router := mux.NewRouter()

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Gets all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

// Gets a specific book, by id
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
		}
	}
}

// Creates a book with id
func CreateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// Deletes a book, by id
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[:index+1]...)
            break
		}
		json.NewEncoder(w).Encode(books)
	}
}