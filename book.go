package main

import (
	"database/sql"
	"fmt"
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

func (u *Book) getBook(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT title, rating FROM books WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Title, &u.Rating)
}
func (u *Book) updateBook(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE books SET title='%s', rating=%d WHERE id=%d", u.Title, u.Rating, u.ID)
	_, err := db.Exec(statement)
	return err
}
func (u *Book) deleteBook(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM books WHERE id=%d", u.ID)
	_, err := db.Exec(statement)
	return err
}
func (u *Book) createBook(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO books(title, rating) VALUES('%s', %d)", u.Title, u.Rating)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}
func getBooks(db *sql.DB, start, count int) ([]Book, error) {
	statement := fmt.Sprintf("SELECT id, title, rating FROM books LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := []Book{}
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Rating); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
