package handlers

import (
	"encoding/json"
	dto "language-learning-app/dto/agents"
	"language-learning-app/services"
	"language-learning-app/utils"
	"net/http"
	"strconv"
)

type exerciseHandlerImpl struct {
	exerciseService services.ExerciseService
	userService     services.UserService
}

// GradeUsageHandler godoc
//
//	@Summary		Grade Usage of language
//	@Description	Grade the usage of the language in a sentence
//	@Tags			exercise
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.GradeUsageRequest	true	"Grade Usage Request"
//	@Success		200		{object}		agents.UsageGrade
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/resources/exercise/usage/grade [post]
func (h *exerciseHandlerImpl) GradeUsageHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GradeUsageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}

	userIDStr := r.Header.Get("User-Id")
	if userIDStr == "" {
		utils.WriteJSONStatus(w, map[string]string{"error": "userId is required"}, http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "invalid userId"}, http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserById(userID)
	if err != nil || user == nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "invalid userId"}, http.StatusBadRequest)
		return
	}

	grade, err := h.exerciseService.GradeUsage(user, req.Sentece, req.GrammarPatternOrWord)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to generate translation exercise"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, grade)
}

// GenerateTranslationExerciseHandler godoc
//
//	@Summary		Get Translation Exercise
//	@Description	Get a new Translation Exercise based on the input lesson
//	@Tags			exercise
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.GeneraterTranslationExerciseRequest	true	"Exercise Request object to be generated"
//	@Success		200		{object}		agents.GeneratedTranslationExercise
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/resources/exercise/translation/generate [post]
func (h *exerciseHandlerImpl) GenerateTranslationExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GeneraterTranslationExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}

	userIDStr := r.Header.Get("User-Id")
	if userIDStr == "" {
		utils.WriteJSONStatus(w, map[string]string{"error": "userId is required"}, http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "invalid userId"}, http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserById(userID)
	if err != nil || user == nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "invalid userId"}, http.StatusBadRequest)
		return
	}

	exercise, err := h.exerciseService.GenerateTranslationExercise(user, req.LessonId)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to generate translation exercise"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, exercise)
}
