package services

import "language-learning-app/storage"

type VocabularyService interface {
	GetVocabulary(userId int, lang string, page, pageSize int) (words []storage.VocabularyMastery, err error)
}

func NewVocabularyService(vocabularyMasteryRepository storage.VocabularyMasteryRepository) VocabularyService {
	return &vocabularyServiceImpl{
		vocabularyMasteryRepository: vocabularyMasteryRepository,
	}
}
