package main

import (
	"ModEd/project/cli_refactor_prototype/assessmentManager"
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
)

func main() {
	db := utils.OpenDatabase("project.db")
	db.Exec("PRAGMA foreign_keys = ON;")

	criteriaCtrl := controller.NewAssessmentCriteriaController(db)
	assessmentCtrl := controller.NewAssessmentController(db)
	linkCtrl := controller.NewAssessmentCriteriaLinkController(db)
	scoreAssessnentAdvisorCtrl := controller.NewScoreAdvisorController[*model.ScoreAssessmentAdvisor](db)
	scoreAssessmentCommitteeCtrl := controller.NewScoreCommitteeController[*model.ScoreAssessmentCommittee](db)

	mainMenu := utils.NewMenuHandler("Main Menu")

	assessmentMenu := assessmentManager.BuildAssessmentManagerMenu(
		criteriaCtrl,
		assessmentCtrl,
		linkCtrl,
		scoreAssessnentAdvisorCtrl,
		scoreAssessmentCommitteeCtrl,
	)
	assessmentMenu.SetParentMenu(mainMenu)
	mainMenu.AppendComponent("Assessment Management", assessmentMenu)

	mainMenu.ExecuteMenuComponent(nil)
}
