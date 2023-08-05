package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/widjaya-hernando/omdb-service/omdb/server"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("NewServer failed with error: %s", err)
	}

	srv.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	signal := <-sigChan
	log.Printf("shutting down products server with signal: %s", signal)
}
