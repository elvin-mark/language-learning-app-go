package services

import "language-learning-app/storage"

type lessonServiceImpl struct {
	lessonRepository storage.LessonRepository
}

func (ls *lessonServiceImpl) GetLessons(userId int, lang string, page, pageSize int) (lessons []storage.Lesson, err error) {
	lessons, err = ls.lessonRepository.GetPaginatedByLanguageAndUser(userId, lang, pageSize, page*pageSize)
	return
}

func (ls *lessonServiceImpl) GetLessonsByGrammar(userId int, lang string, grammarPattern string, page, pageSize int) (lessons []storage.Lesson, err error) {
	lessons, err = ls.lessonRepository.SearchLessonsByGrammarFocus(userId, lang, grammarPattern, pageSize, page*pageSize)
	return
}
