package handlers

import (
	"language-learning-app/services"
	"net/http"
)

type AuthHandler interface {
	GetAuthTokenHandler(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(userService services.UserService) AuthHandler {
	return &authHandlerImpl{
		userService: userService,
	}
}
