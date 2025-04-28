//Chanawat Limpanatewin 65070503445
//MEP-1006

package controller

import (
	"ModEd/eval/model"
	"time"
)

type EvaluationController struct {
	evaluations []*model.Evaluation
}

func NewEvaluationController(evals []*model.Evaluation) *EvaluationController {
	return &EvaluationController{
		evaluations: evals,
	}
}

func (ec *EvaluationController) EvaluateAssignment(studentCode, instructorCode string, assignmentID uint, score uint) {
	for _, e := range ec.evaluations {
		if e.StudentCode == studentCode && e.AssignmentID != nil && *e.AssignmentID == assignmentID {
			e.Score = score
			e.EvaluatedAt = time.Now()
			return
		}
	}
	ec.evaluations = append(ec.evaluations, &model.Evaluation{
		StudentCode:    studentCode,
		InstructorCode: instructorCode,
		AssignmentID:   &assignmentID,
		Score:          score,
		EvaluatedAt:    time.Now(),
	})
}

func (ec *EvaluationController) CommentAssignment(studentCode string, assignmentID uint, comment string) {
	for _, e := range ec.evaluations {
		if e.StudentCode == studentCode && e.AssignmentID != nil && *e.AssignmentID == assignmentID {
			e.Comment = comment
			return
		}
	}
}

func (ec *EvaluationController) EvaluateQuiz(studentCode, instructorCode string, quizID uint, score uint) {
	for _, e := range ec.evaluations {
		if e.StudentCode == studentCode && e.QuizID != nil && *e.QuizID == quizID {
			e.Score = score
			e.EvaluatedAt = time.Now()
			return
		}
	}
	ec.evaluations = append(ec.evaluations, &model.Evaluation{
		StudentCode:    studentCode,
		InstructorCode: instructorCode,
		QuizID:         &quizID,
		Score:          score,
		EvaluatedAt:    time.Now(),
	})
}

func (ec *EvaluationController) ListEvaluations() []*model.Evaluation {
	return ec.evaluations
}

// package controller

// import (
// 	"ModEd/eval/model"
// 	"time"

// 	"gorm.io/gorm"
// )

// type EvaluationController struct {
// 	DB *gorm.DB
// }

// func NewEvaluationController(db *gorm.DB) *EvaluationController {
// 	return &EvaluationController{DB: db}
// }

// func (ec *EvaluationController) EvaluateAssignment(studentCode, instructorCode string, assignmentId uint, score uint, comment string) error {
// 	eval := model.Evaluation{
// 		StudentCode:    studentCode,
// 		InstructorCode: instructorCode,
// 		AssignmentId:   &assignmentId,
// 		Score:          score,
// 		Comment:        comment,
// 		EvaluatedAt:    time.Now(),
// 	}
// 	return ec.DB.Create(&eval).Error
// }

// func (ec *EvaluationController) EvaluateQuiz(studentCode, instructorCode string, quizId uint, score uint, comment string) error {
// 	eval := model.Evaluation{
// 		StudentCode:    studentCode,
// 		InstructorCode: instructorCode,
// 		QuizId:         &quizId,
// 		Score:          score,
// 		Comment:        comment,
// 		EvaluatedAt:    time.Now(),
// 	}
// 	return ec.DB.Create(&eval).Error
// }

// func (ec *EvaluationController) ListEvaluations(studentCode string) ([]model.Evaluation, error) {
// 	var results []model.Evaluation
// 	err := ec.DB.Where("student_code = ?", studentCode).Find(&results).Error
// 	return results, err
// }

//// package controller

// import (
// 	commonModel "ModEd/common/model"

// 	evalModel "ModEd/eval/model"

// 	"time"

// 	"gorm.io/gorm"
// )

// type Evaluation struct {
// 	gorm.Model
// 	Student    commonModel.Student    `gorm:"foreignKey:StudentCode;references:StudentCode"`
// 	Instructor commonModel.Instructor `gorm:"foreignKey:InstructorCode;references:InstructorCode"`
// 	Assignment evalModel.Assignment   `gorm:"foreignKey:AssignmentId;references:AssignmentId"`
// 	Quiz       evalModel.Quiz         `gorm:"foreignKey:QuizId;references:QuizId"`

// 	StudentCode    string `gorm:"not null"`
// 	InstructorCode string `gorm:"not null"`
// 	AssignmentId   *uint
// 	QuizId         *uint

// 	Score       float64 `gorm:"not null"`
// 	Comment     string
// 	EvaluatedAt time.Time `gorm:"not null"`
// }

// type EvaluationController struct {
// 	db *gorm.DB
// }

// func (ec *EvaluationController) SaveEvaluation(studentCode string, instructorCode string, assignmentID, quizID *uint, score float64, comment string) error {
// 	e := Evaluation{
// 		StudentCode:    studentCode,
// 		InstructorCode: instructorCode,
// 		AssignmentId:   assignmentID,
// 		QuizId:         quizID,
// 		Score:          score,
// 		Comment:        comment,
// 		EvaluatedAt:    time.Now(),
// 	}
// 	return ec.db.Create(&e).Error
// }

// func (ec *EvaluationController) ListEvaluations() ([]Evaluation, error) {
// 	var evals []Evaluation
// 	err := ec.db.
// 		Preload("Student").
// 		Preload("Instructor").
// 		Preload("Assignment").
// 		Preload("Quiz").
// 		Find(&evals).Error
// 	return evals, err
// }
