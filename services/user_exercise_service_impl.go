package services

import (
	"language-learning-app/agents"
	"language-learning-app/storage"
	"language-learning-app/utils"
)

type exerciseServiceImpl struct {
	exerciseAgent         agents.ExerciseAgent
	userLessonRepository  storage.UserLessonRepository
	userGrammarRepository storage.UserGrammarRepository
	userWordRepository    storage.UserWordRepository
}

func (es *exerciseServiceImpl) GenerateTranslationExercise(user *storage.User, lessonId int) (exercise agents.GeneratedTranslationExercise, err error) {
	lesson, err := es.userLessonRepository.GetByID(lessonId)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	grammar, err := es.userGrammarRepository.GetByID(lesson.GrammarId)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	wordsId := lesson.WordsId
	words := make([]string, 0)
	var word *storage.UserWord
	for _, id := range wordsId {
		word, err = es.userWordRepository.GetByID(id)
		if err != nil {
			utils.Logger.Error(err.Error())
			break
		}
		words = append(words, word.Word)
	}
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	exercise, err = es.exerciseAgent.GenerateTranslationExercise(user.PreferredLanguage, user.TargetLanguage, grammar.Pattern, words)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	return
}
