package services

import (
	"language-learning-app/agents"
	"language-learning-app/storage"
	"language-learning-app/utils"
)

type chatbotServiceImpl struct {
	chatbotAgent agents.ChatbotAgent
}

func (es *chatbotServiceImpl) GetResponse(user *storage.User, question string) (response agents.ChatbotResponse, err error) {
	response, err = es.chatbotAgent.GetResponse(user.TargetLanguage, question)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	return
}
