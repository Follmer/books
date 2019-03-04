package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    // "fmt"
    // "net/http/httputil"
)

type Book struct {
	ID 			int		`json:id,omitempty`
	Title		string	`json:title,omitempty`
	Author		string	`json:author,omitempty`
	Publisher	string	`json:publisher,omitempty`
	PublishDate	string	`json:publishdate,omitempty`
	Rating		int64		`json:rating,omitempty`
	Status		bool	`json:status,omitempty`
}

var books []Book

// Handle endpoints
func main() {

	books = append(books, Book{ID: 1, Title: "Harry Potter and the Sorcerer's Stone", Author: "J.K. Rowling", Publisher: "Bloomsbury", PublishDate: "06/26/1997", Rating: 3, Status: false})
	books = append(books, Book{ID: 2, Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Publisher: "Allen & Unwin", PublishDate: "07/29/1954", Rating: 2, Status: false})
	books = append(books, Book{ID: 3, Title: "The Cat in the Hat", Author: "Dr. Seuss", Publisher: "Random House", PublishDate: "03/12/1957", Rating: 1, Status: true})

	router := mux.NewRouter()

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")
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
		print(params["id"])
		if book.ID == 3 {
			json.NewEncoder(w).Encode(book)
		}
	}
}

// Creates a book with id
func CreateBook(w http.ResponseWriter, r *http.Request) {
	

	if r.Body == nil {
		http.Error(w, "Please send in a request body", 400)
		return
	}

	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// dump, _ := httputil.DumpRequest(r, true)


	print(r.FormValue("title"))

	if 	(book.ID != 0 && 
		book.Title != "" && 
		book.Author != "" &&
		book.Publisher != "" &&
		book.PublishDate != "" &&
		book.Rating != 0) {
		// book.Status != nil) {

		// book.ID, _ = params["id"]
		print(params["id"])
		book.Title = params["title"]
		book.Author = params["author"]
		book.Publisher = params["publisher"]
		book.PublishDate = params["publishdate"]
		book.Rating, _ = strconv.ParseInt(params["rating"], 0, 64)
		book.Status, _ = strconv.ParseBool(params["status"])	

		books = append(books, book)
		json.NewEncoder(w).Encode(books)

	} else {
		log.Fatal("Missing body param")
	}
	
	
}

// Deletes a book, by id
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	for index, book := range books {
		if book.ID == 3 {
			books = append(books[:index], books[:index+1]...)
            break
		}
		json.NewEncoder(w).Encode(books)
	}
}