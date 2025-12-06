package services

import (
	"language-learning-app/agents"
	dto "language-learning-app/dto/agents"
	"language-learning-app/storage"
)

type AgentService interface {
	GenerateLesson(req dto.GenerateLessonRequest) (lesson storage.Lesson, err error)
	GenerateExercise(req dto.GeneraterExerciseRequest) (lesson storage.Exercise, err error)
}

func NewAgentService(lessonAgent agents.LessonAgent, practiceAgent agents.PracticeAgent) AgentService {
	return &agentServiceImpl{
		lessonAgent:   lessonAgent,
		practiceAgent: practiceAgent,
	}
}
