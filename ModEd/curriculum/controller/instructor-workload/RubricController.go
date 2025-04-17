package controller

import (
	model "ModEd/curriculum/model/instructor-workload"

	"gorm.io/gorm"
)

type RubricController struct {
	DB *gorm.DB
}

func (rc *RubricController) CreateRubric(rubric model.Rubric) error {
	return rc.DB.Create(&rubric).Error
}

func (rc *RubricController) GetRubricByAssignmentId(assignmentId int) (*model.Rubric, error) {
	var rubric model.Rubric
	err := rc.DB.Where("assignment_id = ?", assignmentId).First(&rubric).Error
	if err != nil {
		return nil, err
	}
	return &rubric, nil
}
