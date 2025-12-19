package handlers

import (
	"encoding/json"
	dto "language-learning-app/dto/exercise"
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
//	@Success		200		{object}	agents.UsageGrade
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
//	@Param			user	body		dto.GenerateTranslationExerciseRequest	true	"Exercise Request object to be generated"
//	@Success		200		{object}	agents.GeneratedTranslationExercise
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/resources/exercise/translation/generate [post]
func (h *exerciseHandlerImpl) GenerateTranslationExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GenerateTranslationExerciseRequest
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

// GradeTranslationExerciseHandler godoc
//
//	@Summary		Grade Translation Exercise
//	@Description	Grade a Translation Exercise based on the input lesson
//	@Tags			exercise
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.GradeTranslationExerciseRequest	true	"Exercise Request object to be graded"
//	@Success		200		{array}		agents.UsageGrade
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/resources/exercise/translation/grade [post]
func (h *exerciseHandlerImpl) GradeTranslationExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GradeTranslationExerciseRequest
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

	grades, err := h.exerciseService.GradeTranslationExercise(user, req.LessonId, req.OriginalSentence, req.TranslatedSentence)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to grade translation exercise"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, grades)
}

// GenerateReadingComprehensionExerciseHandler godoc
//
//	@Summary		Get Reading Comprehension Exercise
//	@Description	Get a new reading comprehension Exercise based on the input lesson
//	@Tags			exercise
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.GenerateReadingComprehensionExerciseRequest	true	"Exercise Request object to be generated"
//	@Success		200		{object}	agents.GeneratedReadingComprehensionExercise
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/resources/exercise/reading-comprehension/generate [post]
func (h *exerciseHandlerImpl) GenerateReadingComprehensionExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GenerateReadingComprehensionExerciseRequest
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

	exercise, err := h.exerciseService.GenerateReadingComprehensionExercise(user, req.LessonId)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to generate reading comprehension exercise"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, exercise)
}

// GenerateDialogueInitExerciseHandler godoc
//
//	@Summary		Init Dialogue Exercise
//	@Description	Init dialogue Exercise based on the input lesson
//	@Tags			exercise
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.GenerateDialogueInitExerciseRequest	true	"Exercise Request object to be generated"
//	@Success		200		{object}	agents.GeneratedDialogueInitExercise
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/resources/exercise/dialogue/init [post]
func (h *exerciseHandlerImpl) GenerateDialogueInitExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GenerateDialogueInitExerciseRequest
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

	exercise, err := h.exerciseService.GenerateDialogueInitExercise(user, req.LessonId)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to initialize dialogue"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, exercise)
}

// GenerateDialogueContinuationExerciseHandler godoc
//
//	@Summary		Continue with Dialogue Exercise
//	@Description	Continue with dialogue Exercise based on the input lesson
//	@Tags			exercise
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.GenerateDialogueContinuationExerciseRequest	true	"Exercise Request object to be generated"
//	@Success		200		{object}	agents.GeneratedDialogueContinuationExercise
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/resources/exercise/dialogue/continue [post]
func (h *exerciseHandlerImpl) GenerateDialogueContinuationExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GenerateDialogueContinuationExerciseRequest
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

	exercise, err := h.exerciseService.GenerateDialogueContinuationExercise(user, req.LessonId, req.History)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to generate dialogue continuation"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, exercise)
}
