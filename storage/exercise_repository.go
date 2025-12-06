package storage

import "database/sql"

// -------------------- EXERCISE REPO --------------------

type ExerciseRepository interface {
	Create(e *Exercise) error
}

func NewExerciseRepository(db *sql.DB) ExerciseRepository {
	return &exerciseRepositoryImpl{
		DB: db,
	}
}
