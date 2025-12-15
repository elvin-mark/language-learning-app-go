package storage

import "database/sql"

type UserGrammarRepository interface {
	Upsert(g *UserGrammar) error
	GetByID(id int) (*UserGrammar, error)
	GetForUser(userID int) ([]UserGrammar, error)
	GetPaginatedForUser(userID int, lang string, offset, limit int) ([]UserGrammar, error)
	GetLowestBelowScore(userID int, maxScore int) ([]UserGrammar, error)
	SearchByPattern(userID int, lang string, pattern string, offset, limit int) ([]UserGrammar, error)
}

func NewUserGrammarRepository(db *sql.DB) UserGrammarRepository {
	return &userGrammarRepositoryImpl{
		DB: db,
	}
}
