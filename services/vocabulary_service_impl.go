package services

import "language-learning-app/storage"

type vocabularyServiceImpl struct {
	vocabularyMasteryRepository storage.VocabularyMasteryRepository
}

func (vs *vocabularyServiceImpl) GetVocabulary(userId int, lang string, page, pageSize int) (words []storage.VocabularyMastery, err error) {
	words, err = vs.vocabularyMasteryRepository.GetPaginated(userId, lang, page*pageSize, pageSize)
	return
}
