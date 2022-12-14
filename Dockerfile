FROM golang:1.19.0

WORKDIR /usr/src/app

# Enable hot reload for Go
RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy