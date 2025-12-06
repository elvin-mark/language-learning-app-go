package handlers

import (
	"language-learning-app/services"
	"language-learning-app/utils"
	"net/http"
	"strconv"
)

// ============== STRUCTS ==============

type LessonHandler struct {
	service services.LessonService
}

func NewLessonHandler(service services.LessonService) *LessonHandler {
	return &LessonHandler{service: service}
}

// ============== METHODS ==============

// GetLessonsHandler godoc
// @Summary Get Lessons
// @Description Get paginated lessons for a user and language
// @Tags lessons
// @Accept json
// @Produce json
// @Param userId query int true "User ID"
// @Param language query string true "Language"
// @Param page query int false "Page number (default 1)"
// @Param pageSize query int false "Page size (default 20)"
// @Success 200 {array} storage.Lesson
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /resources/lessons [get]
func (h *LessonHandler) GetLessonsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	userIDStr := query.Get("userId")
	if userIDStr == "" {
		utils.WriteJSONStatus(w, map[string]string{"error": "userId is required"}, http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "invalid userId"}, http.StatusBadRequest)
		return
	}

	lang := query.Get("language")
	if lang == "" {
		utils.WriteJSONStatus(w, map[string]string{"error": "language is required"}, http.StatusBadRequest)
		return
	}

	page, _ := strconv.Atoi(query.Get("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(query.Get("pageSize"))
	if pageSize < 1 {
		pageSize = 20
	}

	lessons, err := h.service.GetLessons(userID, lang, page-1, pageSize)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "failed to get lessons"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, lessons)
}

// GetLessonsByGrammarHandler godoc
// @Summary Get Lessons by Grammar
// @Description Get paginated lessons for a user, language, and grammar pattern
// @Tags lessons
// @Accept json
// @Produce json
// @Param userId query int true "User ID"
// @Param language query string true "Language"
// @Param grammarPattern query string true "Grammar pattern to filter"
// @Param page query int false "Page number (default 1)"
// @Param pageSize query int false "Page size (default 20)"
// @Success 200 {array} storage.Lesson
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /resources/lessons/search [get]
func (h *LessonHandler) GetLessonsByGrammarHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	userIDStr := query.Get("userId")
	if userIDStr == "" {
		utils.WriteJSONStatus(w, map[string]string{"error": "userId is required"}, http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "invalid userId"}, http.StatusBadRequest)
		return
	}

	lang := query.Get("language")
	if lang == "" {
		utils.WriteJSONStatus(w, map[string]string{"error": "language is required"}, http.StatusBadRequest)
		return
	}

	grammarPattern := query.Get("grammarPattern")
	if grammarPattern == "" {
		utils.WriteJSONStatus(w, map[string]string{"error": "grammarPattern is required"}, http.StatusBadRequest)
		return
	}

	page, _ := strconv.Atoi(query.Get("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(query.Get("pageSize"))
	if pageSize < 1 {
		pageSize = 20
	}

	lessons, err := h.service.GetLessonsByGrammar(userID, lang, grammarPattern, page-1, pageSize)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "failed to get lessons"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, lessons)
}
