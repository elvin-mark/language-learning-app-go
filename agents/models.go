package agents

type GeneratedLesson struct {
	GrammarPattern   string   `json:"grammar_pattern"`
	NewVocabulary    []string `json:"new_vocabulary"`
	ExplanationText  string   `json:"explanation_text"`
	ExampleSentences []string `json:"example_sentences"`
}
