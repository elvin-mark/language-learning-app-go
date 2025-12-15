package services

import (
	"language-learning-app/agents"
	"language-learning-app/storage"
)

type UserLessonService interface {
	GetLessons(userId int, lang string, page, pageSize int) (lessons []storage.UserLesson, err error)
	GenerateLesson(user *storage.User) (lesson storage.UserLesson, err error)
}

func NewUserLessonService(userLessonRepository storage.UserLessonRepository, lessonAgent agents.LessonAgent, userGrammarRepository storage.UserGrammarRepository, userWordRepository storage.UserWordRepository) UserLessonService {
	return &userLessonServiceImpl{
		userLessonRepository:  userLessonRepository,
		userGrammarRepository: userGrammarRepository,
		userWordRepository:    userWordRepository,
		lessonAgent:           lessonAgent,
	}
}
