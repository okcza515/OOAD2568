package menus

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
					if err := scoreReportAdvisorController.InsertAdvisorScore(newScore); err != nil {
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
					if err := scoreReportCommitteeController.InsertCommitteeScore(newScore); err != nil {
						io.Println(fmt.Sprintf("Failed to insert committee score: %v", err))
					} else {
						io.Println("Committee score submitted successfully!")
					}
				},
			},
		},
	}
}
