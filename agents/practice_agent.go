package agents

import (
	"language-learning-app/core/llm"
	"language-learning-app/storage"
)

type PracticePattern struct {
	Type        string
	SubType     string
	Instruction string
}

var PracticePatterns = map[string]PracticePattern{
	"Short Story Comprehension": {
		Type:    "Reading",
		SubType: "Short Story Comprehension",
		Instruction: `The response should be a json object. The fields in the json should be: passage, practice (which is an array of json object and each json object contains these fields: questions, answers).
The fields descriptions are:
passage: Contains the short story
practice: An array of json objects that contains the following fields:
	- question: A string of a possible question that can be asked regarding the passage
	- answer: A string containing the answers to this question
`,
	},
	"Sentence Gap-Fill": {
		Type:    "Writing",
		SubType: "Sentence Gap-Fill",
		Instruction: `The response should be an array of json objects. Each json object should have the following fields: sentence, answer.
The fields descriptions are:
sentence: Provide a sentence in which a part of the text is replaced with a * as a blank.
   - Only one word or phrase should be masked per sentence.
   - The masked part should be essential to the meaning.
   - Keep the sentence natural and grammatically correct.
answer: Provide exactly the text that was masked (the correct word or phrase)`,
	},
	"Translation": {
		Type:    "Writing",
		SubType: "Translation",
		Instruction: `The response should be an array of json objects, Each json object should have the following fields: original_sentence, translated_sentence.
The fields descriptions are:
original_sentence: Provide a random sentence in english.
translated_sentence: Provide the correct translation in the language to be practiced.
	- Ensure the translation is accurate and natural.
	- Keep the meaning and grammar correct.
		`,
	},
	"Flashcards": {
		Type:    "Vocabulary",
		SubType: "Flashcards",
		Instruction: `The response should be an array of json objects, Each json object should have the following fields: front, back.
The fields descriptions are:
front: Provide the word, phrase or sentence in english that will appear on the front of the flashcard. The content should be useful for this exercise
back: Provide the correct translation or meaning in the language to be practiced that will appear on the back of the flashcard.
   - The translation should be accurate and natural.
   - Keep it concise and clear.
		`,
	},
}

type PracticeAgent interface {
	GeneratePractice(lang string, practicePattern PracticePattern, lessonId int) (generatedExercise storage.Exercise, err error)
}

func NewPracticeAgent(llm llm.Llm, lessonRepository storage.LessonRepository, exerciseRepository storage.ExerciseRepository) PracticeAgent {
	return &practiceAgent{
		llm:                llm,
		lessonRepository:   lessonRepository,
		exerciseRepository: exerciseRepository,
	}
}
