package services

import (
	"language-learning-app/core/agents"
	"language-learning-app/storage"
)

type ExerciseService interface {
	GradeUsage(user *storage.User, sentence, grammarPatternOrWord string) (grade agents.UsageGrade, err error)
	GenerateTranslationExercise(user *storage.User, lessonId int) (exercise agents.GeneratedTranslationExercise, err error)
	GradeTranslationExercise(user *storage.User, lessonId int, originalSentence, translatedSentence string) (grade []agents.UsageGrade, err error)
	GenerateReadingComprehensionExercise(user *storage.User, lessonId int) (exercise agents.GeneratedReadingComprehensionExercise, err error)
	GradeReadingComprehensionResponse(user *storage.User, lessonId int, shortText, question, answer string) (grade []agents.UsageGrade, err error)
	GenerateDialogueInitExercise(user *storage.User, lessonId int) (exercise agents.GeneratedDialogueInitExercise, err error)
	GenerateDialogueContinuationExercise(user *storage.User, lessonId int, history string) (exercise agents.GeneratedDialogueContinuationExercise, err error)
}

func NewExerciseService(exerciseAgent agents.ExerciseAgent, userLessonRepository storage.UserLessonRepository, userGrammarRepository storage.UserGrammarRepository, userWordRepository storage.UserWordRepository) ExerciseService {
	return &exerciseServiceImpl{
		exerciseAgent:         exerciseAgent,
		userLessonRepository:  userLessonRepository,
		userGrammarRepository: userGrammarRepository,
		userWordRepository:    userWordRepository,
	}
}
