package main

import (
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

	utils.PrintTitle("Senior Project CLI")

	builder := utils.NewMenuBuilder(&utils.MenuItem{
		Title: "Main Menu",
		Children: []*utils.MenuItem{
			{
				Title: "Senior Project Setup",
				Children: []*utils.MenuItem{
					{
						Title: "Create Senior Project",
						Action: func(io *utils.MenuIO) {
							io.Print("Enter the group name (-1 to cancel): ")
							groupNameStr, err := io.ReadInput()
							if err != nil || groupNameStr == "-1" {
								io.Println("Cancelled.")
								return
							}

							if err := seniorProjectController.Insert(&model.SeniorProject{
								GroupName: groupNameStr,
							}); err != nil {
								io.Println(err.Error())
								return
							}
						},
					},
					{
						Title: "List Senior Projects",
						Action: func(io *utils.MenuIO) {
							records, err := seniorProjectController.List(map[string]interface{}{})
							if err != nil {
								io.Println(err.Error())
								return
							}

							io.PrintTableFromSlice(records, []string{"ID", "GroupName", "CreatedAt"})
						},
					},
					{
						Title: "Assign Groups, Advisors, and Committees",
						Children: []*utils.MenuItem{
							BuildAdvisorMenu(advisorController),
							BuildCommitteeMenu(committeeController),
						},
					},
				},
			},
			{
				Title: "Project Execution and Monitoring",
				Children: []*utils.MenuItem{
					BuildReportMenu(reportController),
					BuildProgressMenu(progressController),
				},
			},
			{
				Title: "Evaluation & Assessment",
				Children: []*utils.MenuItem{
					BuildEvaluateAssignmentMenu(
						scoreAssignmentAdvisorController,
						scoreAssignmentCommitteeController,
					),
					BuildEvaluateReportMenu(
						scoreReportAdvisorController,
						scoreReportCommitteeController,
					),
					BuildEvaluatePresentationMenu(
						scorePresentationAdvisorController,
						scorePresentationCommitteeController,
					),
					{
						Title: "Assessment Manager",
						Action: func(io *utils.MenuIO) {
							BuildAssessmentManagerMenu(
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

	builder.Show()
}
