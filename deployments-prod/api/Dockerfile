FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

RUN go mod tidy

RUN go mod verify

ENTRYPOINT go run ./cmd/app/main.go