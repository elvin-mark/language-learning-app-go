package services

import "language-learning-app/storage"

type UserWordService interface {
	GetWords(userId int, lang string, page, pageSize int) (words []storage.UserWord, err error)
}

func NewUserWordService(userWordRepository storage.UserWordRepository) UserWordService {
	return &userWordServiceImpl{
		userWordRepository: userWordRepository,
	}
}
