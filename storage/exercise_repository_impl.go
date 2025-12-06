package storage

import "database/sql"

// -------------------- EXERCISE REPO --------------------

type exerciseRepositoryImpl struct {
	DB *sql.DB
}

func (r *exerciseRepositoryImpl) Create(e *Exercise) error {
	res, err := r.DB.Exec(`
        INSERT INTO exercises (user_id, lesson_id, type, sub_type, question_data, user_response, grade, feedback)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		e.UserID, e.LessonID, e.Type, e.SubType, e.QuestionData, e.UserResponse, e.Grade, e.Feedback)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	e.ExerciseID = int(id)
	return nil
}
