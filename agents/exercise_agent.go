package agents

import (
	"language-learning-app/core/llm"
)

type ExerciseAgent interface {
	GradeUsage(targetLanguage, sentence string, wordOrGrammarPattern string) (grade UsageGrade, err error)
	GenerateTranslationExercise(preferredLanguage, targetLanguage string, grammarPattern string, words []string) (generatedExercise GeneratedTranslationExercise, err error)
	GenerateReadingComprehensionExercise(targetLanguage string, grammarPattern string, words []string) (generatedExercise GeneratedReadingComprehensionExercise, err error)
}

func NewExerciseAgent(llm llm.Llm) ExerciseAgent {
	return &exerciseAgent{
		llm: llm,
	}
}
