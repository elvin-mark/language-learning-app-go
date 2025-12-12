package services

import (
	"language-learning-app/storage"
	"language-learning-app/utils"
)

type vocabularyServiceImpl struct {
	vocabularyMasteryRepository storage.VocabularyMasteryRepository
}

func (vs *vocabularyServiceImpl) GetVocabulary(userId int, lang string, page, pageSize int) (words []storage.VocabularyMastery, err error) {
	words, err = vs.vocabularyMasteryRepository.GetPaginated(userId, lang, page*pageSize, pageSize)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	return
}
