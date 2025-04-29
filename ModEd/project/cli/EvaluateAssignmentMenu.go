package main

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

func BuildEvaluateAssignmentMenu(
	scoreAssignmentAdvisorController *controller.ScoreAdvisorController[*model.ScoreAssignmentAdvisor],
	scoreAssignmentCommitteeController *controller.ScoreCommitteeController[*model.ScoreAssignmentCommittee],
) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Evaluate Assignment",
		Children: []*utils.MenuItem{
			{
				Title: "For Advisor",
				Action: func(io *utils.MenuIO) {
					io.Println("Evaluating Assignment for Advisor...")

					io.Print("Enter Assignment ID (-1 to cancel): ")
					assignmentIdStr, err := io.ReadInput()
					if err != nil || assignmentIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					assignmentId, err := strconv.Atoi(assignmentIdStr)
					if err != nil {
						io.Println("Invalid Assignment ID.")
						return
					}

					io.Print("Enter Advisor ID (-1 to cancel): ")
					advisorIdStr, err := io.ReadInput()
					if err != nil || advisorIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					advisorId, err := strconv.Atoi(advisorIdStr)
					if err != nil {
						io.Println("Invalid Advisor ID.")
						return
					}

					io.Print("Enter Score (0.0 - 100.0): ")
					scoreStr, err := io.ReadInput()
					if err != nil {
						io.Println("Cancelled.")
						return
					}
					score, err := strconv.ParseFloat(scoreStr, 64)
					if err != nil || score < 0 || score > 100 {
						io.Println("Invalid Score. Must be between 0.0 and 100.0.")
						return
					}

					newScore := &model.ScoreAssignmentAdvisor{
						AssignmentId: uint(assignmentId),
						AdvisorId:    uint(advisorId),
						Score:        score,
					}
					if err := scoreAssignmentAdvisorController.InsertAdvisorScore(newScore); err != nil {
						io.Println(fmt.Sprintf("Failed to insert advisor score: %v", err))
					} else {
						io.Println("Advisor score submitted successfully!")
					}
				},
			},
			{
				Title: "For Committee",
				Action: func(io *utils.MenuIO) {
					io.Println("Evaluating Assignment for Committee...")

					io.Print("Enter Assignment ID (-1 to cancel): ")
					assignmentIdStr, err := io.ReadInput()
					if err != nil || assignmentIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					assignmentId, err := strconv.Atoi(assignmentIdStr)
					if err != nil {
						io.Println("Invalid Assignment ID.")
						return
					}

					io.Print("Enter Committee ID (-1 to cancel): ")
					committeeIdStr, err := io.ReadInput()
					if err != nil || committeeIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					committeeId, err := strconv.Atoi(committeeIdStr)
					if err != nil {
						io.Println("Invalid Committee ID.")
						return
					}

					io.Print("Enter Score (0.0 - 100.0): ")
					scoreStr, err := io.ReadInput()
					if err != nil {
						io.Println("Cancelled.")
						return
					}
					score, err := strconv.ParseFloat(scoreStr, 64)
					if err != nil || score < 0 || score > 100 {
						io.Println("Invalid Score. Must be between 0.0 and 100.0.")
						return
					}

					newScore := &model.ScoreAssignmentCommittee{
						AssignmentId: uint(assignmentId),
						CommitteeId:  uint(committeeId),
						Score:        score,
					}
					if err := scoreAssignmentCommitteeController.InsertCommitteeScore(newScore); err != nil {
						io.Println(fmt.Sprintf("Failed to insert committee score: %v", err))
					} else {
						io.Println("Committee score submitted successfully!")
					}
				},
			},
		},
	}
}
