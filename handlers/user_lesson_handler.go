package handlers

import (
	"language-learning-app/services"
	"net/http"
)

type UserLessonHandler interface {
	GetLessonsHandler(w http.ResponseWriter, r *http.Request)
	GenerateLessonHandler(w http.ResponseWriter, r *http.Request)
}

func NewUserLessonHandler(userLessonService services.UserLessonService, userService services.UserService) UserLessonHandler {
	return &userLessonHandlerImpl{userLessonService: userLessonService, userService: userService}
}
