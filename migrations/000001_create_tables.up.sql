PRAGMA foreign_keys = ON;

-- ===========================
-- Table: users
-- ===========================
CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    current_level TEXT,
    known_vocab_count INTEGER DEFAULT 0,
    grammar_mastered_count INTEGER DEFAULT 0,
    most_recent_weak_area TEXT
);

-- ===========================
-- Table: grammar_mastery
-- Track user's mastery per grammar pattern
-- ===========================
CREATE TABLE IF NOT EXISTS grammar_mastery (
    mastery_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    language TEXT NOT NULL,
    pattern TEXT NOT NULL,
    mastery_score REAL DEFAULT 0.0,
    last_reviewed DATETIME DEFAULT CURRENT_TIMESTAMP,
    weakness_flags TEXT DEFAULT '[]',
    times_incorrect INTEGER DEFAULT 0,

    UNIQUE(user_id, pattern),

    FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_grammar_user ON grammar_mastery(user_id);

-- ===========================
-- Table: vocabulary_mastery
-- Track user's mastery per vocabulary word
-- ===========================
CREATE TABLE IF NOT EXISTS vocabulary_mastery (
    mastery_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    language TEXT NOT NULL,
    word TEXT NOT NULL,
    mastery_score REAL DEFAULT 0.0,
    last_reviewed DATETIME DEFAULT CURRENT_TIMESTAMP,
    times_correct INTEGER DEFAULT 0,
    times_incorrect INTEGER DEFAULT 0,

    UNIQUE(user_id, word),

    FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_vocab_user ON vocabulary_mastery(user_id);

-- ===========================
-- Table: lessons
-- Lessons generated for a user
-- ===========================
CREATE TABLE IF NOT EXISTS lessons (
    lesson_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    language TEXT NOT NULL,
    grammar_focus TEXT,
    content TEXT,
    new_vocabulary TEXT DEFAULT '[]',

    FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_lessons_user ON lessons(user_id);

-- ===========================
-- Table: exercises
-- Exercises belonging to a user and optionally a lesson
-- ===========================
CREATE TABLE IF NOT EXISTS exercises (
    exercise_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    lesson_id INTEGER,
    type TEXT,
    sub_type TEXT,
    question_data TEXT,
    user_response TEXT,
    grade INTEGER,
    feedback TEXT,

    FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY(lesson_id) REFERENCES lessons(lesson_id) ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_exercises_user ON exercises(user_id);
CREATE INDEX IF NOT EXISTS idx_exercises_lesson ON exercises(lesson_id);
