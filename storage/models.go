package storage

import "time"

// -----------------------
// Users
// -----------------------
type User struct {
	Id                int    `db:"id"`
	Username          string `db:"username"`
	Password          string `db:"password"`
	PreferredLanguage string `db:"preferred_language"`
	TargetLanguage    string `db:"target_language"`
}

// -----------------------
// User Grammar
// -----------------------
type UserGrammar struct {
	Id           int       `db:"id"`
	UserId       int       `db:"user_id"`
	Language     string    `db:"language"`
	Pattern      string    `db:"pattern"`
	Score        int       `db:"score"`
	LastReviewed time.Time `db:"last_reviewed"`
}

// -----------------------
// User Words
// -----------------------
type UserWord struct {
	Id           int       `db:"id"`
	UserId       int       `db:"user_id"`
	Language     string    `db:"language"`
	Type         string    `db:"type"`
	Word         string    `db:"word"`
	Score        int       `db:"score"`
	LastReviewed time.Time `db:"last_reviewed"`
}

// -----------------------
// User Lessons
// -----------------------
type UserLesson struct {
	Id              int    `db:"id"`
	UserId          int    `db:"user_id"`
	Language        string `db:"language"`
	GrammarId       int    `db:"grammar_id"`
	WordsId         []int  `db:"words_id"`
	Content         string `db:"content"`
	SampleSentences string `db:"sample_sentences"`
	WordsMeaning    string `db:"words_meaning"`
}
