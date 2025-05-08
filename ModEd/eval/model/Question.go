package model

import (
	"ModEd/core"
)

type Question struct {
	core.BaseModel
	SectionID		uint 			`gorm:"type:integer;not null" json:"section_id" csv:"section_id"`
	Section			ExamSection 	`gorm:"foreignKey:SectionID;references:ID" json:"section" csv:"section"`
	Score			float64			`gorm:"type:decimal(5,2);not null" json:"score" csv:"score"`
	ActualQuestion	string			`gorm:"type:text;not null" json:"actual_question" csv:"actual_question"`
	QuestionType	QuestionType	`gorm:"type:text;not null" json:"question_type" csv:"question_type"`
}

// type QuestionProductInterface interface {
// 	GetID() uint
// 	GetDetail() string
// 	GetCorrectAnswer() string
// 	GetScore() float64
// 	Validate() error
// }

// type QuestionFactory interface {
// 	CreateQuestion(base Question) QuestionProductInterface
// }

// type RegularQuestionFactory struct{}

// func (f RegularQuestionFactory) CreateQuestion(base Question) QuestionProductInterface {
// 	return &Question{
// 		ID:         	base.ID,
// 		QuestionDetail: base.QuestionDetail,
// 		CorrectAnswer:  base.CorrectAnswer,
// 		Score:          base.Score,
// 	}
// }

// func (q *Question) GetID() uint {
// 	return q.ID
// }

// func (q *Question) GetDetail() string {
// 	return q.QuestionDetail
// }

// func (q *Question) GetCorrectAnswer() string {
// 	return q.CorrectAnswer
// }

// func (q *Question) GetScore() float64 {
// 	return q.Score
// }

// func (q *Question) Validate() error {
// 	if q.QuestionDetail == "" {
// 		return fmt.Errorf("question detail cannot be empty")
// 	}
// 	if q.Score <= 0 {
// 		return fmt.Errorf("score must be greater than 0")
// 	}
// 	return nil
// }
