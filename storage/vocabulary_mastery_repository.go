package storage

import "database/sql"

// -------------------- VOCABULARY MASTERY REPO --------------------

type VocabularyMasteryRepository interface {
	Upsert(v *VocabularyMastery) error
	GetLowestBelowScore(userID int, maxScore float64) ([]VocabularyMastery, error)
	GetPaginated(userID int, lang string, offset, limit int) ([]VocabularyMastery, error)
}

func NewVocabularyMasteryRepository(db *sql.DB) VocabularyMasteryRepository {
	return &vocabularyMasteryRepositoryImpl{
		DB: db,
	}
}
