package scoreManager

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

type SubmitCommitteeScoreHandler struct {
	scoreCtrl      *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee]
	assessmentCtrl *controller.AssessmentController
	linkCtrl       *controller.AssessmentCriteriaLinkController
}

func (h *SubmitCommitteeScoreHandler) GetInputSequence() []utils.InputPrompt {
	return []utils.InputPrompt{
		{
			Label: "Senior Project ID (number, -1 to cancel)",
			Validate: func(input string) error {
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
		{
			Label: "Criteria ID (number, -1 to cancel)",
			Validate: func(input string) error {
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
		{
			Label: "Committee Member ID (number, -1 to cancel)",
			Validate: func(input string) error {
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
		{
			Label: "Score (0.0-100.0, -1 to cancel)",
			Validate: func(input string) error {
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

func (h *SubmitCommitteeScoreHandler) ExecuteMenuComponent(inputs []string) {
	projectID, _ := strconv.ParseUint(inputs[0], 10, 32)
	criteriaID, _ := strconv.ParseUint(inputs[1], 10, 32)
	committeeID, _ := strconv.ParseUint(inputs[2], 10, 32)
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

	score := model.ScoreAssessmentCommittee{
		AssessmentCriteriaLinkId: link.ID,
		CommitteeId:              uint(committeeID),
		Score:                    scoreVal,
	}

	if err := h.scoreCtrl.Insert(&score); err != nil {
		fmt.Printf("Failed to submit score: %v\n", err)
	} else {
		fmt.Println("Committee score submitted successfully!")
	}
}
