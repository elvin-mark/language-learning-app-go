package handlers

import (
	"net/http"
	"strconv"

	"language-learning-app/services"
	"language-learning-app/utils"
)

type userWordHandlerImpl struct {
	userWordService services.UserWordService
	userService     services.UserService
}

// GetWordsHandler godoc
//
//	@Summary		Get Words
//	@Description	Get paginated words for a user
//	@Tags			vocabulary
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			page		query		int	false	"Page number (default 1)"
//	@Param			pageSize	query		int	false	"Page size (default 20)"
//	@Success		200			{array}		storage.UserWord
//	@Failure		400			{object}	map[string]string
//	@Failure		500			{object}	map[string]string
//	@Router			/resources/words [get]
func (h *userWordHandlerImpl) GetWordsHandler(w http.ResponseWriter, r *http.Request) {
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

	words, err := h.userWordService.GetWords(userID, user.TargetLanguage, page-1, pageSize)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "failed to get vocabulary"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, words)
}
