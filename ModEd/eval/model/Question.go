package model

import (
	"ModEd/core"
)

type Question struct {
	core.BaseModel
	ExamID         uint         `gorm:"not null" csv:"exam_id" json:"exam_id"`
	SectionNo      uint         `gorm:"not null" csv:"section_no" json:"section_no"`
	ActualQuestion string       `gorm:"not null" csv:"actual_question" json:"actual_question"`
	QuestionType   QuestionType `gorm:"not null;type:varchar(50);" csv:"question_type" json:"question_type"`
	Score          float64      `gorm:"not null" csv:"score" json:"score"`
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
