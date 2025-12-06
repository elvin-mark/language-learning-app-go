package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// -------------------- LESSON REPO --------------------

type lessonRepositoryImpl struct {
	DB *sql.DB
}

func (r *lessonRepositoryImpl) Create(l *Lesson) error {
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

func (r *lessonRepositoryImpl) GetByID(id int) (*Lesson, error) {
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

func (r *lessonRepositoryImpl) GetPaginatedByLanguageAndUser(userID int, language string, limit, offset int) ([]Lesson, error) {
	query := `
        SELECT lesson_id, user_id, language, grammar_focus, content, new_vocabulary
        FROM lessons
        WHERE user_id = ? AND language = ?
        ORDER BY lesson_id ASC
        LIMIT ? OFFSET ?
    `

	rows, err := r.DB.Query(query, userID, language, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query lessons by user and language: %w", err)
	}
	defer rows.Close()

	lessons := []Lesson{}

	for rows.Next() {
		var l Lesson
		var newVocabJSON string

		if err := rows.Scan(
			&l.LessonID,
			&l.UserID,
			&l.Language,
			&l.GrammarFocus,
			&l.Content,
			&newVocabJSON,
		); err != nil {
			return nil, fmt.Errorf("failed to scan lesson row: %w", err)
		}

		if err := json.Unmarshal([]byte(newVocabJSON), &l.NewVocabulary); err != nil {
			return nil, fmt.Errorf("failed to unmarshal new_vocabulary JSON: %w", err)
		}

		lessons = append(lessons, l)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return lessons, nil
}

func (r *lessonRepositoryImpl) SearchLessonsByGrammarFocus(userID int, language, searchTerm string, limit, offset int) ([]Lesson, error) {
	likePattern := "%" + searchTerm + "%"

	query := `
        SELECT lesson_id, user_id, language, grammar_focus, content, new_vocabulary
        FROM lessons
        WHERE user_id = ? AND language = ? AND grammar_focus LIKE ?
        ORDER BY lesson_id ASC
        LIMIT ? OFFSET ?
    `

	rows, err := r.DB.Query(query, userID, language, likePattern, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to search lessons by grammar focus: %w", err)
	}
	defer rows.Close()

	lessons := []Lesson{}

	for rows.Next() {
		var l Lesson
		var newVocabJSON string

		if err := rows.Scan(
			&l.LessonID,
			&l.UserID,
			&l.Language,
			&l.GrammarFocus,
			&l.Content,
			&newVocabJSON,
		); err != nil {
			return nil, fmt.Errorf("failed to scan lesson row: %w", err)
		}

		if err := json.Unmarshal([]byte(newVocabJSON), &l.NewVocabulary); err != nil {
			return nil, fmt.Errorf("failed to unmarshal new_vocabulary JSON: %w", err)
		}

		lessons = append(lessons, l)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return lessons, nil
}
