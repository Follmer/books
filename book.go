package main

type Book struct {
	ID          int    `json:id,omitempty`
	Title       string `json:title,omitempty`
	Author      string `json:author,omitempty`
	Publisher   string `json:publisher,omitempty`
	PublishDate string `json:publishdate,omitempty`
	Rating      int64  `json:rating,omitempty`
	Status      bool   `json:status,omitempty`
}

var books []Book
