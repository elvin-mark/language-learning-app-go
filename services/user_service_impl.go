package services

import (
	storage "language-learning-app/storage"
	"language-learning-app/utils"
)

type userServiceImpl struct {
	userRepository storage.UserRepository
}

func NewUserService(userRepository storage.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
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
