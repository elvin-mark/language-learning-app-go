package services

import "language-learning-app/storage"

type UserService interface {
	GetUserByUsername(username string) (user *storage.User, err error)
}
