package agents

import (
	"language-learning-app/core/llm"
	"language-learning-app/storage"
)

type PracticePattern struct {
	Type             string
	SubType          string
	Fields           []string
	FielsDescription string
}

var PracticePatterns = map[string]PracticePattern{
	"Short Story Comprehension": {
		Type:    "Reading",
		SubType: "Short Story Comprehension",
		Fields:  []string{"passage", "questions", "answers"},
		FielsDescription: `
passage: Contains the short story
questions: An array of strings of all possible questions that can be asked regarding the passage
answers: An array of strings containing the answers to the questions`,
	},
	"Sentence Gap-Fill": {
		Type:    "Writing",
		SubType: "Sentence Gap-Fill",
		Fields:  []string{"sentence", "answer"},
		FielsDescription: `
sentence: Provide a sentence in which a part of the text is replaced with a * as a blank.
   - Only one word or phrase should be masked per sentence.
   - The masked part should be essential to the meaning.
   - Keep the sentence natural and grammatically correct.
answer: Provide exactly the text that was masked (the correct word or phrase)`,
	},
	"Translation": {
		Type:    "Writing",
		SubType: "Translation",
		Fields:  []string{"original_sentence", "translated_sentence"},
		FielsDescription: `
original_sentence: Provide a random sentence in english.
translated_sentence: Provide the correct translation in the language to be practiced.
	- Ensure the translation is accurate and natural.
	- Keep the meaning and grammar correct.`,
	},
	"Flashcards": {
		Type:    "Vocabulary",
		SubType: "Flashcards",
		Fields:  []string{"front", "back"},
		FielsDescription: `
front: Provide the word, phrase or sentence in english that will appear on the front of the flashcard. The content should be useful for this exercise
back: Provide the correct translation or meaning in the language to be practiced that will appear on the back of the flashcard.
   - The translation should be accurate and natural.
   - Keep it concise and clear.`,
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
