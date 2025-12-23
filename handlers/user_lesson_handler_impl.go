package handlers

import (
	"encoding/json"
	"io"
	dto "language-learning-app/dto/lesson"
	"language-learning-app/services"
	"language-learning-app/utils"
	"net/http"
	"strconv"
)

type userLessonHandlerImpl struct {
	userLessonService services.UserLessonService
	userService       services.UserService
}

// GetLessonsHandler godoc
//
//	@Summary		Get Lessons
//	@Description	Get paginated lessons for a user
//	@Tags			lessons
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"Page number (default 1)"
//	@Param			pageSize	query		int	false	"Page size (default 20)"
//	@Success		200			{array}		models.LessonItem
//	@Failure		400			{object}	map[string]string
//	@Failure		500			{object}	map[string]string
//	@Router			/resources/lessons [get]
func (h *userLessonHandlerImpl) GetLessonsHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("User-Id")
	query := r.URL.Query()

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
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "invalid userId"}, http.StatusBadRequest)
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

	lessons, err := h.userLessonService.GetLessons(userID, user.TargetLanguage, page-1, pageSize)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "failed to get lessons"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, lessons)
}

// GenerateLessonHandler godoc
//
//	@Summary		Generate New Lesson
//	@Description	Get a new Language Lesson based on the user current status
//	@Tags			lessons
//	@Accept			json
//	@Produce		json
//	@Param			lesson_request	body		dto.GenerateLessonRequest	false	"Generate lesson request"
//	@Success		200				{array}		models.LessonItem
//	@Failure		400				{object}	map[string]string
//	@Failure		500				{object}	map[string]string
//	@Router			/resources/lessons/generate [post]
func (h *userLessonHandlerImpl) GenerateLessonHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GenerateLessonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil && err != io.EOF {
		utils.WriteJSONStatus(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}

	userIdStr := r.Header.Get("User-Id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to get userId"}, http.StatusInternalServerError)
		return
	}

	user, err := h.userService.GetUserById(userId)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "invalid userId"}, http.StatusBadRequest)
		return
	}

	lesson, err := h.userLessonService.GenerateLesson(user, req.GrammarId, req.WordsId)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to generate lesson"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, lesson)
}
