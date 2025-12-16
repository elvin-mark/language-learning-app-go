package handlers

import (
	"encoding/json"
	dto "language-learning-app/dto/chatbot"
	"language-learning-app/services"
	"language-learning-app/utils"
	"net/http"
	"strconv"
)

type chatbotHandlerImpl struct {
	chatbotService services.ChatbotService
	userService    services.UserService
}

// GetResponseHandler godoc
//
//	@Summary		Get Chatbot response
//	@Description	Get chatbot response
//	@Tags			chatbot
//	@Accept			json
//	@Produce		json
//	@Param			chatbot_request	body		dto.GetChatbotResponseRequest	true	"Get Chatbot Response Request"
//	@Success		200				{object}	agents.ChatbotResponse
//	@Failure		400				{object}	map[string]string
//	@Failure		500				{object}	map[string]string
//	@Router			/chatbot/response [post]
func (h *chatbotHandlerImpl) GetResponseHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GetChatbotResponseRequest
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

	grade, err := h.chatbotService.GetResponse(user, req.Question)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed to generate chatbot response"}, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, grade)
}
