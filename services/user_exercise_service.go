package services

import (
	"language-learning-app/agents"
	"language-learning-app/storage"
)

type ExerciseService interface {
	GradeUsage(user *storage.User, sentence, grammarPatternOrWord string) (grade agents.UsageGrade, err error)
	GenerateTranslationExercise(user *storage.User, lessonId int) (exercise agents.GeneratedTranslationExercise, err error)
	GradeTranslationExercise(user *storage.User, lessonId int, sentence string) (grade []agents.UsageGrade, err error)
}

func NewExerciseService(exerciseAgent agents.ExerciseAgent, userLessonRepository storage.UserLessonRepository, userGrammarRepository storage.UserGrammarRepository, userWordRepository storage.UserWordRepository) ExerciseService {
	return &exerciseServiceImpl{
		exerciseAgent:         exerciseAgent,
		userLessonRepository:  userLessonRepository,
		userGrammarRepository: userGrammarRepository,
		userWordRepository:    userWordRepository,
	}
}
