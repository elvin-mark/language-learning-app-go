package services

import "language-learning-app/storage"

type UserService interface {
	GetUserById(userId int) (user *storage.User, err error)
	GetUserByUsername(username string) (user *storage.User, err error)
}
