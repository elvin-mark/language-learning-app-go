package agents

type GeneratedLesson struct {
	Content         string `json:"content"`
	SampleSentences string `json:"sample_sentences"`
	WordsMeaning    string `json:"words_meaning"`
}

type GeneratedTranslationExercise struct {
	Sentences []string `json:"sentences"`
}
