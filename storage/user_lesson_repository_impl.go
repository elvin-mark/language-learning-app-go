package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"language-learning-app/utils"
)

// -------------------- LESSON REPO --------------------

type userLessonRepositoryImpl struct {
	DB *sql.DB
}

func (r *userLessonRepositoryImpl) Create(l *UserLesson) error {
	vocabJSON, _ := utils.ToJSON(l.WordsId)

	res, err := r.DB.Exec(`
        INSERT INTO user_lessons (user_id, language, grammar_id, words_id, content, sample_sentences, words_meaning)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,
		l.UserId, l.Language, l.GrammarId, vocabJSON, l.Content, l.SampleSentences, l.WordsMeaning)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	l.Id = int(id)
	return nil
}

func (r *userLessonRepositoryImpl) GetByID(id int) (*UserLesson, error) {
	query := `
        SELECT id, user_id, language, grammar_id, words_id, content, sample_sentences, words_meaning
        FROM user_lessons
        WHERE id = ?
    `

	row := r.DB.QueryRow(query, id)

	var (
		lesson          UserLesson
		newVocabJSONRaw string
	)

	err := row.Scan(
		&lesson.Id,
		&lesson.UserId,
		&lesson.Language,
		&lesson.GrammarId,
		&newVocabJSONRaw,
		&lesson.Content,
		&lesson.SampleSentences,
		&lesson.WordsMeaning,
	)

	if err == sql.ErrNoRows {
		return nil, nil // not found
	}

	if err != nil {
		return nil, err
	}

	// Parse JSON array into []string
	if err := json.Unmarshal([]byte(newVocabJSONRaw), &lesson.WordsId); err != nil {
		return nil, fmt.Errorf("failed to parse words_id: %w", err)
	}

	return &lesson, nil
}

func (r *userLessonRepositoryImpl) GetPaginatedByLanguageAndUser(userID int, language string, limit, offset int) ([]UserLesson, error) {
	query := `
        SELECT id, user_id, language, grammar_id, words_id, content, sample_sentences, words_meaning
        FROM user_lessons
        WHERE user_id = ? AND language = ?
        ORDER BY id ASC
        LIMIT ? OFFSET ?
    `

	rows, err := r.DB.Query(query, userID, language, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query lessons by user and language: %w", err)
	}
	defer rows.Close()

	lessons := []UserLesson{}

	for rows.Next() {
		var l UserLesson
		var newVocabJSON string

		if err := rows.Scan(
			&l.Id,
			&l.UserId,
			&l.Language,
			&l.GrammarId,
			&newVocabJSON,
			&l.Content,
			&l.SampleSentences,
			&l.WordsMeaning,
		); err != nil {
			return nil, fmt.Errorf("failed to scan lesson row: %w", err)
		}

		if err := json.Unmarshal([]byte(newVocabJSON), &l.WordsId); err != nil {
			return nil, fmt.Errorf("failed to unmarshal words_id JSON: %w", err)
		}

		lessons = append(lessons, l)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return lessons, nil
}
