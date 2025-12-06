package storage

import (
	"database/sql"
	"errors"
)

// -------------------- USERS STATUS REPO --------------------

type userRepositoryImpl struct {
	DB *sql.DB
}

func (r *userRepositoryImpl) Create(user *User) error {
	res, err := r.DB.Exec(`
        INSERT INTO users (current_level, known_vocab_count, grammar_mastered_count, most_recent_weak_area)
        VALUES (?, ?, ?, ?)`,
		user.CurrentLevel, user.KnownVocabCount, user.GrammarMasteredCount, user.MostRecentWeakArea)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	user.UserID = int(id)
	return nil
}

func (r *userRepositoryImpl) GetByID(id int) (*User, error) {
	var u User
	row := r.DB.QueryRow(`
        SELECT user_id, current_level, known_vocab_count, grammar_mastered_count, most_recent_weak_area
        FROM users WHERE user_id = ?`, id)

	if err := row.Scan(&u.UserID, &u.CurrentLevel, &u.KnownVocabCount,
		&u.GrammarMasteredCount, &u.MostRecentWeakArea); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
