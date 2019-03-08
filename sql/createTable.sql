CREATE DATABASE IF NOT EXISTS booksdb;
use booksdb;
CREATE TABLE IF NOT EXISTS books
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) DEFAULT NULL,
    author VARCHAR(100) DEFAULT NULL,
    publisher VARCHAR(100) DEFAULT NULL,
    publishdate VARCHAR(100) DEFAULT NULL,
    rating INT DEFAULT 0,
    status INT DEFAULT 0
);

INSERT INTO books (title, author, publisher, publishdate, rating, status)
VALUES("Harry Potter and the Sorcerer's Stone", "J.K. Rowling", "Bloomsbury", "06/26/1997", 3, false)

INSERT INTO books (title, author, publisher, publishdate, rating, status)
VALUES("The Lord of the Rings", "J.R.R. Tolkien", "Allen & Unwin", "07/29/1954", 1, true)

INSERT INTO books (title, author, publisher, publishdate, rating, status)
VALUES("The Cat in the Hat", "Dr. Seuss", "Random House", "03/12/1957", 2, false)
