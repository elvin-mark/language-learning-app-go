package handlers

import (
	"encoding/json"
	dto "language-learning-app/dto/agents"
	"language-learning-app/services"
	"language-learning-app/utils"
	"net/http"
)

// ============== STRUCTS ==============

type AgentHandler struct {
	service services.AgentService
}

func NewAgentHandler(service services.AgentService) *AgentHandler {
	return &AgentHandler{service: service}
}

// ============== METHODS ==============

// GenerateLessonHandler godoc
// @Summary Get Lesson
// @Description Get a new Language Lesson based on the user current status
// @Tags agents
// @Accept json
// @Produce json
// @Param user body dto.GenerateLessonRequest true "Lesson Request object to be generated"
// @Success 200 {array} storage.Lesson
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /agents/lessons [post]
func (h *AgentHandler) GenerateLessonHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GenerateLessonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}
	lesson, err := h.service.GenerateLesson(req)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to get users"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, lesson)
}

// GenerateExerciseHandler godoc
// @Summary Get Practice Exercises
// @Description Get a new Exercise based on the input lesson
// @Tags agents
// @Accept json
// @Produce json
// @Param user body dto.GeneraterExerciseRequest true "Exercise Request object to be generated"
// @Success 200 {array} storage.Exercise
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /agents/exercises [post]
func (h *AgentHandler) GenerateExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GeneraterExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}
	exercise, err := h.service.GenerateExercise(req)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to get users"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, exercise)
}
