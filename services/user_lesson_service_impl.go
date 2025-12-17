package services

import (
	"fmt"
	"language-learning-app/agents"
	models "language-learning-app/models/lesson"
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

func (ls *userLessonServiceImpl) GetLessons(userId int, lang string, page, pageSize int) (lessons []models.LessonItem, err error) {
	lessonEntities, err := ls.userLessonRepository.GetPaginatedByLanguageAndUser(userId, lang, pageSize, page*pageSize)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	lessons = make([]models.LessonItem, 0)
	for _, entity := range lessonEntities {
		grammar, err := ls.userGrammarRepository.GetByID(entity.GrammarId)
		if err != nil {
			utils.Logger.Error(err.Error())
			continue
		}
		words := make([]string, 0)
		for _, wordId := range entity.WordsId {
			word, err := ls.userWordRepository.GetByID(wordId)
			if err != nil {
				utils.Logger.Error(err.Error())
				continue
			}
			words = append(words, word.Word)
		}
		lessons = append(lessons, models.LessonItem{
			Id:              entity.Id,
			UserId:          entity.UserId,
			Language:        entity.Language,
			Grammar:         grammar.Pattern,
			Words:           words,
			Content:         entity.Content,
			SampleSentences: entity.SampleSentences,
			WordsMeaning:    entity.WordsMeaning,
		})
	}
	return
}

func (ls *userLessonServiceImpl) GenerateLesson(user *storage.User) (lesson models.LessonItem, err error) {
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

	lessonEntity := storage.UserLesson{
		UserId:          user.Id,
		Language:        user.TargetLanguage,
		GrammarId:       randomGrammar.Id,
		WordsId:         wordsId,
		Content:         generatedLesson.Content,
		SampleSentences: sampleSentencesJSON,
		WordsMeaning:    wordsMeaning,
	}

	err = ls.userLessonRepository.Create(&lessonEntity)

	lesson = models.LessonItem{
		Id:              lessonEntity.Id,
		UserId:          lessonEntity.UserId,
		Language:        lessonEntity.Language,
		Grammar:         randomGrammar.Pattern,
		Words:           wordsList,
		Content:         lessonEntity.Content,
		SampleSentences: lessonEntity.SampleSentences,
		WordsMeaning:    lessonEntity.WordsMeaning,
	}
	return
}
