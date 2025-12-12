package services

import (
	"language-learning-app/storage"
	"language-learning-app/utils"
)

type grammarServiceImpl struct {
	grammarMasterRepository storage.GrammarMasteryRepository
}

func (gs *grammarServiceImpl) GetGrammarPatterns(userId int, lang string, page, pageSize int) (grammars []storage.GrammarMastery, err error) {
	grammars, err = gs.grammarMasterRepository.GetPaginatedForUser(userId, lang, page*pageSize, pageSize)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	return
}

func (gs *grammarServiceImpl) GetGrammarPatternsByPattern(userId int, lang string, pattern string, page, pageSize int) (grammars []storage.GrammarMastery, err error) {
	grammars, err = gs.grammarMasterRepository.SearchByPattern(userId, lang, pattern, page*pageSize, page)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	return
}
