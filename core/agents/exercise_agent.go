package agents

import (
	"language-learning-app/core/llm"
)

type ExerciseAgent interface {
	GradeUsage(targetLanguage, sentence string, wordOrGrammarPattern string) (grade UsageGrade, err error)
	GenerateTranslationExercise(preferredLanguage, targetLanguage string, grammarPattern string, words []string) (generatedExercise GeneratedTranslationExercise, err error)
	GradeTranslation(targetLanguage, originalSentence string, translatedSentence string) (grade UsageGrade, err error)
	GenerateReadingComprehensionExercise(targetLanguage string, grammarPattern string, words []string) (generatedExercise GeneratedReadingComprehensionExercise, err error)
	GradeReadingComprehensionResponse(targetLanguage, shortText, question, answer string) (grade UsageGrade, err error)
	GenerateDialogueInitExercise(targetLanguage string, grammarPattern string, words []string) (generatedExercise GeneratedDialogueInitExercise, err error)
	GenerateDialogueContinuationExercise(targetLanguage string, grammarPattern string, words []string, history string) (generatedExercise GeneratedDialogueContinuationExercise, err error)
}

func NewExerciseAgent(llm llm.Llm) ExerciseAgent {
	return &exerciseAgentImpl{
		llm: llm,
	}
}
