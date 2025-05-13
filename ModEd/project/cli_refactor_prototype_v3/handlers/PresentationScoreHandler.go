package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
)

type PresentationScoreHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewPresentationScoreHandler(instance *controller.InstanceStorer) *PresentationScoreHandler {
	return &PresentationScoreHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instance,
	}
}

func (h *PresentationScoreHandler) InsertAdvisorScore(io *core.MenuIO) {
	io.Println("Evaluating Presentation for Advisor...")

	io.Print("Enter Presentation ID (-1 to cancel): ")
	presentationId, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter Advisor ID (-1 to cancel): ")
	advisorId, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter Score (0.0 - 100.0): ")
	score, err := io.ReadInputFloat()
	if err != nil || score < 0 || score > 100 {
		io.Println("Invalid Score. Must be between 0.0 and 100.0.")
		return
	}

	err = h.instanceStorer.ScorePresentationAdvisor.Insert(&model.ScorePresentationAdvisor{
		PresentationId: presentationId,
		AdvisorId:      advisorId,
		Score:          score,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Failed to insert advisor score: %v", err))
	} else {
		io.Println("Advisor score submitted successfully!")
	}
}

func (h *PresentationScoreHandler) InsertCommitteeScore(io *core.MenuIO) {
	io.Println("Evaluating Presentation for Committee...")

	io.Print("Enter Presentation ID (-1 to cancel): ")
	presentationId, err := io.ReadInputID()
	if err != nil {
		io.Println("Cancelled.")
		return
	}

	io.Print("Enter Committee ID (-1 to cancel): ")
	committeeId, err := io.ReadInputID()
	if err != nil {
		io.Println("Cancelled.")
		return
	}

	io.Print("Enter Score (0.0 - 100.0): ")
	score, err := io.ReadInputFloat()
	if err != nil || score < 0 || score > 100 {
		io.Println("Invalid Score. Must be between 0.0 and 100.0.")
		return
	}

	err = h.instanceStorer.ScorePresentationCommittee.Insert(&model.ScorePresentationCommittee{
		PresentationId: presentationId,
		CommitteeId:    committeeId,
		Score:          score,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Failed to insert committee score: %v", err))
	} else {
		io.Println("Committee score submitted successfully!")
	}
}

func (h *PresentationScoreHandler) CheckScore(io *core.MenuIO) {
	io.Println("Checking Scores for Presentation...")

	io.Print("Enter Presentation ID (-1 to cancel): ")
	presentationId, err := io.ReadInputID()
	if err != nil {
		io.Println("Cancelled.")
		return
	}

	advisorScores, err := h.instanceStorer.ScorePresentationAdvisor.List(map[string]interface{}{
		"presentation_id": presentationId,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Error fetching advisor scores: %v", err))
	} else if len(advisorScores) == 0 {
		io.Println("No advisor scores found for this presentation.")
	} else {
		io.Println("Advisor Scores:")
		io.PrintTableFromSlice(advisorScores, []string{"AdvisorId", "Score"})
	}

	committeeScores, err := h.instanceStorer.ScorePresentationCommittee.List(map[string]interface{}{
		"presentation_id": presentationId,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Error fetching committee scores: %v", err))
	} else if len(committeeScores) == 0 {
		io.Println("No committee scores found for this presentation.")
	} else {
		io.Println("Committee Scores:")
		io.PrintTableFromSlice(committeeScores, []string{"CommitteeId", "Score"})
	}
}
