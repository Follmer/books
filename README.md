# Books

## Getting started

Begin by cloning the repository by using `git clone https://github.com/Follmer/books.git` in a terminal.

##### Run the container
`docker-compose up --build`

##### Run the tests
`go test -v`

##### Expected output
![](https://i.imgur.com/GvJYa8l.png)

##### Optional
If you'd like, you can run `go build` to create the executable.

Run `./books` to run the executable. You can then visit `http://localhost:8080/books` to perform any other CRUD operations you'd like.

An easier way to view/modify these operations is to use [Postman](https://www.getpostman.com/).
