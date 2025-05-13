package scoreManager

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

type SubmitAdvisorScoreHandler struct {
	scoreCtrl      *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor]
	assessmentCtrl *controller.AssessmentController
	linkCtrl       *controller.AssessmentCriteriaLinkController
}

func (h *SubmitAdvisorScoreHandler) GetInputSequence() []utils.InputPrompt {
	return []utils.InputPrompt{
		{
			Label: "Senior Project ID (number, -1 to cancel)",
			Validate: func(input string) error {
				if input == "-1" {
					return fmt.Errorf("operation cancelled")
				}
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
		{
			Label: "Criteria ID (number, -1 to cancel)",
			Validate: func(input string) error {
				if input == "-1" {
					return fmt.Errorf("operation cancelled")
				}
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
		{
			Label: "Advisor ID (number)",
			Validate: func(input string) error {
				if input == "-1" {
					return fmt.Errorf("operation cancelled")
				}
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
		{
			Label: "Score (0.0-100.0)",
			Validate: func(input string) error {
				if input == "-1" {
					return fmt.Errorf("operation cancelled")
				}
				score, err := strconv.ParseFloat(input, 64)
				if err != nil {
					return err
				}
				if score < 0 || score > 100 {
					return fmt.Errorf("score must be between 0 and 100")
				}
				return nil
			},
		},
	}
}

func (h *SubmitAdvisorScoreHandler) ExecuteMenuComponent(inputs []string) {
	projectID, _ := strconv.ParseUint(inputs[0], 10, 32)
	criteriaID, _ := strconv.ParseUint(inputs[1], 10, 32)
	advisorID, _ := strconv.ParseUint(inputs[2], 10, 32)
	scoreVal, _ := strconv.ParseFloat(inputs[3], 64)

	assessment, err := h.assessmentCtrl.RetrieveAssessmentBySeniorProjectId(uint(projectID))
	if err != nil {
		fmt.Println("Project not found")
		return
	}

	link, err := h.linkCtrl.RetrieveAssessmentCriteriaLink(assessment.ID, uint(criteriaID))
	if err != nil {
		fmt.Println("Criteria not found for this project")
		return
	}

	score := model.ScoreAssessmentAdvisor{
		AssessmentCriteriaLinkId: link.ID,
		AdvisorId:                uint(advisorID),
		Score:                    scoreVal,
	}

	if err := h.scoreCtrl.Insert(&score); err != nil {
		fmt.Printf("Failed to submit score: %v\n", err)
	} else {
		fmt.Println("Advisor score submitted successfully!")
	}
}
