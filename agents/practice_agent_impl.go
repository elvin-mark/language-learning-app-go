package agents

import (
	"language-learning-app/core/llm"
	"language-learning-app/storage"
	"strings"
)

type practiceAgent struct {
	llm                llm.Llm
	lessonRepository   storage.LessonRepository
	exerciseRepository storage.ExerciseRepository
}

func (pa *practiceAgent) GeneratePractice(lang string, practicePattern PracticePattern, lessonId int) (generatedExercise storage.Exercise, err error) {
	lesson, err := pa.lessonRepository.GetByID(lessonId)
	if err != nil {
		return
	}
	prompt := generatePracticeGenerationPrompt(lang, practicePattern, *lesson)
	resp, err := pa.llm.GetResponse(prompt)
	if err != nil {
		return
	}

	cleaned := strings.TrimPrefix(resp.Choices[0].Message.Content, "```json")
	cleaned = strings.TrimSuffix(cleaned, "```")
	cleaned = strings.TrimSpace(cleaned)

	generatedExercise = storage.Exercise{
		UserID:       lesson.UserID,
		LessonID:     &lesson.LessonID,
		Type:         practicePattern.Type,
		SubType:      practicePattern.SubType,
		QuestionData: cleaned,
	}

	err = pa.exerciseRepository.Create(&generatedExercise)
	return
}
