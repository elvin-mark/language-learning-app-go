PRAGMA foreign_keys = ON;

-- ===========================
-- Table: users
-- ===========================
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    preferred_language TEXT NOT NULL DEFAULT 'English',
    target_language TEXT NOT NULL DEFAULT 'Korean'
);

-- ===========================
-- Table: user_grammar
-- Track user's grammar per grammar pattern
-- ===========================
CREATE TABLE IF NOT EXISTS user_grammar (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    language TEXT NOT NULL,
    pattern TEXT NOT NULL,
    score INTEGER DEFAULT 0,
    last_reviewed DATETIME DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(user_id, pattern),

    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_grammar_user ON user_grammar(user_id);

-- ===========================
-- Table: user_words
-- Track user's words per vocabulary word
-- ===========================
CREATE TABLE IF NOT EXISTS user_words (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    language TEXT NOT NULL,
    type TEXT NOT NULL, -- VERB, NOUN, ADJ, ADV
    word TEXT NOT NULL,
    score INTEGER DEFAULT 0,
    last_reviewed DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, word),

    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_vocab_user ON user_words(user_id);

-- ===========================
-- Table: lessons
-- Lessons generated for a user
-- ===========================
CREATE TABLE IF NOT EXISTS user_lessons (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    language TEXT NOT NULL,
    grammar_id INTEGER NOT NULL,
    words_id TEXT DEFAULT '[]',
    content TEXT,
    sample_sentences TEXT,
    words_meaning TEXT,

    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_lessons_user ON user_lessons(id);
