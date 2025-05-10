package assessmentManager

import (
	"ModEd/project/cli_refactor_prototype/criteriaManager"
	"ModEd/project/cli_refactor_prototype/scoreManager"
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
)

func BuildAssessmentManagerMenu(
	criteriaCtrl *controller.AssessmentCriteriaController,
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
	scoreAdvisorCtrl *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor],
	scoreCommitteeCtrl *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee],
) *utils.MenuHandler {
	menu := utils.NewMenuHandler("Assessment Manager")

	menu.AppendComponent("List Criteria Links",
		&ListCriteriaHandler{
			criteriaCtrl: criteriaCtrl,
			linkCtrl:     linkCtrl,
		})

	menu.AppendComponent("Link Criteria",
		&LinkCriteriaHandler{
			criteriaCtrl:   criteriaCtrl,
			assessmentCtrl: assessmentCtrl,
			linkCtrl:       linkCtrl,
		})

	menu.AppendComponent("Update Criteria Link",
		&UpdateLinkHandler{
			criteriaCtrl:   criteriaCtrl,
			assessmentCtrl: assessmentCtrl,
			linkCtrl:       linkCtrl,
		})

	menu.AppendComponent("Delete Criteria Link",
		&DeleteLinkHandler{
			assessmentCtrl: assessmentCtrl,
			linkCtrl:       linkCtrl,
			criteriaCtrl:   criteriaCtrl,
		})

	scoreMenu := scoreManager.BuildAssessmentScoreManagerMenu(
		assessmentCtrl,
		linkCtrl,
		criteriaCtrl,
		scoreAdvisorCtrl,
		scoreCommitteeCtrl,
	)
	scoreMenu.SetParentMenu(menu)
	menu.AppendComponent("Score Manager", scoreMenu)

	criteriaMenu := criteriaManager.BuildAssessmentCriteriaManagerMenu(criteriaCtrl)
	criteriaMenu.SetParentMenu(menu)
	menu.AppendComponent("Criteria Management", criteriaMenu)

	return menu
}
