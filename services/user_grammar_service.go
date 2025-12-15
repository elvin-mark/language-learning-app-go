package services

import "language-learning-app/storage"

type UserGrammarService interface {
	GetGrammarPatterns(userId int, lang string, page, pageSize int) (grammars []storage.UserGrammar, err error)
	GetGrammarPatternsByPattern(userId int, lang string, pattern string, page, pageSize int) (grammars []storage.UserGrammar, err error)
}

func NewUserGrammarService(userGrammarRepository storage.UserGrammarRepository) UserGrammarService {
	return &userGrammarServiceImpl{
		userGrammarRepository: userGrammarRepository,
	}
}
