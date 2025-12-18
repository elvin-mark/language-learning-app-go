package services

import (
	models "language-learning-app/models/user"
	"language-learning-app/storage"
)

type UserService interface {
	GetUserById(userId int) (user *storage.User, err error)
	GetUserByUsername(username string) (user *storage.User, err error)
	UpdateUserSettings(userId int, preferredLanguage, targetLanguage string) (user *storage.User, err error)
	GetUserStatus(userId int, targetLanguage string) (report models.UserStatusReport, err error)
}

func NewUserService(userRepository storage.UserRepository, userGrammarRepository storage.UserGrammarRepository, userWordRepository storage.UserWordRepository) UserService {
	return &userServiceImpl{
		userRepository:        userRepository,
		userGrammarRepository: userGrammarRepository,
		userWordRepository:    userWordRepository,
	}
}
