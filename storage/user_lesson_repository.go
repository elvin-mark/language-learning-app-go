package storage

import (
	"database/sql"
)

// -------------------- LESSON REPO --------------------

type UserLessonRepository interface {
	Create(l *UserLesson) error
	GetByID(id int) (*UserLesson, error)
	GetPaginatedByLanguageAndUser(userID int, language string, limit, offset int) ([]UserLesson, error)
}

func NewUserLessonRepository(db *sql.DB) UserLessonRepository {
	return &userLessonRepositoryImpl{
		DB: db,
	}
}
