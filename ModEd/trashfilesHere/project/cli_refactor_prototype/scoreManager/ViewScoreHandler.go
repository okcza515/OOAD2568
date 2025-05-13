package scoreManager

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

type ViewScoresHandler struct {
	assessmentCtrl     *controller.AssessmentController
	linkCtrl           *controller.AssessmentCriteriaLinkController
	criteriaCtrl       *controller.AssessmentCriteriaController
	advisorScoreCtrl   *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor]
	committeeScoreCtrl *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee]
}

func (h *ViewScoresHandler) GetInputSequence() []utils.InputPrompt {
	return []utils.InputPrompt{
		{
			Label: "Senior Project ID (number, -1 to cancel)",
			Validate: func(input string) error {
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
	}
}

func (h *ViewScoresHandler) ExecuteMenuComponent(inputs []string) {
	projectID, _ := strconv.ParseUint(inputs[0], 10, 32)

	_, err := h.assessmentCtrl.RetrieveAssessmentBySeniorProjectId(uint(projectID))
	if err != nil {
		fmt.Printf("Error retrieving assessment: %v\n", err)
		return
	}

	links, err := h.linkCtrl.ListProjectAssessmentCriteriaLinks(uint(projectID))
	if err != nil || len(links) == 0 {
		fmt.Println("No assessment criteria linked to this project.")
		return
	}

	for _, link := range links {
		criteria, err := h.criteriaCtrl.RetrieveByID(link.AssessmentCriteriaId)
		if err != nil {
			continue
		}
		fmt.Printf("\nCriteria ID: %d | Name: %s\n", criteria.ID, criteria.CriteriaName)

		h.displayAdvisorScore(link.ID)
		h.displayCommitteeScores(link.ID)
	}
}

func (h *ViewScoresHandler) displayAdvisorScore(linkId uint) {
	score, err := h.advisorScoreCtrl.RetrieveByCondition(
		map[string]interface{}{"assessment_criteria_link_id": linkId},
	)
	if err != nil {
		fmt.Println("  Advisor Score: -")
		return
	}

	if score != nil {
		fmt.Printf("  Advisor Score: %.2f, By Advisor ID: %d\n", score.Score, score.AdvisorId)
	} else {
		fmt.Println("  Advisor Score: -")
	}
}

func (h *ViewScoresHandler) displayCommitteeScores(linkId uint) {
	scores, err := h.committeeScoreCtrl.List(
		map[string]interface{}{
			"assessment_criteria_link_id": linkId,
		},
	)
	if err != nil {
		fmt.Println("  Committee Scores: -")
		return
	}

	if len(scores) == 0 {
		fmt.Println("  Committee Scores: -")
		return
	}

	fmt.Println("  Committee Scores:")
	for _, cs := range scores {
		fmt.Printf("    - Score: %.2f, By Committee ID: %d\n", cs.Score, cs.CommitteeId)
	}
}
