package agents

type UsageGrade struct {
	Score    int    `json:"score"`
	Feedback string `json:"feedback"`
}

type GeneratedLesson struct {
	Content         string            `json:"content"`
	SampleSentences []string          `json:"sample_sentences"`
	WordsMeaning    map[string]string `json:"words_meaning"`
}

type GeneratedTranslationExercise struct {
	Sentences []string `json:"sentences"`
}
