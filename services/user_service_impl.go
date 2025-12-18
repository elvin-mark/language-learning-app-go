package services

import (
	models "language-learning-app/models/user"
	storage "language-learning-app/storage"
	"language-learning-app/utils"
)

type userServiceImpl struct {
	userRepository        storage.UserRepository
	userGrammarRepository storage.UserGrammarRepository
	userWordRepository    storage.UserWordRepository
}

func (us *userServiceImpl) GetUserById(userId int) (user *storage.User, err error) {
	user, err = us.userRepository.GetByID(userId)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	return
}

func (us *userServiceImpl) GetUserByUsername(username string) (user *storage.User, err error) {
	user, err = us.userRepository.GetByUsername(username)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	return
}

func (us *userServiceImpl) UpdateUserSettings(userId int, preferredLanguage, targetLanguage string) (user *storage.User, err error) {
	user, err = us.userRepository.GetByID(userId)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	user.PreferredLanguage = preferredLanguage
	user.TargetLanguage = targetLanguage
	err = us.userRepository.Update(user)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	return
}

func (us *userServiceImpl) GetUserStatus(userId int, targetLanguage string) (report models.UserStatusReport, err error) {
	report.MasteredWords, err = us.userWordRepository.GetUserLearnedWords(userId, targetLanguage, 70)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	report.TotalWords, err = us.userWordRepository.GetUserTotalWords(userId, targetLanguage)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	report.MasteredGrammarPatterns, err = us.userGrammarRepository.GetUserLearnedGrammarPatterns(userId, targetLanguage, 70)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	report.TotalGrammarPatterns, err = us.userGrammarRepository.GetUserTotalGrammarPatterns(userId, targetLanguage)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	return
}
