package models

type LessonItem struct {
	Id              int
	UserId          int
	Language        string
	Grammar         string
	Words           []string
	Content         string
	SampleSentences string
	WordsMeaning    string
}
