package handlers

import (
	"net/http"
	"strconv"

	"language-learning-app/services"
	"language-learning-app/utils"
)

// ============== STRUCTS ==============

type GrammarHandler struct {
	service services.GrammarService
}

func NewGrammarHandler(service services.GrammarService) *GrammarHandler {
	return &GrammarHandler{service: service}
}

// ============== METHODS ==============

// GetGrammarPatternsHandler godoc
//
//	@Summary		Get Grammar Patterns
//	@Description	Get paginated grammar patterns for a user and language
//	@Tags			grammar
//	@Accept			json
//	@Produce		json
//	@Param			userId		query		int		true	"User ID"
//	@Param			language	query		string	true	"Language"
//	@Param			page		query		int		false	"Page number (default 1)"
//	@Param			pageSize	query		int		false	"Page size (default 20)"
//	@Success		200			{array}		storage.GrammarMastery
//	@Failure		400			{object}	map[string]string
//	@Failure		500			{object}	map[string]string
//	@Router			/resources/grammar [get]
func (h *GrammarHandler) GetGrammarPatternsHandler(w http.ResponseWriter, r *http.Request) {
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

	grammars, err := h.service.GetGrammarPatterns(userID, lang, page-1, pageSize)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "failed to get grammar patterns"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, grammars)
}

// GetGrammarPatternsByPatternHandler godoc
//
//	@Summary		Search Grammar Patterns
//	@Description	Get paginated grammar patterns for a user and language filtered by pattern (LIKE search)
//	@Tags			grammar
//	@Accept			json
//	@Produce		json
//	@Param			userId		query		int		true	"User ID"
//	@Param			language	query		string	true	"Language"
//	@Param			pattern		query		string	true	"Grammar pattern to filter"
//	@Param			page		query		int		false	"Page number (default 1)"
//	@Param			pageSize	query		int		false	"Page size (default 20)"
//	@Success		200			{array}		storage.GrammarMastery
//	@Failure		400			{object}	map[string]string
//	@Failure		500			{object}	map[string]string
//	@Router			/resources/grammar/search [get]
func (h *GrammarHandler) GetGrammarPatternsByPatternHandler(w http.ResponseWriter, r *http.Request) {
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

	pattern := query.Get("pattern")
	if pattern == "" {
		utils.WriteJSONStatus(w, map[string]string{"error": "pattern is required"}, http.StatusBadRequest)
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

	grammars, err := h.service.GetGrammarPatternsByPattern(userID, lang, pattern, page-1, pageSize)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "failed to search grammar patterns"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, grammars)
}
