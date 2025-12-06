package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// -------------------- LESSON REPO --------------------

type LessonRepository struct {
	DB *sql.DB
}

func (r *LessonRepository) Create(l *Lesson) error {
	vocabJSON, _ := toJSON(l.NewVocabulary)

	res, err := r.DB.Exec(`
        INSERT INTO lessons (user_id, language, grammar_focus, content, new_vocabulary)
        VALUES (?, ?, ?, ?, ?)`,
		l.UserID, l.Language, l.GrammarFocus, l.Content, vocabJSON)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	l.LessonID = int(id)
	return nil
}

func (r *LessonRepository) GetByID(id int) (*Lesson, error) {
	query := `
        SELECT lesson_id, user_id, language, grammar_focus, content, new_vocabulary
        FROM lessons
        WHERE lesson_id = ?
    `

	row := r.DB.QueryRow(query, id)

	var (
		lesson          Lesson
		newVocabJSONRaw string
	)

	err := row.Scan(
		&lesson.LessonID,
		&lesson.UserID,
		&lesson.Language,
		&lesson.GrammarFocus,
		&lesson.Content,
		&newVocabJSONRaw,
	)

	if err == sql.ErrNoRows {
		return nil, nil // not found
	}

	if err != nil {
		return nil, err
	}

	// Parse JSON array into []string
	if err := json.Unmarshal([]byte(newVocabJSONRaw), &lesson.NewVocabulary); err != nil {
		return nil, fmt.Errorf("failed to parse new_vocabulary: %w", err)
	}

	return &lesson, nil
}
