package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type AssessmentController struct {
	*core.BaseController[*model.Assessment]
	SeniorProjectController *core.BaseController[*model.SeniorProject]
	DB                      *gorm.DB
}

func NewAssessmentController(db *gorm.DB) *AssessmentController {
	return &AssessmentController{
		BaseController:          core.NewBaseController[*model.Assessment](db),
		SeniorProjectController: core.NewBaseController[*model.SeniorProject](db),
		DB:                      db,
	}
}

func (c *AssessmentController) RetrieveAssessmentBySeniorProjectId(seniorProjectId uint) (*model.Assessment, error) {
	// check if senior project exists
	_, err := c.SeniorProjectController.RetrieveByID(seniorProjectId)
	if err != nil {
		return nil, fmt.Errorf("the project does not exist: %w", err)
	}

	existingAssessments, err := c.List(map[string]interface{}{"senior_project_id": seniorProjectId})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the assessment: %w", err)
	}

	//if no assessment exists, auto create only once
	if len(existingAssessments) == 0 {
		assessment := model.Assessment{
			SeniorProjectId: seniorProjectId,
		}

		err = c.Insert(&assessment)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve the assessment: %w", err)
		}
	}
	return existingAssessments[0], nil
}

// func (c *AssessmentController) InsertAssessment(seniorProjectId uint) (*model.Assessment, error) {
// 	existingRecords, err := c.List(map[string]interface{}{"senior_project_id": seniorProjectId})
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to check existing assessments: %w", err)
// 	}

// 	if len(existingRecords) > 0 {
// 		return nil, fmt.Errorf("assessments already exist for project %d", seniorProjectId)
// 	}

// 	assessment := model.Assessment{
// 		SeniorProjectId: seniorProjectId,
// 	}

// 	return &assessment, c.Insert(&assessment)
// }
