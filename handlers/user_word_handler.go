package handlers

import (
	"net/http"

	"language-learning-app/services"
)

type UserWordHandler interface {
	GetWordsHandler(w http.ResponseWriter, r *http.Request)
}

func NewUserWordHandler(userWordService services.UserWordService, userService services.UserService) UserWordHandler {
	return &userWordHandlerImpl{userWordService: userWordService, userService: userService}
}
