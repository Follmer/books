package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "fmt"
	// "net/http/httputil"
)

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

	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// dump, _ := httputil.DumpRequest(r, true)

	print(r.FormValue("title"))

	var missingAttributes []string

	if book.Title == "" {
		missingAttributes = append(missingAttributes, "Title")
	}
	if book.Author == "" {
		missingAttributes = append(missingAttributes, "Author")
	}
	if book.Publisher == "" {
		missingAttributes = append(missingAttributes, "Publisher")
	}
	if book.PublishDate == "" {
		missingAttributes = append(missingAttributes, "PublishDate")
	}
	if book.Rating == 0 {
		missingAttributes = append(missingAttributes, "Rating")
	}

	if len(missingAttributes) > 0 {
		for _, element := range missingAttributes {
			fmt.Println("Missing body attribute:", element)
		}
		return
	}

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
