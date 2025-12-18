package storage

import (
	"database/sql"
	"errors"
)

// -------------------- VOCABULARY MASTERY REPO --------------------

type userWordRepositoryImpl struct {
	DB *sql.DB
}

func (r *userWordRepositoryImpl) Upsert(v *UserWord) error {
	_, err := r.DB.Exec(`
        INSERT INTO user_words (user_id, language, type, word, score)
        VALUES (?, ?, ?, ?, ?)
        ON CONFLICT(user_id, word)
        DO UPDATE SET 
            score = excluded.score,
            last_reviewed = CURRENT_TIMESTAMP;
    `, v.UserId, v.Language, v.Type, v.Word, v.Score)

	return err
}

func (r *userWordRepositoryImpl) GetByID(id int) (*UserWord, error) {
	var u UserWord
	row := r.DB.QueryRow(`
        SELECT id, user_id, language, type, word, score, last_reviewed
        FROM user_words WHERE id = ?`, id)

	if err := row.Scan(&u.Id, &u.UserId, &u.Language, &u.Type, &u.Word, &u.Score, &u.LastReviewed); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *userWordRepositoryImpl) GetLowestBelowScore(userID int, maxScore int) ([]UserWord, error) {
	rows, err := r.DB.Query(`
        SELECT id, user_id, language, type, word, score, last_reviewed
        FROM user_words
        WHERE user_id = ? AND score < ?
        ORDER BY score ASC
        LIMIT 20;
    `, userID, maxScore)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []UserWord
	for rows.Next() {
		var v UserWord
		if err := rows.Scan(
			&v.Id, &v.UserId, &v.Language, &v.Type, &v.Word, &v.Score, &v.LastReviewed,
		); err != nil {
			return nil, err
		}
		results = append(results, v)
	}

	return results, nil
}

func (r *userWordRepositoryImpl) GetPaginated(userID int, lang string, offset, limit int) ([]UserWord, error) {
	rows, err := r.DB.Query(`
        SELECT id, user_id, language, type, word, score, last_reviewed
        FROM user_words
        WHERE user_id = ?
		AND language = ?
        ORDER BY id
        LIMIT ? OFFSET ?`,
		userID, lang, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []UserWord

	for rows.Next() {
		var v UserWord
		if err := rows.Scan(
			&v.Id,
			&v.UserId,
			&v.Language,
			&v.Type,
			&v.Word,
			&v.Score,
			&v.LastReviewed,
		); err != nil {
			return nil, err
		}
		list = append(list, v)
	}

	return list, nil
}

func (r *userWordRepositoryImpl) GetUserTotalWords(userId int, targetLanguage string) (int, error) {
	var resp int
	row := r.DB.QueryRow(`
        SELECT COUNT(*)
        FROM user_words WHERE user_id = ? AND language = ?`, userId, targetLanguage)

	if err := row.Scan(&resp); err != nil {
		return 0, err
	}
	return resp, nil
}

func (r *userWordRepositoryImpl) GetUserLearnedWords(userId int, targetLanguage string, scoreTrigger int) (int, error) {
	var resp int
	row := r.DB.QueryRow(`
        SELECT COUNT(*)
        FROM user_words WHERE user_id = ? AND language = ? AND score > ?`, userId, targetLanguage, scoreTrigger)

	if err := row.Scan(&resp); err != nil {
		return 0, err
	}
	return resp, nil
}
