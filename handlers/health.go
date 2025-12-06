package handlers

import (
	"language-learning-app/utils"
	"net/http"
)

// HealthCheckHandler godoc
// @Summary Show the status of the server.
// @Description get the status of the server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, map[string]interface{}{
		"status":  "healthy",
		"server":  "Language Learning API",
		"version": "1.0.0",
	})
}
