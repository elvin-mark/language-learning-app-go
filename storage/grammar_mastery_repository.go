package storage

import "database/sql"

// -------------------- GRAMMAR MASTERY REPO --------------------

type GrammarMasteryRepository struct {
	DB *sql.DB
}

func (r *GrammarMasteryRepository) Upsert(g *GrammarMastery) error {
	flagsJSON, _ := toJSON(g.WeaknessFlags)

	_, err := r.DB.Exec(`
        INSERT INTO grammar_mastery (user_id, language, pattern, mastery_score, weakness_flags, times_incorrect)
        VALUES (?, ?, ?, ?, ?, ?)
        ON CONFLICT(user_id, pattern)
        DO UPDATE SET 
            mastery_score = excluded.mastery_score,
            weakness_flags = excluded.weakness_flags,
            times_incorrect = excluded.times_incorrect,
            last_reviewed = CURRENT_TIMESTAMP;
    `, g.UserID, g.Language, g.Pattern, g.MasteryScore, flagsJSON, g.TimesIncorrect)

	return err
}

func (r *GrammarMasteryRepository) GetForUser(userID int) ([]GrammarMastery, error) {
	rows, err := r.DB.Query(`
        SELECT mastery_id, user_id, language, pattern, mastery_score, last_reviewed, weakness_flags, times_incorrect
        FROM grammar_mastery WHERE user_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []GrammarMastery
	for rows.Next() {
		var gm GrammarMastery
		var flags string

		if err := rows.Scan(&gm.MasteryID, &gm.UserID, &gm.Language, &gm.Pattern,
			&gm.MasteryScore, &gm.LastReviewed, &flags, &gm.TimesIncorrect); err != nil {
			return nil, err
		}

		fromJSON(flags, &gm.WeaknessFlags)
		list = append(list, gm)
	}
	return list, nil
}

func (r *GrammarMasteryRepository) GetLowestBelowScore(userID int, maxScore float64) ([]GrammarMastery, error) {
	rows, err := r.DB.Query(`
        SELECT mastery_id, user_id, language, pattern, mastery_score, last_reviewed,
               weakness_flags, times_incorrect
        FROM grammar_mastery
        WHERE user_id = ? AND mastery_score < ?
        ORDER BY mastery_score ASC
        LIMIT 20;
    `, userID, maxScore)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []GrammarMastery

	for rows.Next() {
		var g GrammarMastery
		var flagsStr string

		if err := rows.Scan(
			&g.MasteryID, &g.UserID, &g.Language, &g.Pattern, &g.MasteryScore,
			&g.LastReviewed, &flagsStr, &g.TimesIncorrect,
		); err != nil {
			return nil, err
		}

		// Parse JSON string → []string
		_ = fromJSON(flagsStr, &g.WeaknessFlags)

		results = append(results, g)
	}

	return results, nil
}
