package assessmentScoreManager

import (
	"fmt"
)

type ViewScoresHandler struct {
	*BaseAssessmentScoreStrategy
}

func (h *ViewScoresHandler) Execute() error {
	projectID, err := h.getProjectInput("Senior Project ID (number, -1 to cancel): ")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// Just check if project exists
	_, err = h.assessmentCtrl.RetrieveAssessmentBySeniorProjectId(projectID)
	if err != nil {
		fmt.Println("Project not found")
		return nil
	}

	links, err := h.linkCtrl.ListProjectAssessmentCriteriaLinks(projectID)
	if err != nil || len(links) == 0 {
		fmt.Println("No assessment criteria linked to this project.")
		return nil
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
	return nil
}

func (h *ViewScoresHandler) displayAdvisorScore(linkId uint) {
	score, err := h.advisorScoreCtrl.RetrieveByCondition(
		map[string]interface{}{"assessment_criteria_link_id": linkId},
	)
	if err != nil || score == nil {
		fmt.Println("  Advisor Score: -")
		return
	}
	fmt.Printf("  Advisor Score: %.2f, By Advisor ID: %d\n", score.Score, score.AdvisorId)
}

func (h *ViewScoresHandler) displayCommitteeScores(linkId uint) {
	scores, err := h.committeeScoreCtrl.List(
		map[string]interface{}{"assessment_criteria_link_id": linkId},
	)
	if err != nil || len(scores) == 0 {
		fmt.Println("  Committee Scores: -")
		return
	}

	fmt.Println("  Committee Scores:")
	for _, cs := range scores {
		fmt.Printf("    - Score: %.2f, By Committee ID: %d\n", cs.Score, cs.CommitteeId)
	}
}
