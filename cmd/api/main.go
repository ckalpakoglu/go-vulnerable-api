package main

import (
	stdlog "log"

	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env if present
	_ = godotenv.Load()

	e := server.New()

	if err := server.Start(e, ":8080"); err != nil {
		stdlog.Fatalf("server error: %v", err)
	}
}
