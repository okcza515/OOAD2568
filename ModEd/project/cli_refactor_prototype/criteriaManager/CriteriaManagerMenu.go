package criteriaManager

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
)

func BuildAssessmentCriteriaManagerMenu(
	criteriaCtrl *controller.AssessmentCriteriaController,
) *utils.MenuHandler {
	menu := utils.NewMenuHandler("Criteria Management")

	menu.AppendComponent("Define Criteria",
		&DefineCriteriaHandler{controller: criteriaCtrl})

	menu.AppendComponent("List All Criteria",
		&ListCriteriaHandler{controller: criteriaCtrl})

	menu.AppendComponent("Update Criteria",
		&UpdateCriteriaHandler{controller: criteriaCtrl})

	menu.AppendComponent("Delete Criteria",
		&DeleteCriteriaHandler{controller: criteriaCtrl})

	return menu
}
