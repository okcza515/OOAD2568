package assessmentScoreManager

import (
	"ModEd/project/model"
	"fmt"
)

type SubmitCommitteeScoreHandler struct {
	*BaseAssessmentScoreStrategy
}

func (h *SubmitCommitteeScoreHandler) Execute() error {
	projectID, err := h.getProjectInput("Senior Project ID (number, -1 to cancel): ")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	criteriaID, err := h.getUintInput("Criteria ID (number, -1 to cancel): ")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	committeeID, err := h.getUintInput("Committee Member ID (number): ")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	scoreVal, err := h.getScoreInput()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	assessment, err := h.assessmentCtrl.RetrieveAssessmentBySeniorProjectId(projectID)
	if err != nil {
		fmt.Println("Project not found")
		return nil
	}

	link, err := h.linkCtrl.RetrieveAssessmentCriteriaLink(assessment.ID, criteriaID)
	if err != nil {
		fmt.Println("Criteria not found for this project")
		return nil
	}

	score := model.ScoreAssessmentCommittee{
		AssessmentCriteriaLinkId: link.ID,
		CommitteeId:              committeeID,
		Score:                    scoreVal,
	}

	if err := h.committeeScoreCtrl.Insert(&score); err != nil {
		return fmt.Errorf("failed to submit score: %w", err)
	}

	fmt.Println("Committee score submitted successfully!")
	return nil
}
