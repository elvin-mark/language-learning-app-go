package agents

import (
	"encoding/json"
	"language-learning-app/core/llm"
	"language-learning-app/utils"
	"strings"
)

type lessonAgentImpl struct {
	llm llm.Llm
}

func (la *lessonAgentImpl) GenerateLesson(targetLanguage string, grammarPattern string, words []string) (generatedLesson GeneratedLesson, err error) {
	prompt := generateLessonGenerationPrompt(targetLanguage, grammarPattern, words)
	resp, err := la.llm.GetResponse(prompt)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	cleaned := strings.TrimPrefix(resp.Choices[0].Message.Content, "```json")
	cleaned = strings.TrimSuffix(cleaned, "```")
	cleaned = strings.TrimSpace(cleaned)

	if err = json.Unmarshal([]byte(cleaned), &generatedLesson); err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	return
}
