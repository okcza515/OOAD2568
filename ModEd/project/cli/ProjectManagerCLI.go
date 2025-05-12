package main

import (
	"ModEd/core"
	"ModEd/project/cli/menu"
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
)

func main() {
	db := utils.OpenDatabase("project.db")
	db.Exec("PRAGMA foreign_keys = ON;")

	seniorProjectController := controller.NewSeniorProjectController(db)
	advisorController := controller.NewAdvisorController(db)
	committeeController := controller.NewCommitteeController(db)
	reportController := controller.NewReportController(db)
	progressController := controller.NewProgressController(db)
	assessmentController := controller.NewAssessmentController(db)
	assessmentCriteriaController := controller.NewAssessmentCriteriaController(db)
	assessmentCriteriaLinkController := controller.NewAssessmentCriteriaLinkController(db)
	scoreAssignmentAdvisorController := controller.NewScoreAdvisorController[*model.ScoreAssignmentAdvisor](db)
	scoreAssignmentCommitteeController := controller.NewScoreCommitteeController[*model.ScoreAssignmentCommittee](db)
	scoreReportAdvisorController := controller.NewScoreAdvisorController[*model.ScoreReportAdvisor](db)
	scoreReportCommitteeController := controller.NewScoreCommitteeController[*model.ScoreReportCommittee](db)
	scorePresentationAdvisorController := controller.NewScoreAdvisorController[*model.ScorePresentationAdvisor](db)
	scorePresentationCommitteeController := controller.NewScoreCommitteeController[*model.ScorePresentationCommittee](db)
	scoreAssessmentAdvisorController := controller.NewScoreAdvisorController[*model.ScoreAssessmentAdvisor](db)
	scoreAssessmentCommitteeController := controller.NewScoreCommitteeController[*model.ScoreAssessmentCommittee](db)

	utils.MenuTitle("Senior Project CLI")
	builder := utils.NewMenuBuilder(&utils.MenuItem{
		Title: "Main Menu",
		Children: []*utils.MenuItem{
			{
				Title: "Senior Project Setup",
				Children: []*utils.MenuItem{
					menu.BuildSeniorProjectMenu(seniorProjectController),
					{
						Title: "Assign Groups, Advisors, and Committees",
						Children: []*utils.MenuItem{
							menu.BuildAdvisorMenu(advisorController),
							menu.BuildCommitteeMenu(committeeController),
						},
					},
				},
			},
			{
				Title: "Project Execution and Monitoring",
				Children: []*utils.MenuItem{
					menu.BuildReportMenu(reportController),
					menu.BuildProgressMenu(progressController),
				},
			},
			{
				Title: "Evaluation & Assessment",
				Children: []*utils.MenuItem{
					menu.BuildEvaluateAssignmentMenu(
						scoreAssignmentAdvisorController,
						scoreAssignmentCommitteeController,
					),
					menu.BuildEvaluateReportMenu(
						scoreReportAdvisorController,
						scoreReportCommitteeController,
					),
					menu.BuildEvaluatePresentationMenu(
						scorePresentationAdvisorController,
						scorePresentationCommitteeController,
					),
					{
						Title: "Assessment Manager",
						Action: func(io *core.MenuIO) {
							menu.BuildAssessmentManagerMenu(
								assessmentCriteriaController,
								assessmentController,
								assessmentCriteriaLinkController,
								scoreAssessmentAdvisorController,
								scoreAssessmentCommitteeController,
							)
						},
					},
				},
			},
		},
	}, nil, nil)

	builder.AddMenuChild([]string{"Main Menu"}, &utils.MenuItem{
		Title:  "Example",
		Action: func(io *core.MenuIO) {},
	})

	builder.Show()
}
