package main

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
)

func BuildEvaluateAssignmentMenu(
	scoreAssignmentAdvisorController *controller.ScoreAdvisorController[*model.ScoreAssignmentAdvisor],
	scoreAssignmentCommitteeController *controller.ScoreCommitteeController[*model.ScoreAssignmentCommittee],
) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Evaluate Assignment",
		Children: []*utils.MenuItem{
			{
				Title: "Insert Score For Advisor",
				Action: func(io *utils.MenuIO) {
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
					if err != nil {
						return
					}

					if score < 0 || score > 100 {
						io.Println("Invalid Score. Must be between 0.0 and 100.0.")
					}

					if err := scoreAssignmentAdvisorController.Insert(&model.ScoreAssignmentAdvisor{
						AssignmentId: uint(assignmentId),
						AdvisorId:    uint(advisorId),
						Score:        score,
					}); err != nil {
						io.Println(fmt.Sprintf("Failed to insert advisor score: %v", err))
					} else {
						io.Println("Advisor score submitted successfully!")
					}
				},
			},
			{
				Title: "Insert Score For Committee",
				Action: func(io *utils.MenuIO) {
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
					if err != nil {
						return
					}

					if err := scoreAssignmentCommitteeController.Insert(&model.ScoreAssignmentCommittee{
						AssignmentId: assignmentId,
						CommitteeId:  committeeId,
						Score:        score,
					}); err != nil {
						io.Println(fmt.Sprintf("Failed to insert committee score: %v", err))
					} else {
						io.Println("Committee score submitted successfully!")
					}
				},
			},
			{
				Title: "Check Score",
				Action: func(io *utils.MenuIO) {
					io.Println("Checking Scores for Assignment...")

					io.Print("Enter Assignment ID (-1 to cancel): ")
					assignmentId, err := io.ReadInputID()
					if err != nil {
						return
					}

					advisorScores, err := scoreAssignmentAdvisorController.List(
						map[string]interface{}{
							"assignment_id": assignmentId,
						})
					if err != nil {
						io.Println(fmt.Sprintf("Error fetching advisor scores: %v", err))
					} else if len(advisorScores) == 0 {
						io.Println("No advisor scores found for this assignment.")
					} else {
						io.Println("Advisor Scores:")
						for _, score := range advisorScores {
							io.Println(fmt.Sprintf("Advisor ID: %d, Score: %.2f", score.AdvisorId, score.Score))
						}
					}

					// Fetch committee scores
					committeeScores, err := scoreAssignmentCommitteeController.List(map[string]interface{}{"assignment_id": assignmentId})
					if err != nil {
						io.Println(fmt.Sprintf("Error fetching committee scores: %v", err))
					} else if len(committeeScores) == 0 {
						io.Println("No committee scores found for this assignment.")
					} else {
						io.Println("Committee Scores:")
						for _, score := range committeeScores {
							io.Println(fmt.Sprintf("Committee ID: %d, Score: %.2f", score.CommitteeId, score.Score))
						}
					}
				},
			},
		},
	}
}
