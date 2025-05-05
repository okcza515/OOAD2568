package main

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

func BuildEvaluateReportMenu(
	scoreReportAdvisorController *controller.ScoreAdvisorController[*model.ScoreReportAdvisor],
	scoreReportCommitteeController *controller.ScoreCommitteeController[*model.ScoreReportCommittee],
) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Evaluate Report",
		Children: []*utils.MenuItem{
			{
				Title: "For Advisor",
				Action: func(io *utils.MenuIO) {
					io.Println("Evaluating Report for Advisor...")

					io.Print("Enter Report ID (-1 to cancel): ")
					reportIdStr, err := io.ReadInput()
					if err != nil || reportIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					reportId, err := strconv.Atoi(reportIdStr)
					if err != nil {
						io.Println("Invalid Report ID.")
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

					newScore := &model.ScoreReportAdvisor{
						ReportId:  uint(reportId),
						AdvisorId: uint(advisorId),
						Score:     score,
					}
					if err := scoreReportAdvisorController.Insert(newScore); err != nil {
						io.Println(fmt.Sprintf("Failed to insert advisor score: %v", err))
					} else {
						io.Println("Advisor score submitted successfully!")
					}
				},
			},
			{
				Title: "For Committee",
				Action: func(io *utils.MenuIO) {
					io.Println("Evaluating Report for Committee...")

					io.Print("Enter Report ID (-1 to cancel): ")
					reportIdStr, err := io.ReadInput()
					if err != nil || reportIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					reportId, err := strconv.Atoi(reportIdStr)
					if err != nil {
						io.Println("Invalid Report ID.")
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

					newScore := &model.ScoreReportCommittee{
						ReportId:    uint(reportId),
						CommitteeId: uint(committeeId),
						Score:       score,
					}
					if err := scoreReportCommitteeController.Insert(newScore); err != nil {
						io.Println(fmt.Sprintf("Failed to insert committee score: %v", err))
					} else {
						io.Println("Committee score submitted successfully!")
					}
				},
				},
			{
				Title: "Check Score",
				Action: func(io *utils.MenuIO) {
					io.Println("Checking Scores for Report...")

					io.Print("Enter Report ID (-1 to cancel): ")
					reportIdStr, err := io.ReadInput()
					if err != nil || reportIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					reportId, err := strconv.Atoi(reportIdStr)
					if err != nil {
						io.Println("Invalid Report ID.")
						return
					}

					// Fetch advisor scores
					advisorScores, err := scoreReportAdvisorController.ListAdvisorScoresByCondition("report_id", reportId)
					if err != nil {
						io.Println(fmt.Sprintf("Error fetching advisor scores: %v", err))
					} else if len(advisorScores) == 0 {
						io.Println("No advisor scores found for this report.")
					} else {
						io.Println("Advisor Scores:")
						for _, score := range advisorScores {
							io.Println(fmt.Sprintf("Advisor ID: %d, Score: %.2f", score.AdvisorId, score.Score))
						}
					}

					// Fetch committee scores
					committeeScores, err := scoreReportCommitteeController.ListCommitteeScoresByCondition("report_id", reportId)
					if err != nil {
						io.Println(fmt.Sprintf("Error fetching committee scores: %v", err))
					} else if len(committeeScores) == 0 {
						io.Println("No committee scores found for this report.")
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
