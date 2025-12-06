package storage

import (
	"database/sql"
)

// -------------------- LESSON REPO --------------------

type LessonRepository interface {
	Create(l *Lesson) error
	GetByID(id int) (*Lesson, error)
	GetPaginatedByLanguageAndUser(userID int, language string, limit, offset int) ([]Lesson, error)
	SearchLessonsByGrammarFocus(userID int, language, searchTerm string, limit, offset int) ([]Lesson, error)
}

func NewLessonRepository(db *sql.DB) LessonRepository {
	return &lessonRepositoryImpl{
		DB: db,
	}
}
