package main

import (
	"database/sql"
	"errors"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"publishdate"`
	Rating      int64  `json:"rating"`
	Status      bool   `json:"status"`
}

func (b *Book) getBook(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (b *Book) updateBook(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (b *Book) deleteBook(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (b *Book) createBook(db *sql.DB) error {
	return errors.New("Not implemented")
}
func getBooks(db *sql.DB, start, count int) ([]Book, error) {
	return nil, errors.New("Not implemented")
}
