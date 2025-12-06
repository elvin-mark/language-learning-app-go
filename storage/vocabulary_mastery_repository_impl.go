package storage

import "database/sql"

// -------------------- VOCABULARY MASTERY REPO --------------------

type vocabularyMasteryRepositoryImpl struct {
	DB *sql.DB
}

func (r *vocabularyMasteryRepositoryImpl) Upsert(v *VocabularyMastery) error {
	_, err := r.DB.Exec(`
        INSERT INTO vocabulary_mastery (user_id, language, word, mastery_score, times_correct, times_incorrect)
        VALUES (?, ?, ?, ?, ?, ?)
        ON CONFLICT(user_id, word)
        DO UPDATE SET 
            mastery_score = excluded.mastery_score,
            times_correct = excluded.times_correct,
            times_incorrect = excluded.times_incorrect,
            last_reviewed = CURRENT_TIMESTAMP;
    `, v.UserID, v.Language, v.Word, v.MasteryScore, v.TimesCorrect, v.TimesIncorrect)

	return err
}

func (r *vocabularyMasteryRepositoryImpl) GetLowestBelowScore(userID int, maxScore float64) ([]VocabularyMastery, error) {
	rows, err := r.DB.Query(`
        SELECT mastery_id, user_id, language, word, mastery_score, last_reviewed,
               times_correct, times_incorrect
        FROM vocabulary_mastery
        WHERE user_id = ? AND mastery_score < ?
        ORDER BY mastery_score ASC
        LIMIT 20;
    `, userID, maxScore)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []VocabularyMastery
	for rows.Next() {
		var v VocabularyMastery
		if err := rows.Scan(
			&v.MasteryID, &v.UserID, &v.Language, &v.Word, &v.MasteryScore,
			&v.LastReviewed, &v.TimesCorrect, &v.TimesIncorrect,
		); err != nil {
			return nil, err
		}
		results = append(results, v)
	}

	return results, nil
}

func (r *vocabularyMasteryRepositoryImpl) GetPaginated(userID int, lang string, offset, limit int) ([]VocabularyMastery, error) {
	rows, err := r.DB.Query(`
        SELECT mastery_id, user_id, language, word, mastery_score, last_reviewed, times_correct, times_incorrect
        FROM vocabulary_mastery
        WHERE user_id = ?
		AND language = ?
        ORDER BY mastery_id
        LIMIT ? OFFSET ?`,
		userID, lang, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []VocabularyMastery

	for rows.Next() {
		var v VocabularyMastery
		if err := rows.Scan(
			&v.MasteryID,
			&v.UserID,
			&v.Language,
			&v.Word,
			&v.MasteryScore,
			&v.LastReviewed,
			&v.TimesCorrect,
			&v.TimesIncorrect,
		); err != nil {
			return nil, err
		}
		list = append(list, v)
	}

	return list, nil
}
