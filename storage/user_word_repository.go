package storage

import "database/sql"

type UserWordRepository interface {
	GetByID(id int) (*UserWord, error)
	Upsert(v *UserWord) error
	GetLowestBelowScore(userID int, maxScore int) ([]UserWord, error)
	GetPaginated(userID int, lang string, offset, limit int) ([]UserWord, error)
	GetUserTotalWords(userId int, targetLanguage string) (int, error)
	GetUserLearnedWords(userId int, targetLanguage string, scoreTrigger int) (int, error)
}

func NewUserWordRepository(db *sql.DB) UserWordRepository {
	return &userWordRepositoryImpl{
		DB: db,
	}
}
