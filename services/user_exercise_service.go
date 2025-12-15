package services

import (
	"language-learning-app/agents"
	"language-learning-app/storage"
)

type ExerciseService interface {
	GenerateTranslationExercise(user *storage.User, lessonId int) (exercise agents.GeneratedTranslationExercise, err error)
}

func NewExerciseService(exerciseAgent agents.ExerciseAgent) ExerciseService {
	return &exerciseServiceImpl{
		exerciseAgent: exerciseAgent,
	}
}
