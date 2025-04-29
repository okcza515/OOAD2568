package main

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
	"strconv"
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
	advisorScoreCtrl *controller.ScoreAdvisorController,
	committeeScoreCtrl *controller.ScoreCommitteeController,
) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
		io.Print("Enter Senior Project ID to view scores: ")
		input, _ := io.ReadInput()
		projectId, _ := strconv.Atoi(input)

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
			criteria, err := criteriaCtrl.RetrieveAssessmentCriteria(link.AssessmentCriteriaId)
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
	scoreCtrl *controller.ScoreAdvisorController,
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
		io.Print("Enter Senior Project ID: ")
		projectIDStr, _ := io.ReadInput()
		projectId, _ := strconv.Atoi(projectIDStr)

		assessment, err := assessmentCtrl.RetrieveAssessmentBySeniorProjectId(uint(projectId))
		if err != nil {
			io.Println("Project not found")
			return
		}

		io.Print("Enter Criteria ID to score: ")
		criteriaIDStr, _ := io.ReadInput()
		criteriaID, _ := strconv.Atoi(criteriaIDStr)

		link, err := linkCtrl.RetrieveAssessmentCriteriaLink(assessment.ID, uint(criteriaID))
		if err != nil {
			io.Println("Criteria not found for this project")
			return
		}

		io.Print("Enter Advisor ID: ")
		advisorIDStr, _ := io.ReadInput()
		advisorID, _ := strconv.ParseUint(advisorIDStr, 10, 64)

		io.Print("Enter score (0.0 - 100.0): ")
		scoreStr, _ := io.ReadInput()
		scoreVal, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil {
			io.Println("Invalid score.")
			return
		}

		score := model.ScoreAssessmentAdvisor{
			AssessmentCriteriaLinkId: link.ID,
			AdvisorId:                uint(advisorID),
			Score:                    scoreVal,
		}

		if err := scoreCtrl.InsertAdvisorScore(&score); err != nil {
			io.Println(fmt.Sprintf("Failed to submit score: %v", err))
		} else {
			io.Println("Advisor score submitted successfully!")
		}
	}
}

func submitCommitteeScore(
	scoreCtrl *controller.ScoreCommitteeController,
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
		io.Print("Enter Senior Project ID: ")
		projectIDStr, _ := io.ReadInput()
		projectId, _ := strconv.Atoi(projectIDStr)

		assessment, err := assessmentCtrl.RetrieveAssessmentBySeniorProjectId(uint(projectId))
		if err != nil {
			io.Println("Project not found")
			return
		}

		io.Print("Enter Criteria ID to score: ")
		criteriaIDStr, _ := io.ReadInput()
		criteriaID, _ := strconv.Atoi(criteriaIDStr)

		link, err := linkCtrl.RetrieveAssessmentCriteriaLink(assessment.ID, uint(criteriaID))
		if err != nil {
			io.Println("Criteria not found for this project")
			return
		}

		io.Print("Enter Committee Member ID: ")
		committeeIDStr, _ := io.ReadInput()
		committeeID, _ := strconv.ParseUint(committeeIDStr, 10, 64)

		io.Print("Enter score (0.0 - 100.0): ")
		scoreStr, _ := io.ReadInput()
		scoreVal, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil {
			io.Println("Invalid score.")
			return
		}

		score := model.ScoreAssessmentCommittee{
			AssessmentCriteriaLinkId: link.ID,
			CommitteeId:              uint(committeeID),
			Score:                    scoreVal,
		}

		if err := scoreCtrl.InsertCommitteeScore(&score); err != nil {
			io.Println(fmt.Sprintf("Failed to submit score: %v", err))
		} else {
			io.Println("Committee score submitted successfully!")
		}
	}
}

func displayAdvisorScore(io *utils.MenuIO, scoreCtrl *controller.ScoreAdvisorController, linkId uint) {
	advisorScore, err := scoreCtrl.RetrieveAdvisorScoreByCondition(
		"assessment", "assessment_criteria_link_id = ?", linkId,
	)
	if err == nil {
		if score, ok := advisorScore.(*model.ScoreAssessmentAdvisor); ok {
			io.Println(fmt.Sprintf("  Advisor Score: %.2f, By Advisor ID: %d", score.Score, score.AdvisorId))
		} else {
			io.Println("  Advisor Score: -")
		}
	} else {
		io.Println("  Advisor Score: -")
	}
}

func displayCommitteeScores(io *utils.MenuIO, scoreCtrl *controller.ScoreCommitteeController, linkId uint) {
	committeeScores, err := scoreCtrl.ListCommitteeScoresByCondition(
		"assessment", "assessment_criteria_link_id = ?", linkId,
	)
	if err != nil {
		io.Println("  Committee Scores: -")
		return
	}

	scoreList, ok := committeeScores.(*[]model.ScoreAssessmentCommittee)
	if !ok || len(*scoreList) == 0 {
		io.Println("  Committee Scores: -")
		return
	}

	io.Println("  Committee Scores:")
	for _, cs := range *scoreList {
		if cs.AssessmentCriteriaLinkId == linkId {
			io.Println(fmt.Sprintf("    - Score: %.2f, By Committee ID: %d", cs.Score, cs.CommitteeId))
		}
	}
}
