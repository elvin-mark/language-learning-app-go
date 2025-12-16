package handlers

import (
	"language-learning-app/services"
	"net/http"
)

type ExerciseHandler interface {
	GradeUsageHandler(w http.ResponseWriter, r *http.Request)
	GenerateTranslationExerciseHandler(w http.ResponseWriter, r *http.Request)
	GradeTranslationExerciseHandler(w http.ResponseWriter, r *http.Request)
	GenerateReadingComprehensionExerciseHandler(w http.ResponseWriter, r *http.Request)
}

func NewExerciseHandler(exerciseService services.ExerciseService, userService services.UserService) ExerciseHandler {
	return &exerciseHandlerImpl{exerciseService: exerciseService, userService: userService}
}
