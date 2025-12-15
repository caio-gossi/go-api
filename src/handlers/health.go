package handlers

import "net/http"

// @Summary Health endpoint
// @Description Checks if the API is working
// @Tags health
// @Accept json
// @Produce json
// @Router /health [get]
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
