FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

# installing CompileDaemon for hot-reload
RUN go get github.com/githubnemo/CompileDaemon

# --build - what will CompileDaemon run, after it detects file change. --command - executes after build (executes made main file) 
ENTRYPOINT CompileDaemon --build="go build cmd/app/main.go" --command=./main