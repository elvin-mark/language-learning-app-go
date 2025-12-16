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
2.  **sample_sentences**: Create 3-4 diverse and practical example sentences that use the grammar pattern correctly and also use the words in the vocabulary to learn. This should be an array of strings in which each string is a sample sentence
3.  **words_meaning**: Explain the meaning of each word in the vocabulary to learn. This should be a JSON object where the key is the word and the value is the meaning

Make sure to reply just with the JSON object, no need for other text.`

func generateLessonGenerationPrompt(lang string, grammarPattern string, words []string) string {
	return fmt.Sprintf(lessonGenerationPrompt, lang, grammarPattern, strings.Join(words, ","))
}

const gradeUsageGenerationPrompt = `You are an expert and strict %s language teacher. Your task is to grade and give a feedback to the sentence created by the user. Your response should be solely a JSON object, no extra text.
**Input:**
- **User's sentence:** %s
- **Grammar Pattern or Word to evaluate usage:** %s

**Response:**
- **score**: From 0-5, grade how well the user's sentence is using the grammar pattern or word. If the grammar pattern or word is not used return 0 as score.
- **feedback**: Provide a feedback on how well the user was using the grammar pattern or word and what the user should fix or improve. Don't provide any feedback if the grammar pattern or word is not in used in the sentence

If the user does not use the grammar pattern or word in the provided sentence (be very strict with this), response with an empty JSON object
`

func generateGradeUsagePrompt(language, sentence, grammarPatternOrWord string) string {
	return fmt.Sprintf(gradeUsageGenerationPrompt, language, sentence, grammarPatternOrWord)
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
