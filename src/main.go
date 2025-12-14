package main

import (
	"log"
	"net/http"
)

func main() {

	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	log.Println("Starting server on :8080")

	var err error = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
