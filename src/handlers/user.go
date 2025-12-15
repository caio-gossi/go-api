package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/go-api/domain"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var response domain.GenericMessage = domain.GenericMessage{Message: "Hellow world"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
