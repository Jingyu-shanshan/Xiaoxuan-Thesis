FROM golang:1.19.0

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
COPY .env /usr/src/app

RUN go mod tidy