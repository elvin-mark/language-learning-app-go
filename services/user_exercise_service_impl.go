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

func (es *exerciseServiceImpl) GradeUsage(user *storage.User, sentence, grammarPatternOrWord string) (grade agents.UsageGrade, err error) {
	grade, err = es.exerciseAgent.GradeUsage(user.TargetLanguage, sentence, grammarPatternOrWord)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	return
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

func (es *exerciseServiceImpl) GradeTranslationExercise(user *storage.User, lessonId int, originallSentence, translatedSentence string) (grades []agents.UsageGrade, err error) {
	grades = make([]agents.UsageGrade, 0)
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

	gradeGrammarUsage, err := es.exerciseAgent.GradeUsage(user.TargetLanguage, translatedSentence, grammar.Pattern)
	if err != nil {
		utils.Logger.Error(err.Error())
	} else if gradeGrammarUsage.Score > 0 {
		grades = append(grades, gradeGrammarUsage)
		grammar.Score += gradeGrammarUsage.Score
		err = es.userGrammarRepository.Upsert(grammar)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
	}

	for _, id := range lesson.WordsId {
		word, err := es.userWordRepository.GetByID(id)
		if err != nil {
			utils.Logger.Error(err.Error())
			continue
		}
		gradeWordUsage, err := es.exerciseAgent.GradeUsage(user.TargetLanguage, translatedSentence, word.Word)
		if err != nil {
			utils.Logger.Error(err.Error())
		} else if gradeWordUsage.Score > 0 {
			grades = append(grades, gradeWordUsage)
			word.Score += gradeWordUsage.Score
			err = es.userWordRepository.Upsert(word)
			if err != nil {
				utils.Logger.Error(err.Error())
			}
		}
	}

	translationGrade, err := es.exerciseAgent.GradeTranslation(user.TargetLanguage, originallSentence, translatedSentence)
	if err != nil {
		utils.Logger.Error(err.Error())
	} else {
		grades = append(grades, translationGrade)
	}
	return
}

func (es *exerciseServiceImpl) GenerateReadingComprehensionExercise(user *storage.User, lessonId int) (exercise agents.GeneratedReadingComprehensionExercise, err error) {
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

	exercise, err = es.exerciseAgent.GenerateReadingComprehensionExercise(user.TargetLanguage, grammar.Pattern, words)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	return
}

func (es *exerciseServiceImpl) GradeReadingComprehensionResponse(user *storage.User, lessonId int, shortText, question, answer string) (grades []agents.UsageGrade, err error) {
	grades = make([]agents.UsageGrade, 0)
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

	gradeGrammarUsage, err := es.exerciseAgent.GradeUsage(user.TargetLanguage, answer, grammar.Pattern)
	if err != nil {
		utils.Logger.Error(err.Error())
	} else if gradeGrammarUsage.Score > 0 {
		grades = append(grades, gradeGrammarUsage)
		grammar.Score += gradeGrammarUsage.Score
		err = es.userGrammarRepository.Upsert(grammar)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
	}

	for _, id := range lesson.WordsId {
		word, err := es.userWordRepository.GetByID(id)
		if err != nil {
			utils.Logger.Error(err.Error())
			continue
		}
		gradeWordUsage, err := es.exerciseAgent.GradeUsage(user.TargetLanguage, answer, word.Word)
		if err != nil {
			utils.Logger.Error(err.Error())
		} else if gradeWordUsage.Score > 0 {
			grades = append(grades, gradeWordUsage)
			word.Score += gradeWordUsage.Score
			err = es.userWordRepository.Upsert(word)
			if err != nil {
				utils.Logger.Error(err.Error())
			}
		}
	}

	responseGrade, err := es.exerciseAgent.GradeReadingComprehensionResponse(user.TargetLanguage, shortText, question, answer)
	if err != nil {
		utils.Logger.Error(err.Error())
	} else {
		grades = append(grades, responseGrade)
	}
	return
}

func (es *exerciseServiceImpl) GenerateDialogueInitExercise(user *storage.User, lessonId int) (exercise agents.GeneratedDialogueInitExercise, err error) {
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

	exercise, err = es.exerciseAgent.GenerateDialogueInitExercise(user.TargetLanguage, grammar.Pattern, words)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	return
}

func (es *exerciseServiceImpl) GenerateDialogueContinuationExercise(user *storage.User, lessonId int, history string) (exercise agents.GeneratedDialogueContinuationExercise, err error) {
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

	exercise, err = es.exerciseAgent.GenerateDialogueContinuationExercise(user.TargetLanguage, grammar.Pattern, words, history)
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	return
}
