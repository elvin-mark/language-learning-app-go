package agents

import (
	"encoding/json"
	"language-learning-app/core/llm"
	"language-learning-app/utils"
	"strings"
)

type exerciseAgent struct {
	llm llm.Llm
}

func (pa *exerciseAgent) GenerateTranslationExercise(preferredLanguage, targetLanguage string, grammarPattern string, words []string) (generatedExercise GeneratedTranslationExercise, err error) {
	prompt := generateTranslationExercisePrompt(preferredLanguage, targetLanguage, grammarPattern, words)
	resp, err := pa.llm.GetResponse(prompt)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	cleaned := strings.TrimPrefix(resp.Choices[0].Message.Content, "```json")
	cleaned = strings.TrimSuffix(cleaned, "```")
	cleaned = strings.TrimSpace(cleaned)

	if err = json.Unmarshal([]byte(cleaned), &generatedExercise); err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	return
}
