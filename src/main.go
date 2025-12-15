package main

import (
	"log"
	"net/http"

	"example.com/go-api/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/test-message", handlers.UserHandler)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Println("Starting server on :8080")

	var err error = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
