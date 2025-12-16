package services

import (
	"fmt"
	"language-learning-app/agents"
	"language-learning-app/storage"
	"language-learning-app/utils"
	"math/rand"
)

type userLessonServiceImpl struct {
	userLessonRepository  storage.UserLessonRepository
	userGrammarRepository storage.UserGrammarRepository
	userWordRepository    storage.UserWordRepository
	lessonAgent           agents.LessonAgent
}

func (ls *userLessonServiceImpl) GetLessons(userId int, lang string, page, pageSize int) (lessons []storage.UserLesson, err error) {
	lessons, err = ls.userLessonRepository.GetPaginatedByLanguageAndUser(userId, lang, pageSize, page*pageSize)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	return
}

func (ls *userLessonServiceImpl) GenerateLesson(user *storage.User) (lesson storage.UserLesson, err error) {
	grammars, err := ls.userGrammarRepository.GetLowestBelowScore(user.Id, 70)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	if len(grammars) < 1 {
		err = fmt.Errorf("no grammars below score")
		utils.Logger.Error(err.Error())
		return
	}

	words, err := ls.userWordRepository.GetLowestBelowScore(user.Id, 70)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	if len(words) < 1 {
		err = fmt.Errorf("no words below score")
		utils.Logger.Error(err.Error())
		return
	}

	randomGrammar := grammars[rand.Intn(len(grammars))]

	wordsList := make([]string, 0)
	wordsId := make([]int, 0)
	for _, word := range words {
		wordsList = append(wordsList, word.Word)
		wordsId = append(wordsId, word.Id)
	}
	generatedLesson, err := ls.lessonAgent.GenerateLesson(user.TargetLanguage, randomGrammar.Pattern, wordsList)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	sampleSentencesJSON, err := utils.ToJSON(generatedLesson.SampleSentences)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	wordsMeaning, err := utils.ToJSON(generatedLesson.WordsMeaning)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	lesson = storage.UserLesson{
		UserId:          user.Id,
		Language:        user.TargetLanguage,
		GrammarId:       randomGrammar.Id,
		WordsId:         wordsId,
		Content:         generatedLesson.Content,
		SampleSentences: sampleSentencesJSON,
		WordsMeaning:    wordsMeaning,
	}

	err = ls.userLessonRepository.Create(&lesson)
	return
}
