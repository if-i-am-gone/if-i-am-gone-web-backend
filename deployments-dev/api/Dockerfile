FROM golang:latest

WORKDIR /go/src/ifidie_backend

COPY . /go/src/ifidie_backend

RUN go mod download

RUN go mod tidy

RUN go mod verify

# installing CompileDaemon for hot-reload
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

# --build - what will CompileDaemon run, after it detects file change. --command - executes after build (executes made main file) 
# ENTRYPOINT CompileDaemon --build="go build /app/cmd/app/main.go" --command=./main
ENTRYPOINT CompileDaemon --build="go build /go/src/ifidie_backend/cmd/app/main.go" --command="./main"