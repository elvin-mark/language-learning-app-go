package storage

import (
	"database/sql"
	"errors"
)

type userGrammarRepositoryImpl struct {
	DB *sql.DB
}

func (r *userGrammarRepositoryImpl) Upsert(g *UserGrammar) error {
	_, err := r.DB.Exec(`
        INSERT INTO user_grammar (user_id, language, pattern, score)
        VALUES (?, ?, ?, ?)
        ON CONFLICT(user_id, pattern)
        DO UPDATE SET 
            score = excluded.score,
            last_reviewed = CURRENT_TIMESTAMP;
    `, g.UserId, g.Language, g.Pattern, g.Score)

	return err
}

func (r *userGrammarRepositoryImpl) GetByID(id int) (*UserGrammar, error) {
	var u UserGrammar
	row := r.DB.QueryRow(`
        SELECT id, user_id, language, pattern, score, last_reviewed
        FROM user_grammar WHERE id = ?`, id)

	if err := row.Scan(&u.Id, &u.UserId, &u.Language, &u.Pattern, &u.Score, &u.LastReviewed); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *userGrammarRepositoryImpl) GetForUser(userID int) ([]UserGrammar, error) {
	rows, err := r.DB.Query(`
        SELECT id, user_id, language, pattern, score, last_reviewed
        FROM user_grammar WHERE user_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []UserGrammar
	for rows.Next() {
		var gm UserGrammar

		if err := rows.Scan(&gm.Id, &gm.UserId, &gm.Language, &gm.Pattern, &gm.Score, &gm.LastReviewed); err != nil {
			return nil, err
		}

		list = append(list, gm)
	}
	return list, nil
}

func (r *userGrammarRepositoryImpl) GetPaginatedForUser(userID int, lang string, offset, limit int) ([]UserGrammar, error) {
	rows, err := r.DB.Query(`
        SELECT id, user_id, language, pattern, score, last_reviewed
        FROM user_grammar 
        WHERE user_id = ?
		AND language = ?
        ORDER BY id
        LIMIT ? OFFSET ?`, userID, lang, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []UserGrammar
	for rows.Next() {
		var gm UserGrammar

		if err := rows.Scan(&gm.Id, &gm.UserId, &gm.Language, &gm.Pattern, &gm.Score, &gm.LastReviewed); err != nil {
			return nil, err
		}

		list = append(list, gm)
	}
	return list, nil
}

func (r *userGrammarRepositoryImpl) GetLowestBelowScore(userID int, maxScore int) ([]UserGrammar, error) {
	rows, err := r.DB.Query(`
        SELECT id, user_id, language, pattern, score, last_reviewed
        FROM user_grammar
        WHERE user_id = ? AND score < ?
        ORDER BY score ASC
        LIMIT 20;
    `, userID, maxScore)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []UserGrammar

	for rows.Next() {
		var g UserGrammar

		if err := rows.Scan(
			&g.Id, &g.UserId, &g.Language, &g.Pattern, &g.Score, &g.LastReviewed,
		); err != nil {
			return nil, err
		}

		results = append(results, g)
	}

	return results, nil
}

func (r *userGrammarRepositoryImpl) SearchByPattern(userID int, lang string, pattern string, offset, limit int) ([]UserGrammar, error) {
	likePattern := "%" + pattern + "%"

	rows, err := r.DB.Query(`
        SELECT id, user_id, language, pattern, score, last_reviewed
        FROM user_grammar
        WHERE user_id = ? AND pattern LIKE ?
		AND language = ?
        ORDER BY id
        LIMIT ? OFFSET ?`, userID, likePattern, lang, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []UserGrammar
	for rows.Next() {
		var gm UserGrammar

		if err := rows.Scan(&gm.Id, &gm.UserId, &gm.Language, &gm.Pattern, &gm.Score, &gm.LastReviewed); err != nil {
			return nil, err
		}

		list = append(list, gm)
	}
	return list, nil
}
