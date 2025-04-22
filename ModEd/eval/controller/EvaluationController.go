//Chanawat Limpanatewin 65070503445
//MEP-1006

package controller

import (
	commonModel "ModEd/common/model"

	evalModel "ModEd/eval/model"

	"time"

	"gorm.io/gorm"
)

type Evaluation struct {
	gorm.Model
	Student    commonModel.Student    `gorm:"foreignKey:StudentCode;references:StudentCode"`
	Instructor commonModel.Instructor `gorm:"foreignKey:InstructorCode;references:InstructorCode"`
	Assignment evalModel.Assignment   `gorm:"foreignKey:AssignmentId;references:AssignmentId"`
	Quiz       evalModel.Quiz         `gorm:"foreignKey:QuizId;references:QuizId"`

	StudentCode    string `gorm:"not null"`
	InstructorCode string `gorm:"not null"`
	AssignmentId   *uint
	QuizId         *uint

	Score       float64 `gorm:"not null"`
	Comment     string
	EvaluatedAt time.Time `gorm:"not null"`
}

type EvaluationController struct {
	db *gorm.DB
}

func (ec *EvaluationController) SaveEvaluation(studentCode string, instructorCode string, assignmentID, quizID *uint, score float64, comment string) error {
	e := Evaluation{
		StudentCode:    studentCode,
		InstructorCode: instructorCode,
		AssignmentId:   assignmentID,
		QuizId:         quizID,
		Score:          score,
		Comment:        comment,
		EvaluatedAt:    time.Now(),
	}
	return ec.db.Create(&e).Error
}

func (ec *EvaluationController) ListEvaluations() ([]Evaluation, error) {
	var evals []Evaluation
	err := ec.db.
		Preload("Student").
		Preload("Instructor").
		Preload("Assignment").
		Preload("Quiz").
		Find(&evals).Error
	return evals, err
}
