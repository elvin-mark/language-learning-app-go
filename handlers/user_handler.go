package handlers

import (
	"language-learning-app/services"
	"net/http"
)

type UserHandler interface {
	GetUserProfileHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserSettingsHandler(w http.ResponseWriter, r *http.Request)
	GetUserStatusReport(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandlerImpl{userService: userService}
}
