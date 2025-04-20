// MEP-1007
package controller

import (
	// commonmodel "ModEd/common/model"
	model "ModEd/eval/model"

	"gorm.io/gorm"
)

type IResultController interface {
	// CreateResults() error
	GetAllResults() ([]model.Result, error)
	UpdateResult(id uint, updatedResult *model.Result) error
	DeleteResult(id uint) error
}

type ResultController struct {
	db               *gorm.DB
	AnswerController *AnswerController
}

func NewResultController(db *gorm.DB) *ResultController {
	return &ResultController{db: db}
}

// func (c *ResultController) CreateResults() error {
// 	var exams []model.Examination
// 	var students []commonmodel.Student

// 	if err := c.db.Find(&exams).Error; err != nil {
// 		return err
// 	}

// 	if err := c.db.Find(&students).Error; err != nil {
// 		return err
// 	}

// 	for _, exam := range exams {
// 		status := GetExamStatusByID(exam.ID)
// 		questions := GetQuestionByExam(exam.ID)

// 		for _, student := range students {
// 			newResult := model.Result{
// 				ExaminationID: exam.ID,
// 				StudentID:     student.ID,
// 				Status:        status,
// 				Feedback:      "",
// 				Score:         0,
// 			}
// 			for _, question := range questions {
// 				answer, err := c.AnswerController.GetAnswerByQuestionAndStudent(question.ID, student.ID)
// 				if err != nil {
// 					return err
// 				}
				
// 				if answer.Question.Correct_answer == answer.Answer {
// 					newResult.Score += question.Score
// 				}
// 			}

// 			if err := c.db.Create(&newResult).Error; err != nil {
// 				return err
// 			}
// 			return nil
// 		}
// 	}
// }

func (c *ResultController) GetAllResults() ([]model.Result, error) {
	var results []model.Result
	if err := c.db.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (c *ResultController) UpdateResult(id uint, updatedData map[string]interface{}) error {
	var result model.Result
	if err := c.db.First(&result, id).Error; err != nil {
		return err
	}

	if err := c.db.Model(&result).Updates(updatedData).Error; err != nil {
		return err
	}
	return nil
}

func (c *ResultController) DeleteResult(id uint) error {
	var result model.Result
	if err := c.db.First(&result, id).Error; err != nil {
		return err
	}

	if err := c.db.Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}
