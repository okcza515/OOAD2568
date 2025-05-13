package assessmentScoreManager

import (
	"ModEd/project/model"
	"fmt"
)

type SubmitAdvisorScoreHandler struct {
	*BaseAssessmentScoreStrategy
}

func (h *SubmitAdvisorScoreHandler) Execute() error {
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

	advisorID, err := h.getUintInput("Advisor ID (number): ")
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

	score := model.ScoreAssessmentAdvisor{
		AssessmentCriteriaLinkId: link.ID,
		AdvisorId:                advisorID,
		Score:                    scoreVal,
	}

	if err := h.advisorScoreCtrl.Insert(&score); err != nil {
		return fmt.Errorf("failed to submit score: %w", err)
	}

	fmt.Println("Advisor score submitted successfully!")
	return nil
}

func (h *BaseAssessmentScoreStrategy) getUintInput(prompt string) (uint, error) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return h.parseUintInput(input)
}

func (h *BaseAssessmentScoreStrategy) getScoreInput() (float64, error) {
	fmt.Print("Score (0.0-100.0): ")
	var input string
	fmt.Scanln(&input)

	score, err := h.parseFloatInput(input)
	if err != nil {
		return 0, err
	}

	if score < 0 || score > 100 {
		return 0, fmt.Errorf("score must be between 0 and 100")
	}
	return score, nil
}
