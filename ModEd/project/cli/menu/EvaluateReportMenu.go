package menu

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
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
				Action: func(io *core.MenuIO) {
					io.Println("Evaluating Report for Advisor...")

					io.Print("Enter Report ID (-1 to cancel): ")
					reportId, err := io.ReadInputID()
					if err != nil {
						io.Println("Cancelled.")
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
						return
					}

					if err := scoreReportAdvisorController.Insert(&model.ScoreReportAdvisor{
						ReportId:  reportId,
						AdvisorId: advisorId,
						Score:     score,
					}); err != nil {
						io.Println(fmt.Sprintf("Failed to insert advisor score: %v", err))
					} else {
						io.Println("Advisor score submitted successfully!")
					}
				},
			},
			{
				Title: "For Committee",
				Action: func(io *core.MenuIO) {
					io.Println("Evaluating Report for Committee...")

					io.Print("Enter Report ID (-1 to cancel): ")
					reportId, err := io.ReadInputID()
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
					if score < 0 || score > 100 {
						io.Println("Invalid Score. Must be between 0.0 and 100.0.")
						return
					}

					if err := scoreReportCommitteeController.Insert(&model.ScoreReportCommittee{
						ReportId:    uint(reportId),
						CommitteeId: uint(committeeId),
						Score:       score,
					}); err != nil {
						io.Println(fmt.Sprintf("Failed to insert committee score: %v", err))
					} else {
						io.Println("Committee score submitted successfully!")
					}
				},
			},
			{
				Title: "Check Score",
				Action: func(io *core.MenuIO) {
					io.Println("Checking Scores for Report...")

					io.Print("Enter Report ID (-1 to cancel): ")
					reportId, err := io.ReadInputID()
					if err != nil {
						return
					}

					advisorScores, err := scoreReportAdvisorController.List(map[string]interface{}{
						"report_id": reportId,
					})
					if err != nil {
						io.Println(fmt.Sprintf("Error fetching advisor scores: %v", err))
					} else if len(advisorScores) == 0 {
						io.Println("No advisor scores found for this report.")
					} else {
						io.Println("Advisor Scores:")
						io.PrintTableFromSlice(advisorScores, []string{"AdvisorId", "Score"})
					}

					committeeScores, err := scoreReportCommitteeController.List(map[string]interface{}{"report_id": reportId})
					if err != nil {
						io.Println(fmt.Sprintf("Error fetching committee scores: %v", err))
					} else if len(committeeScores) == 0 {
						io.Println("No committee scores found for this report.")
					} else {
						io.Println("Committee Scores:")
						io.PrintTableFromSlice(advisorScores, []string{"CommitteeId", "Score"})
					}
				},
			},
		},
	}
}
