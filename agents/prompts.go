package agents

import (
	"fmt"
	"strings"
)

const lessonGenerationPrompt = `You are an expert and friendly %s language teacher. Your task is to create a concise, personalized lesson and return it as a JSON object.

**Topic:**
- **Grammar Pattern to Learn:** %s
- **Vocabulary to Learn:** %s

**Lesson Requirements:**
1.  **content**: Write a clear and simple explanation of the grammar pattern.
2.  **sample_sentences**: Create 3-4 diverse and practical example sentences that use the grammar pattern correctly and also use the words in the vocabulary to learn.
3.  **words_meaning**: Explain the meaning of each word in the vocabulary to learn.

Make sure to reply just with the JSON object, no need for other text. Each field of this JSON object are just strings (Mardown style strings)`

func generateLessonGenerationPrompt(lang string, grammarPattern string, words []string) string {
	return fmt.Sprintf(lessonGenerationPrompt, lang, grammarPattern, strings.Join(words, ","))
}

const translationExerciseGenerationPrompt = `You are a creative %s language teacher. Generate a sequence of sentences in %s language as a JSON object.

**Exercise Goal:**
- **Grammar Pattern**: %s
- **Vocabulary**: %s

**Instructions:**
- Create a list of sentences in %s language that helps the user to practice the grammar pattern and vocabulary of %s language by trying to translate it.

**Response:**
- sentences: an array of strings with the sentences
`

func generateTranslationExercisePrompt(preferredLanguage, targetLanguage string, grammarPattern string, words []string) string {
	vocab := strings.Join(words, ",")
	return fmt.Sprintf(translationExerciseGenerationPrompt, targetLanguage, preferredLanguage, grammarPattern, vocab, preferredLanguage, targetLanguage)
}
