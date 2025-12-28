package agents

import (
	"language-learning-app/core/llm"
)

type LessonAgent interface {
	GenerateLesson(targetLanguage string, grammarPattern string, words []string) (generatedLesson GeneratedLesson, err error)
}

func NewLessonAgent(llm llm.Llm) LessonAgent {
	return &lessonAgentImpl{
		llm: llm,
	}
}
