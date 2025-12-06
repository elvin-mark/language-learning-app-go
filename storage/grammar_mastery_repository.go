package storage

import "database/sql"

// -------------------- GRAMMAR MASTERY REPO --------------------

type GrammarMasteryRepository interface {
	Upsert(g *GrammarMastery) error
	GetForUser(userID int) ([]GrammarMastery, error)
	GetPaginatedForUser(userID int, lang string, offset, limit int) ([]GrammarMastery, error)
	GetLowestBelowScore(userID int, maxScore float64) ([]GrammarMastery, error)
	SearchByPattern(userID int, lang string, pattern string, offset, limit int) ([]GrammarMastery, error)
}

func NewGrammarMasteryRepository(db *sql.DB) GrammarMasteryRepository {
	return &grammarMasteryRepositoryImpl{
		DB: db,
	}
}
