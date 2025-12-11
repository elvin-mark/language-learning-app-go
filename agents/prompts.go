package agents

import (
	"fmt"
	"language-learning-app/storage"
	"strings"
)

const lessonGenerationPrompt = `You are an expert and friendly %s language teacher. Your task is to create a concise, personalized lesson and return it as a JSON object.

**User's Current Status:**
- **Grammar Pattern to Learn:** %s
- **Current Mastery Score:** %f (A score from 0.0 to 1.0)
- **Known Issues/Weakness Flags:** %s (These are specific errors the user has made before. Address them in your explanation.)

**Lesson Requirements:**
1.  **explanation_text**: Write a clear and simple explanation of the grammar pattern. If there are weakness flags, provide examples that specifically correct those mistakes.
2.  **example_sentences**: Create 3-4 diverse and practical example sentences that use the grammar pattern correctly.
3.  **new_vocabulary**: Naturally include some of these **new vocabulary words** (%s) within your example sentences.

Fill the grammar_pattern and new_vocabulary fields in the output with the exact data provided.`

func generateLessonGenerationPrompt(lang string, grammar storage.GrammarMastery, vocabs []storage.VocabularyMastery) string {
	newWords := make([]string, 0)
	for _, vocab := range vocabs {
		newWords = append(newWords, vocab.Word)
	}
	return fmt.Sprintf(lessonGenerationPrompt, lang, grammar.Pattern, grammar.MasteryScore, strings.Join(grammar.WeaknessFlags, ","), strings.Join(newWords, ","))
}

const practiceGenerationPrompt = `You are a creative %s language teacher. Generate a single practice exercise as a JSON object.

**Exercise Goal:**
- **Type**: %s
- **Sub-Type**: %s
- **Grammar Pattern**: %s
- **Vocabulary**: %s

**Instructions:**
- Create a proper exercise to review %s using %s kind of exercises to make the user practice %s.
- Make sure that the exercise require using the vocabulary: %s.
- The type and sub_type in the output must match the goal.
- The output should contain the following fields: %s
- The description of the fields are: %s`

func generatePracticeGenerationPrompt(lang string, practicePattern PracticePattern, lesson storage.Lesson) string {
	words := strings.Join(lesson.NewVocabulary, ",")
	fields := strings.Join(practicePattern.Fields, ",")
	return fmt.Sprintf(practiceGenerationPrompt, lang, practicePattern.Type, practicePattern.SubType, lesson.GrammarFocus, words, practicePattern.Type, practicePattern.SubType, lesson.GrammarFocus, words, fields, practicePattern.FielsDescription)
}
