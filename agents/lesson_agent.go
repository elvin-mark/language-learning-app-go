package agents

import (
	"language-learning-app/core/llm"
	"language-learning-app/storage"
)

type LessonAgent interface {
	GenerateLesson(userId int, lang string, masteryScoreThreshold float64) (lesson storage.Lesson, err error)
}

func NewLessonAgent(llm llm.Llm, lessonRepository storage.LessonRepository, vocabularyMasteryRepository storage.VocabularyMasteryRepository, grammarMasteryRepository storage.GrammarMasteryRepository) LessonAgent {
	return &lessonAgentImpl{
		llm:                         llm,
		lessonRepository:            lessonRepository,
		vocabularyMasteryRepository: vocabularyMasteryRepository,
		grammarMasteryRepository:    grammarMasteryRepository,
	}
}
