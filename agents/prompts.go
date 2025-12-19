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

const gradeTranslationGenerationPrompt = `You are an expert and strict %s language teacher. Your task is to grade and give a feedback to the translation created by the user. Your response should be solely a JSON object, no extra text.
**Input:**
- **Original Sentence:** %s
- **User's translation:** %s

**Response:**
- **score**: From 0-5, grade how well the user's sentence has translated the sentence.
- **feedback**: Provide a feedback on how well the user has translated the sentence and what the user should fix or improve.
`

func generateGradeTranslationPrompt(language, originalSentence, translatedSentence string) string {
	return fmt.Sprintf(gradeTranslationGenerationPrompt, language, originalSentence, translatedSentence)
}

const readingComprehensionExerciseGenerationPrompt = `You are a creative %s language teacher. Generate a short passage with questions so the user can practice the grammar pattern and vocabulary in this lesson. Your response should be a JSON object.

**Exercise Goal:**
- **Grammar Pattern**: %s
- **Vocabulary**: %s

**Instructions:**
- Create a short passage (short text) in %s language that helps the user to practice the grammar pattern and vocabulary of this lesson. Also generates some questions regarding the text so the user can reply and practice the use of the language.

**Response:**
- short_text: a short passage of any topic that is more likely to use the grammar pattern and vocabulary
- questions: an array of strings in which each string is a question regarding the short_text

Make sure to reply just with the JSON object, no need for other text.`

func generateReadingComprehensionExercisePrompt(targetLanguage string, grammarPattern string, words []string) string {
	vocab := strings.Join(words, ",")
	return fmt.Sprintf(readingComprehensionExerciseGenerationPrompt, targetLanguage, grammarPattern, vocab, targetLanguage)
}

const dialogueInitGenerationPrompt = `You are a creative %s language teacher. Generate a situation in which the user can engage in dialogue in which it can use the grammar pattern and vocabulary in this lesson.

**Exercise Goal:**
- **Grammar Pattern**: %s
- **Vocabulary**: %s

**Instructions:**
- Create a situation in which the user will need to interact with a couple questions and answers so the user can practice the grammar pattern and vocabulary.

**Response:**
- situation: Detail description of the situation so the user can understand the context
- init: First sentence or question so the conversation can start

Make sure to reply just with the JSON object, no need for other text.`

func generateDialogueInitExercisePrompt(targetLanguage string, grammarPattern string, words []string) string {
	vocab := strings.Join(words, ",")
	return fmt.Sprintf(dialogueInitGenerationPrompt, targetLanguage, grammarPattern, vocab)
}

const dialogueContinuationGenerationPrompt = `You are a creative %s language teacher. Continue with this dialogue in which the user is trying to use the grammar pattern and vocabulary in this lesson.

**Exercise Goal:**
- **Grammar Pattern**: %s
- **Vocabulary**: %s

**Instructions:**
- Create a proper response or/and ask some question so the user keep engaged in this dialogue and facilitate the usage of the grammar pattern and vocabulary in the user's response

**Dialogue History:**
%s

**Response:**
- next: Next sentence or question so the conversation can continue. If you consider that this dialogue is long enough you can end it properly in this sentence.

Make sure to reply just with the JSON object, no need for other text.`

func generateDialogueContinuationExercisePrompt(targetLanguage string, grammarPattern string, words []string, history string) string {
	vocab := strings.Join(words, ",")
	return fmt.Sprintf(dialogueContinuationGenerationPrompt, targetLanguage, grammarPattern, vocab, history)
}

const chatbotGenerationPrompt = `You are a creative %s language teacher. Answer to the user's question in detail and with examples.

**Instructions:**
- Create a proper response that can clear any doubts that the user has regarding the question. Give samples, scenarios, usages or anything you consider that will be helpful so the user has a clear understanding.

**Question:** %s

**Response:**
- response: your response to the user's question

Make sure to reply just with the JSON object, no need for other text.`

func generateChatbotPrompt(targetLanguage string, question string) string {
	return fmt.Sprintf(chatbotGenerationPrompt, targetLanguage, question)
}
