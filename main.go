package main

func main() {

	// books = append(books, Book{ID: 1, Title: "Harry Potter and the Sorcerer's Stone", Author: "J.K. Rowling", Publisher: "Bloomsbury", PublishDate: "06/26/1997", Rating: 3, Status: false})
	// books = append(books, Book{ID: 2, Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Publisher: "Allen & Unwin", PublishDate: "07/29/1954", Rating: 2, Status: false})
	// books = append(books, Book{ID: 3, Title: "The Cat in the Hat", Author: "Dr. Seuss", Publisher: "Random House", PublishDate: "03/12/1957", Rating: 1, Status: true})

	a := App{}
	a.Initialize("admin", "admin", "books")
	a.Run(":8080")
}
