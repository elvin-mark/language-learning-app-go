package services

import (
	"language-learning-app/core/agents"
	"language-learning-app/storage"
)

type ChatbotService interface {
	GetResponse(user *storage.User, question string) (response agents.ChatbotResponse, err error)
}

func NewChatbotService(chatbotAgent agents.ChatbotAgent) ChatbotService {
	return &chatbotServiceImpl{
		chatbotAgent: chatbotAgent,
	}
}
