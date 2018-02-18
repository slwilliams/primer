package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	// defaultPort is the default HTTP server port
	defaultPort int = 8080

	// defaultMaxNumber to check for primality
	// increase for a longer test
	defaultMaxNumber = 50000000
)

var (
	logger   = log.New(os.Stdout, "[primer] ", log.Lshortfile|log.Ldate|log.Ltime)
	hostname = "undefined"
)

func main() {

	host, err := os.Hostname()
	if err == nil {
		hostname = host
	}

	logger.Printf("Starting server on port %d...", defaultPort)
	srv := startServer()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	logger.Printf("Shutting down server on port %d...", defaultPort)
	srv.Shutdown(context.Background())
	os.Exit(0)

}
