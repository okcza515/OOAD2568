package controller

import (
	// "ModEd/eval/model"
	"gorm.io/gorm"
)

type IEvaluationFacade interface {

}

type EvaluationFacade struct {
	answerCtrl IAnswerController
	resultCtrl IResultController
}

func NewEvaluationFacade(db *gorm.DB) *EvaluationFacade {
	answerController := NewAnswerController(db)
	resultController := NewResultController(db)
	return &EvaluationFacade{answerCtrl: answerController, resultCtrl: resultController}
}