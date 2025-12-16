package handlers

import (
	"language-learning-app/services"
	"net/http"
)

type ChatbotHandler interface {
	GetResponseHandler(w http.ResponseWriter, r *http.Request)
}

func NewChatbotHandler(chatbotService services.ChatbotService, userService services.UserService) ChatbotHandler {
	return &chatbotHandlerImpl{
		chatbotService: chatbotService,
		userService:    userService,
	}
}
