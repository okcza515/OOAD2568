package menu

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

func BuildEvaluatePresentationMenu(
	scorePresentationAdvisorController *controller.ScoreAdvisorController[*model.ScorePresentationAdvisor],
	scorePresentationCommitteeController *controller.ScoreCommitteeController[*model.ScorePresentationCommittee],
) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Evaluate Presentation",
		Children: []*utils.MenuItem{
			{
				Title: "For Advisor",
				Action: func(io *utils.MenuIO) {
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
					if err != nil {
						return
					}

					if score < 0 || score > 100 {
						io.Println("Invalid Score. Must be between 0.0 and 100.0.")
						return
					}

					newScore := &model.ScorePresentationAdvisor{
						PresentationId: uint(presentationId),
						AdvisorId:      uint(advisorId),
						Score:          score,
					}
					if err := scorePresentationAdvisorController.Insert(newScore); err != nil {
						io.Println(fmt.Sprintf("Failed to insert advisor score: %v", err))
					} else {
						io.Println("Advisor score submitted successfully!")
					}
				},
			},
			{
				Title: "For Committee",
				Action: func(io *utils.MenuIO) {
					io.Println("Evaluating Presentation for Committee...")

					io.Print("Enter Presentation ID (-1 to cancel): ")
					presentationIdStr, err := io.ReadInput()
					if err != nil || presentationIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					presentationId, err := strconv.Atoi(presentationIdStr)
					if err != nil {
						io.Println("Invalid Presentation ID.")
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
					score, err := io.ReadInputFloat()
					if err != nil {
						return
					}
					if score < 0 || score > 100 {
						io.Println("Invalid Score. Must be between 0.0 and 100.0.")
						return
					}

					if err := scorePresentationCommitteeController.Insert(&model.ScorePresentationCommittee{
						PresentationId: uint(presentationId),
						CommitteeId:    uint(committeeId),
						Score:          score,
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
					io.Println("Checking Scores for Presentation...")

					io.Print("Enter Presentation ID (-1 to cancel): ")
					presentationId, err := io.ReadInputID()
					if err != nil {
						io.Println("Cancelled.")
						return
					}

					// Fetch advisor scores
					advisorScores, err := scorePresentationAdvisorController.List(map[string]interface{}{
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

					// Fetch committee scores
					committeeScores, err := scorePresentationCommitteeController.List(map[string]interface{}{"presentation_id": presentationId})
					if err != nil {
						io.Println(fmt.Sprintf("Error fetching committee scores: %v", err))
					} else if len(committeeScores) == 0 {
						io.Println("No committee scores found for this presentation.")
					} else {
						io.Println("Committee Scores:")
						io.PrintTableFromSlice(advisorScores, []string{"CommitteeId", "Score"})
					}
				},
			},
		},
	}
}
