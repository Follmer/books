// main_test.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("admin", "admin", "books")
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM books WHERE 1=1")
	a.DB.Exec("ALTER TABLE books AUTO_INCREMENT = 1")
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/books", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentBook(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/book/45", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "Book not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Book not found'. Got '%s'", m["error"])
	}
}

func TestCreateBook(t *testing.T) {
	clearTable()
	payload := []byte(`{"title": "test book", "author": "bob", "publisher": "jerry", "publishdate": "04/10/2010", "rating": 3, "status": false}`)
	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "test book" {
		t.Errorf("Expected book title to be 'test book'. Got '%v'", m["title"])
	}

	if m["author"] != "bob" {
		t.Errorf("Expected book author to be 'bob'. Got '%v'", m["author"])
	}

	if m["publisher"] != "jerry" {
		t.Errorf("Expected book publisher to be 'jerry'. Got '%v'", m["publisher"])
	}

	if m["publishdate"] != "04/10/2010" {
		t.Errorf("Expected book publishdate to be '04/10/2010'. Got '%v'", m["publishdate"])
	}

	if m["rating"] != 3.0 {
		t.Errorf("Expected book rating to be '3'. Got '%v'", m["age"])
	}

	if m["status"] != false {
		t.Errorf("Expected book status to be 'false'. Got '%v'", m["status"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected book ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetBook(t *testing.T) {
	clearTable()
	addBooks(1)
	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateBook(t *testing.T) {
	clearTable()
	addBooks(1)
	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)
	var originalBook map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalBook)
	payload := []byte(`{"title": "test book updated", "rating": 2}`)
	req, _ = http.NewRequest("PUT", "/book/1", bytes.NewBuffer(payload))
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalBook["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalBook["id"], m["id"])
	}

	if m["title"] == originalBook["title"] {
		t.Errorf("Expected the title to change from '%v' to '%v'. Got '%v'", originalBook["title"], m["title"], m["title"])
	}
}

func TestDeleteBook(t *testing.T) {
	clearTable()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/book/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/book/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func addBooks(count int) {
	if count < 1 {
		count = 1
	}
	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO books(title, rating) VALUES('%s', %d)", ("Book " + strconv.Itoa(i+1)), rand.Intn(3-1)+1)
		a.DB.Exec(statement)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS books
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) DEFAULT NULL,
    author VARCHAR(100) DEFAULT NULL,
    publisher VARCHAR(100) DEFAULT NULL,
    publishdate VARCHAR(100) DEFAULT NULL,
    rating INT DEFAULT 0,
    status INT DEFAULT FALSE
)`
