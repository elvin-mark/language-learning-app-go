package agents

import (
	"encoding/json"
	"language-learning-app/core/llm"
	"language-learning-app/storage"
	"math/rand"
	"strings"
)

type lessonAgentImpl struct {
	llm                         llm.Llm
	lessonRepository            storage.LessonRepository
	vocabularyMasteryRepository storage.VocabularyMasteryRepository
	grammarMasteryRepository    storage.GrammarMasteryRepository
}

func (la *lessonAgentImpl) GenerateLesson(userId int, lang string, masteryScoreThreshold float64) (lesson storage.Lesson, err error) {

	grammars, err := la.grammarMasteryRepository.GetLowestBelowScore(userId, masteryScoreThreshold)
	if err != nil {
		return
	}
	vocabs, err := la.vocabularyMasteryRepository.GetLowestBelowScore(userId, masteryScoreThreshold)
	if err != nil {
		return
	}

	randomGrammar := grammars[rand.Intn(len(grammars))]
	prompt := generateLessonGenerationPrompt(lang, randomGrammar, vocabs)
	resp, err := la.llm.GetResponse(prompt)
	if err != nil {
		return
	}

	cleaned := strings.TrimPrefix(resp.Choices[0].Message.Content, "```json")
	cleaned = strings.TrimSuffix(cleaned, "```")
	cleaned = strings.TrimSpace(cleaned)

	var generatedLesson GeneratedLesson
	if err = json.Unmarshal([]byte(cleaned), &generatedLesson); err != nil {
		return
	}

	lesson = storage.Lesson{
		UserID:        userId,
		Language:      lang,
		Content:       generatedLesson.ExplanationText + "\nExample Sentences: \n" + strings.Join(generatedLesson.ExampleSentences, "\n"),
		GrammarFocus:  generatedLesson.GrammarPattern,
		NewVocabulary: generatedLesson.NewVocabulary,
	}

	la.lessonRepository.Create(&lesson)
	return
}
