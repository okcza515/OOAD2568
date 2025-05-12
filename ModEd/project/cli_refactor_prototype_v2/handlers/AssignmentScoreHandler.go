package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
)

type AssignmentScoreHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewAssignmentScoreHandler(instance *controller.InstanceStorer) *AssignmentScoreHandler {
	return &AssignmentScoreHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instance,
	}
}

func (h *AssignmentScoreHandler) InsertAdvisorScore(io *core.MenuIO) {
	io.Println("Evaluating Assignment for Advisor...")

	io.Print("Enter Assignment ID (-1 to cancel): ")
	assignmentId, err := io.ReadInputID()
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
		io.Println("Invalid score. Must be between 0.0 and 100.0.")
		return
	}

	err = h.instanceStorer.ScoreAssignmentAdvisor.Insert(&model.ScoreAssignmentAdvisor{
		AssignmentId: assignmentId,
		AdvisorId:    advisorId,
		Score:        score,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Failed to insert advisor score: %v", err))
	} else {
		io.Println("Advisor score submitted successfully!")
	}
}

func (h *AssignmentScoreHandler) InsertCommitteeScore(io *core.MenuIO) {
	io.Println("Evaluating Assignment for Committee...")

	io.Print("Enter Assignment ID (-1 to cancel): ")
	assignmentId, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter Committee ID (-1 to cancel): ")
	committeeId, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter Score (0.0 - 100.0): ")
	score, err := io.ReadInputFloat()
	if err != nil || score < 0 || score > 100 {
		io.Println("Invalid score. Must be between 0.0 and 100.0.")
		return
	}

	err = h.instanceStorer.ScoreAssignmentCommittee.Insert(&model.ScoreAssignmentCommittee{
		AssignmentId: assignmentId,
		CommitteeId:  committeeId,
		Score:        score,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Failed to insert committee score: %v", err))
	} else {
		io.Println("Committee score submitted successfully!")
	}
}

func (h *AssignmentScoreHandler) CheckScore(io *core.MenuIO) {
	io.Println("Checking Scores for Assignment...")

	io.Print("Enter Assignment ID (-1 to cancel): ")
	assignmentId, err := io.ReadInputID()
	if err != nil {
		return
	}

	advisorScores, err := h.instanceStorer.ScoreAssignmentAdvisor.List(map[string]interface{}{
		"assignment_id": assignmentId,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Error fetching advisor scores: %v", err))
	} else if len(advisorScores) == 0 {
		io.Println("No advisor scores found for this assignment.")
	} else {
		io.Println("Advisor Scores:")
		io.PrintTableFromSlice(advisorScores, []string{"AdvisorId", "Score"})
	}

	committeeScores, err := h.instanceStorer.ScoreAssignmentCommittee.List(map[string]interface{}{
		"assignment_id": assignmentId,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Error fetching committee scores: %v", err))
	} else if len(committeeScores) == 0 {
		io.Println("No committee scores found for this assignment.")
	} else {
		io.Println("Committee Scores:")
		io.PrintTableFromSlice(committeeScores, []string{"CommitteeId", "Score"})
	}
}
