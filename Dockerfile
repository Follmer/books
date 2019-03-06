
# Dockerfile References: https://docs.docker.com/engine/reference/builder/


RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql

FROM mysql

ENV MYSQL_ROOT_PASSWORD password

# Add db
ENV MYSQL_DATABASE books

# Add the sql dir
COPY ./sql/ /docker-entrypoint-initdb.d/

# Start from golang latest base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Kyle Follmer <follmerka@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/books

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["go-docker"]
