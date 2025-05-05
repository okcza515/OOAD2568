// MEP-1007
package controller

import (
	commonmodel "ModEd/common/model"
	model "ModEd/eval/model"

	"gorm.io/gorm"
)

type IResultController interface {
	CreateResults() error
	GetResultByStudent(studentID uint) ([]model.Result, error)
	UpdateResult(id uint, updatedData map[string]interface{}) error
	DeleteResult(id uint) error
}

type ResultController struct {
	db                 *gorm.DB
	QuestionController *QuestionController
	AnswerController   *AnswerController
	GradingFactory     *GradingStrategyFactory 
}

func NewResultController(db *gorm.DB) *ResultController {
	return &ResultController{db: db}
}

func (c *ResultController) CreateResultByExamID(examID uint) error {
	var students []commonmodel.Student
 	var count int

	if err := c.db.Find(&students).Error; err != nil {
		return err
	}

	questions, err := c.QuestionController.GetQuestionsByExamID(examID)
	if err != nil {
		return err
	}

	for _, student := range students {
		newResult := model.Result {
			ExaminationID: examID,
			StudentID:     student.ID,
			Status:        "Pending",
			Feedback:      "",
			Score:         0,
		}
		count = 0

		for _, question := range questions {
			strategy := c.GradingFactory.GetStrategy(question.Question_type)

			answer, err := c.AnswerController.GetAnswerByQuestionAndStudent(question.ID, student.ID)
			if err != nil {
				return err
			}

			newResult.Score += strategy.Grade(question, *answer)
			count++
		}

		if count == len(questions) {
			newResult.Status = "Success"
		}

		if err := c.db.Create(&newResult).Error; err != nil {
			return err
		}
	}
	return nil
}

func (c *ResultController) GetResultByExamAndStudent(examID uint, studentID uint) ([]model.Result, error) {
	var results []model.Result
	if err := c.db.Where("examination_id = ? AND student_id = ?", examID, studentID).Preload("Examination").Preload("Student").Find(&results).Error; err != nil {
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
