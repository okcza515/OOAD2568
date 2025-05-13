package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
)

type AssessmentScoreHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewAssessmentScoreHandler(instanceStorer *controller.InstanceStorer) *AssessmentScoreHandler {
	return &AssessmentScoreHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

func (h *AssessmentScoreHandler) ViewProjectScores(io *core.MenuIO) {
	io.Print("Enter Senior Project ID to view scores: ")
	projectId, _ := io.ReadInputID()

	_, err := h.instanceStorer.Assessment.RetrieveAssessmentBySeniorProjectId(uint(projectId))
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving assessment: %v", err))
		return
	}

	links, err := h.instanceStorer.AssessmentCriteriaLink.ListProjectAssessmentCriteriaLinks(uint(projectId))
	if err != nil || len(links) == 0 {
		io.Println("No assessment criteria linked to this project.")
		return
	}

	for _, link := range links {
		criteria, err := h.instanceStorer.AssessmentCriteria.RetrieveByID(link.AssessmentCriteriaId)
		if err != nil {
			continue
		}
		io.Println(fmt.Sprintf("\nCriteria ID: %d | Name: %s", criteria.ID, criteria.CriteriaName))

		h.displayAdvisorScore(io, link.ID)
		h.displayCommitteeScores(io, link.ID)
	}
}

func (h *AssessmentScoreHandler) SubmitAdvisorScore(io *core.MenuIO) {
	io.Print("Enter Senior Project ID: ")
	projectId, _ := io.ReadInputID()

	assessment, err := h.instanceStorer.Assessment.RetrieveAssessmentBySeniorProjectId(projectId)
	if err != nil {
		io.Println("Project not found")
		return
	}

	io.Print("Enter Criteria ID to score: ")
	criteriaID, _ := io.ReadInputID()

	link, err := h.instanceStorer.AssessmentCriteriaLink.RetrieveAssessmentCriteriaLink(assessment.ID, criteriaID)
	if err != nil {
		io.Println("Criteria not found for this project")
		return
	}

	io.Print("Enter Advisor ID: ")
	advisorID, _ := io.ReadInputID()

	io.Print("Enter score (0.0 - 100.0): ")
	scoreVal, err := io.ReadInputFloat()
	if err != nil {
		io.Println("Invalid score.")
		return
	}

	err = h.instanceStorer.ScoreAssessmentAdvisor.Insert(&model.ScoreAssessmentAdvisor{
		AssessmentCriteriaLinkId: link.ID,
		AdvisorId:                advisorID,
		Score:                    scoreVal,
	})

	if err != nil {
		io.Println(fmt.Sprintf("Failed to submit score: %v", err))
	} else {
		io.Println("Advisor score submitted successfully!")
	}
}

func (h *AssessmentScoreHandler) SubmitCommitteeScore(io *core.MenuIO) {
	io.Print("Enter Senior Project ID: ")
	projectId, _ := io.ReadInputID()

	assessment, err := h.instanceStorer.Assessment.RetrieveAssessmentBySeniorProjectId(projectId)
	if err != nil {
		io.Println("Project not found")
		return
	}

	io.Print("Enter Criteria ID to score: ")
	criteriaID, _ := io.ReadInputID()

	link, err := h.instanceStorer.AssessmentCriteriaLink.RetrieveAssessmentCriteriaLink(assessment.ID, criteriaID)
	if err != nil {
		io.Println("Criteria not found for this project")
		return
	}

	io.Print("Enter Committee Member ID: ")
	committeeID, _ := io.ReadInputID()

	io.Print("Enter score (0.0 - 100.0): ")
	scoreVal, _ := io.ReadInputFloat()

	err = h.instanceStorer.ScoreAssessmentCommittee.Insert(&model.ScoreAssessmentCommittee{
		AssessmentCriteriaLinkId: link.ID,
		CommitteeId:              committeeID,
		Score:                    scoreVal,
	})

	if err != nil {
		io.Println(fmt.Sprintf("Failed to submit score: %v", err))
	} else {
		io.Println("Committee score submitted successfully!")
	}
}

func (h *AssessmentScoreHandler) displayAdvisorScore(io *core.MenuIO, linkId uint) {
	score, err := h.instanceStorer.ScoreAssessmentAdvisor.RetrieveByCondition(
		map[string]interface{}{"assessment_criteria_link_id": linkId},
	)
	if err != nil || score == nil {
		io.Println("  Advisor Score: -")
		return
	}

	io.Println(fmt.Sprintf("  Advisor Score: %.2f, By Advisor ID: %d", score.Score, score.AdvisorId))
}

func (h *AssessmentScoreHandler) displayCommitteeScores(io *core.MenuIO, linkId uint) {
	scores, err := h.instanceStorer.ScoreAssessmentCommittee.List(
		map[string]interface{}{"assessment_criteria_link_id": linkId},
	)
	if err != nil || len(scores) == 0 {
		io.Println("  Committee Scores: -")
		return
	}

	io.Println("  Committee Scores:")
	io.PrintTableFromSlice(scores, []string{"Score", "CommitteeId"})
}
