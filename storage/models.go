package storage

import "time"

// -----------------------
// Users
// -----------------------
type User struct {
	UserID               int     `db:"user_id"`
	CurrentLevel         string  `db:"current_level"`
	KnownVocabCount      int     `db:"known_vocab_count"`
	GrammarMasteredCount int     `db:"grammar_mastered_count"`
	MostRecentWeakArea   *string `db:"most_recent_weak_area"`
}

// -----------------------
// Grammar Mastery
// -----------------------
type GrammarMastery struct {
	MasteryID      int       `db:"mastery_id"`
	UserID         int       `db:"user_id"`
	Language       string    `db:"language"`
	Pattern        string    `db:"pattern"`
	MasteryScore   float64   `db:"mastery_score"`
	LastReviewed   time.Time `db:"last_reviewed"`
	WeaknessFlags  []string  `db:"weakness_flags"` // JSON array
	TimesIncorrect int       `db:"times_incorrect"`
}

// -----------------------
// Vocabulary Mastery
// -----------------------
type VocabularyMastery struct {
	MasteryID      int       `db:"mastery_id"`
	UserID         int       `db:"user_id"`
	Language       string    `db:"language"`
	Word           string    `db:"word"`
	MasteryScore   float64   `db:"mastery_score"`
	LastReviewed   time.Time `db:"last_reviewed"`
	TimesCorrect   int       `db:"times_correct"`
	TimesIncorrect int       `db:"times_incorrect"`
}

// -----------------------
// Lessons
// -----------------------
type Lesson struct {
	LessonID      int      `db:"lesson_id"`
	UserID        int      `db:"user_id"`
	Language      string   `db:"language"`
	GrammarFocus  string   `db:"grammar_focus"`
	Content       string   `db:"content"`
	NewVocabulary []string `db:"new_vocabulary"` // JSON array
}

// -----------------------
// Exercises
// -----------------------
type Exercise struct {
	ExerciseID   int     `db:"exercise_id"`
	UserID       int     `db:"user_id"`
	LessonID     *int    `db:"lesson_id"`
	Type         string  `db:"type"`
	SubType      string  `db:"sub_type"`
	QuestionData string  `db:"question_data"`
	UserResponse string  `db:"user_response"`
	Grade        *int    `db:"grade"`
	Feedback     *string `db:"feedback"`
}
