package menu

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
)

func BuildAssessmentScoreManagerMenu(
	assessmentController *controller.AssessmentController,
	assessmentCriteriaLinkController *controller.AssessmentCriteriaLinkController,
	assessmentCriteriaController *controller.AssessmentCriteriaController,
	scoreAdvisorController *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor],
	scoreCommitteeController *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee],
) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Assessment Score Manager",
		Children: []*utils.MenuItem{
			{
				Title: "View Project Scores",
				Action: viewProjectScores(
					assessmentController,
					assessmentCriteriaLinkController,
					assessmentCriteriaController,
					scoreAdvisorController,
					scoreCommitteeController,
				),
			},
			{
				Title: "Submit Advisor Score",
				Action: submitAdvisorScore(
					scoreAdvisorController,
					assessmentController,
					assessmentCriteriaLinkController,
				),
			},
			{
				Title: "Submit Committee Score",
				Action: submitCommitteeScore(
					scoreCommitteeController,
					assessmentController,
					assessmentCriteriaLinkController,
				),
			},
		},
	}
}

func viewProjectScores(
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
	criteriaCtrl *controller.AssessmentCriteriaController,
	advisorScoreCtrl *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor],
	committeeScoreCtrl *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee],
) func(*core.MenuIO) {
	return func(io *core.MenuIO) {
		io.Print("Enter Senior Project ID to view scores: ")
		projectId, _ := io.ReadInputID()

		_, err := assessmentCtrl.RetrieveAssessmentBySeniorProjectId(uint(projectId))
		if err != nil {
			io.Println(fmt.Sprintf("Error retrieving assessment: %v", err))
			return
		}

		links, err := linkCtrl.ListProjectAssessmentCriteriaLinks(uint(projectId))
		if err != nil || len(links) == 0 {
			io.Println("No assessment criteria linked to this project.")
			return
		}

		for _, link := range links {
			criteria, err := criteriaCtrl.RetrieveByID(link.AssessmentCriteriaId)
			if err != nil {
				continue
			}
			io.Println(fmt.Sprintf("\nCriteria ID: %d | Name: %s", criteria.ID, criteria.CriteriaName))

			displayAdvisorScore(io, advisorScoreCtrl, link.ID)

			displayCommitteeScores(io, committeeScoreCtrl, link.ID)
		}
	}
}

func submitAdvisorScore(
	scoreCtrl *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor],
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
) func(*core.MenuIO) {
	return func(io *core.MenuIO) {
		io.Print("Enter Senior Project ID: ")
		projectId, _ := io.ReadInputID()

		assessment, err := assessmentCtrl.RetrieveAssessmentBySeniorProjectId(projectId)
		if err != nil {
			io.Println("Project not found")
			return
		}

		io.Print("Enter Criteria ID to score: ")
		criteriaID, _ := io.ReadInputID()

		link, err := linkCtrl.RetrieveAssessmentCriteriaLink(assessment.ID, criteriaID)
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

		if err := scoreCtrl.Insert(&model.ScoreAssessmentAdvisor{
			AssessmentCriteriaLinkId: link.ID,
			AdvisorId:                advisorID,
			Score:                    scoreVal,
		}); err != nil {
			io.Println(fmt.Sprintf("Failed to submit score: %v", err))
		} else {
			io.Println("Advisor score submitted successfully!")
		}
	}
}

func submitCommitteeScore(
	scoreCtrl *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee],
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
) func(*core.MenuIO) {
	return func(io *core.MenuIO) {
		io.Print("Enter Senior Project ID: ")
		projectId, _ := io.ReadInputID()

		assessment, err := assessmentCtrl.RetrieveAssessmentBySeniorProjectId(projectId)
		if err != nil {
			io.Println("Project not found")
			return
		}

		io.Print("Enter Criteria ID to score: ")
		criteriaID, _ := io.ReadInputID()

		link, err := linkCtrl.RetrieveAssessmentCriteriaLink(assessment.ID, criteriaID)
		if err != nil {
			io.Println("Criteria not found for this project")
			return
		}

		io.Print("Enter Committee Member ID: ")
		committeeID, _ := io.ReadInputID()

		io.Print("Enter score (0.0 - 100.0): ")
		scoreVal, _ := io.ReadInputFloat()

		score := model.ScoreAssessmentCommittee{
			AssessmentCriteriaLinkId: link.ID,
			CommitteeId:              committeeID,
			Score:                    scoreVal,
		}

		if err := scoreCtrl.Insert(&score); err != nil {
			io.Println(fmt.Sprintf("Failed to submit score: %v", err))
		} else {
			io.Println("Committee score submitted successfully!")
		}
	}
}

func displayAdvisorScore(io *core.MenuIO, scoreCtrl *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor], linkId uint) {
	advisorScore, err := scoreCtrl.RetrieveByCondition(
		map[string]interface{}{"assessment_criteria_link_id": linkId},
	)
	if err != nil {
		io.Println("  Advisor Score: -")
		return
	}

	if advisorScore != nil {
		io.Println(fmt.Sprintf("  Advisor Score: %.2f, By Advisor ID: %d", advisorScore.Score, advisorScore.AdvisorId))
	} else {
		io.Println("  Advisor Score: -")
	}
}

func displayCommitteeScores(io *core.MenuIO, scoreCtrl *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee], linkId uint) {
	committeeScores, err := scoreCtrl.List(map[string]interface{}{"assessment_criteria_link_id": linkId})
	if err != nil {
		io.Println("  Committee Scores: -")
		return
	}

	if len(committeeScores) == 0 {
		io.Println("  Committee Scores: -")
		return
	}

	io.Println("  Committee Scores:")
	io.PrintTableFromSlice(committeeScores, []string{"Score", "CommitteeId"})
}
