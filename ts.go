package main

import (
	"log"
	"net/http"
	"os"

	"hackathon/TS/handlers"
	"hackathon/TS/server"
)

func main() {
	logger := log.New(os.Stdout, "Trusting Social ", log.LstdFlags|log.Lshortfile)

	h := handlers.NewHandlers(logger)
	r := server.New()
	h.SetupRoutes(r)

	ServiceAddr := "5000"
	logger.Println("server listening on " + ServiceAddr)
	err := http.ListenAndServe(":"+ServiceAddr, r)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
