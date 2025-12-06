package services

import "language-learning-app/storage"

type LessonService interface {
	GetLessons(userId int, lang string, page, pageSize int) (lessons []storage.Lesson, err error)
	GetLessonsByGrammar(userId int, lang string, grammarPattern string, page, pageSize int) (lessons []storage.Lesson, err error)
}

func NewLessonService(lessonRepository storage.LessonRepository) LessonService {
	return &lessonServiceImpl{
		lessonRepository: lessonRepository,
	}
}
