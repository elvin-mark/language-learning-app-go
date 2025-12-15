package handlers

import (
	"net/http"

	"language-learning-app/services"
)

type UserGrammarHandler interface {
	GetGrammarPatternsHandler(w http.ResponseWriter, r *http.Request)
	GetGrammarPatternsByPatternHandler(w http.ResponseWriter, r *http.Request)
}

func NewUserGrammarHandler(userGrammarService services.UserGrammarService, userService services.UserService) UserGrammarHandler {
	return &userGrammarHandlerImpl{userGrammarService: userGrammarService, userService: userService}
}
