package storage

import (
	"database/sql"
)

type UserRepository interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}
