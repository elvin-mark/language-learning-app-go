package services

import (
	"language-learning-app/storage"
	"language-learning-app/utils"
)

type userGrammarServiceImpl struct {
	userGrammarRepository storage.UserGrammarRepository
}

func (gs *userGrammarServiceImpl) GetGrammarPatterns(userId int, lang string, page, pageSize int) (grammars []storage.UserGrammar, err error) {
	grammars, err = gs.userGrammarRepository.GetPaginatedForUser(userId, lang, page*pageSize, pageSize)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	return
}

func (gs *userGrammarServiceImpl) GetGrammarPatternsByPattern(userId int, lang string, pattern string, page, pageSize int) (grammars []storage.UserGrammar, err error) {
	grammars, err = gs.userGrammarRepository.SearchByPattern(userId, lang, pattern, page*pageSize, pageSize)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	return
}
