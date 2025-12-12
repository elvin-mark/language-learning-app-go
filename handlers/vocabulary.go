package handlers

import (
	"net/http"
	"strconv"

	"language-learning-app/services"
	"language-learning-app/utils"
)

// ============== STRUCTS ==============

type VocabularyHandler struct {
	service services.VocabularyService
}

func NewVocabularyHandler(service services.VocabularyService) *VocabularyHandler {
	return &VocabularyHandler{service: service}
}

// ============== METHODS ==============

// GetVocabularyHandler godoc
//
//	@Summary		Get Vocabulary
//	@Description	Get paginated vocabulary for a user and language
//	@Tags			vocabulary
//	@Accept			json
//	@Produce		json
//	@Param			language	query		string	true	"Language"
//	@Param			page		query		int		false	"Page number (default 1)"
//	@Param			pageSize	query		int		false	"Page size (default 20)"
//	@Success		200			{array}		storage.VocabularyMastery
//	@Failure		400			{object}	map[string]string
//	@Failure		500			{object}	map[string]string
//	@Router			/resources/vocabulary [get]
func (h *VocabularyHandler) GetVocabularyHandler(w http.ResponseWriter, r *http.Request) {
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

	words, err := h.service.GetVocabulary(userID, lang, page-1, pageSize)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "failed to get vocabulary"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, words)
}
