package agents

import (
	"language-learning-app/core/llm"
)

type ChatbotAgent interface {
	GetResponse(targetLanguage, question string) (response ChatbotResponse, err error)
}

func NewChatbotAgent(llm llm.Llm) ChatbotAgent {
	return &chatbotAgentImpl{
		llm: llm,
	}
}
