// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"fmt"

	"gorm.io/gorm"
)

type InterviewCriteriaCtrl struct {
	DB *gorm.DB
}

func NewInterviewCriteriaCtrl(db *gorm.DB) *InterviewCriteriaCtrl {
	return &InterviewCriteriaCtrl{
		DB: db,
	}
}

func (ctrl *InterviewCriteriaCtrl) ReadInterviewCriteriaFromCSV(filePath string) error {

	if err := ctrl.DB.Exec("DELETE FROM interview_criteria").Error; err != nil {
		fmt.Println("Error clearing table:", err)
		return err
	}

	criteria, err := util.InsertFromCSVOrJSON[model.InterviewCriteria](filePath, ctrl.DB)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted %d interview criteria from file.\n", len(criteria))
	return nil
}

func (ctrl *InterviewCriteriaCtrl) GetFullInterviewCriteria() ([]model.InterviewCriteria, error) {
	var criteria []model.InterviewCriteria
	err := ctrl.DB.Preload("ApplicationRound").
		Preload("Faculty").
		Preload("Department").
		Find(&criteria).Error
	if err != nil {
		return nil, fmt.Errorf("failed to load interview criteria with related data: %w", err)
	}
	return criteria, nil
}
