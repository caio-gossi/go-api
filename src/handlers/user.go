package handlers

import (
	"encoding/json"
	"net/http"

	_ "example.com/go-api/docs"
	"example.com/go-api/domain"
)

// @Summary Test endpoint
// @Description Returns a message
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} domain.GenericMessage
// @Failure 404 {object} domain.GenericMessage
// @Router /test-message [get]
func UserHandler(w http.ResponseWriter, r *http.Request) {
	var response domain.GenericMessage = domain.GenericMessage{Message: "Hello world!"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
