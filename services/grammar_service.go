package services

import "language-learning-app/storage"

type GrammarService interface {
	GetGrammarPatterns(userId int, lang string, page, pageSize int) (grammars []storage.GrammarMastery, err error)
	GetGrammarPatternsByPattern(userId int, lang string, pattern string, page, pageSize int) (grammars []storage.GrammarMastery, err error)
}

func NewGrammarService(grammarMasterRepository storage.GrammarMasteryRepository) GrammarService {
	return &grammarServiceImpl{
		grammarMasterRepository: grammarMasterRepository,
	}
}
