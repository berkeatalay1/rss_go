package main

import (
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	const port = "8080"
	godotenv.Load()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}
}
