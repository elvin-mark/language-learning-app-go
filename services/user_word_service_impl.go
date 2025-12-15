package services

import (
	"language-learning-app/storage"
	"language-learning-app/utils"
)

type userWordServiceImpl struct {
	userWordRepository storage.UserWordRepository
}

func (vs *userWordServiceImpl) GetWords(userId int, lang string, page, pageSize int) (words []storage.UserWord, err error) {
	words, err = vs.userWordRepository.GetPaginated(userId, lang, page*pageSize, pageSize)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	return
}
