package services

import (
	"language-learning-app/core/agents"
	models "language-learning-app/models/lesson"
	"language-learning-app/storage"
)

type UserLessonService interface {
	GetLessons(userId int, lang string, page, pageSize int) (lessons []models.LessonItem, err error)
	GenerateLesson(user *storage.User, grammarId int, wordsId []int) (lesson models.LessonItem, err error)
}

func NewUserLessonService(userLessonRepository storage.UserLessonRepository, lessonAgent agents.LessonAgent, userGrammarRepository storage.UserGrammarRepository, userWordRepository storage.UserWordRepository) UserLessonService {
	return &userLessonServiceImpl{
		userLessonRepository:  userLessonRepository,
		userGrammarRepository: userGrammarRepository,
		userWordRepository:    userWordRepository,
		lessonAgent:           lessonAgent,
	}
}
