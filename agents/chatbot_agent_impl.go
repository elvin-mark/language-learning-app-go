package agents

import (
	"encoding/json"
	"language-learning-app/core/llm"
	"language-learning-app/utils"
	"strings"
)

type chatbotAgentImpl struct {
	llm llm.Llm
}

func (ca *chatbotAgentImpl) GetResponse(targetLanguage, question string) (response ChatbotResponse, err error) {
	prompt := generateChatbotPrompt(targetLanguage, question)

	resp, err := ca.llm.GetResponse(prompt)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	cleaned := strings.TrimPrefix(resp.Choices[0].Message.Content, "```json")
	cleaned = strings.TrimSuffix(cleaned, "```")
	cleaned = strings.TrimSpace(cleaned)

	utils.Logger.Debug("Response from LLM: " + cleaned)
	if err = json.Unmarshal([]byte(cleaned), &response); err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	return
}
