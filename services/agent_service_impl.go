package services

import (
	"language-learning-app/agents"
	dto "language-learning-app/dto/agents"
	"language-learning-app/storage"
	"language-learning-app/utils"
)

type agentServiceImpl struct {
	lessonAgent   agents.LessonAgent
	practiceAgent agents.PracticeAgent
}

func (as *agentServiceImpl) GenerateLesson(req dto.GenerateLessonRequest) (lesson storage.Lesson, err error) {
	lesson, err = as.lessonAgent.GenerateLesson(req.UserId, req.Lang, 0.7)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	return
}

func (as *agentServiceImpl) GenerateExercise(req dto.GeneraterExerciseRequest) (lesson storage.Exercise, err error) {

	lesson, err = as.practiceAgent.GeneratePractice(req.Lang, agents.PracticePatterns[req.PracticePattern], req.LessonId)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	return
}
