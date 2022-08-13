FROM golang:1.17-alpine

WORKDIR /app

COPY .. /app

RUN apk add --update make
RUN apk add --update nodejs npm
RUN go get github.com/githubnemo/CompileDaemon
RUN go mod download


ENTRYPOINT CompileDaemon --build="go build -o main transactions/main.go" --command=./main