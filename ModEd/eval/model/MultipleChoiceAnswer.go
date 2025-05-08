package model

import "ModEd/core"

type MultipleChoiceAnswer struct {
	core.BaseModel
	QuestionID  uint		`gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	Question  	Question 	`gorm:"foreignKey:QuestionID;references:ID" json:"question" csv:"question"`
	AnswerLabel string		`gorm:"type:text;not null" json:"answer_label" csv:"answer_label"`
	IsExpected  bool		`gorm:"type:boolean;not null" json:"is_expected" csv:"is_expected"`
}

// type MultipleChoiceQuestionFactory struct {
// 	Choice []string
// }

// func (f MultipleChoiceQuestionFactory) CreateQuestion(base Question) QuestionProductInterface {
// 	return &MultipleChoiceQuestion{
// 		Question: Question{
// 			ID:         base.ID,
// 			QuestionDetail: base.QuestionDetail,
// 			CorrectAnswer:  base.CorrectAnswer,
// 			Score:          base.Score,
// 		},
// 		Choice: f.Choice,
// 	}
// }
