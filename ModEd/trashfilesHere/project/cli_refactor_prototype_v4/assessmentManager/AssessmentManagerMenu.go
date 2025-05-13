package assessmentManager

import (
	"ModEd/core/handler"
	"ModEd/project/cli_refactor_prototype_v4/assessmentScoreManager"
	"ModEd/project/cli_refactor_prototype_v4/criteriaManager"
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
)

func BuildAssessmentManagerMenu(
	criteriaCtrl *controller.AssessmentCriteriaController,
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
	scoreAdvisorCtrl *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor],
	scoreCommitteeCtrl *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee],
) *handler.HandlerContext {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Assessment Manager")

	base := NewBaseAssessmentStrategy(criteriaCtrl, assessmentCtrl, linkCtrl)

	ctx.AddHandler("1", "List Criteria Links", &ListCriteriaHandler{base})
	ctx.AddHandler("2", "Link Criteria", &LinkCriteriaHandler{base})
	ctx.AddHandler("3", "Update Criteria Link", &UpdateLinkHandler{base})
	ctx.AddHandler("4", "Delete Criteria Link", &DeleteLinkHandler{base})

	// Submenu handlers
	ctx.AddHandler("5", "Score Manager", createSubmenuHandler(
		assessmentScoreManager.BuildAssessmentScoreManagerMenu(
			assessmentCtrl,
			linkCtrl,
			criteriaCtrl,
			scoreAdvisorCtrl,
			scoreCommitteeCtrl,
		),
	))

	ctx.AddHandler("6", "Criteria Management", createSubmenuHandler(
		criteriaManager.BuildAssessmentCriteriaManagerMenu(criteriaCtrl),
	))

	ctx.AddBackHandler(handler.FuncStrategy{
		Action: func() error {
			fmt.Println("Exiting Assessment Manager...")
			return nil
		},
	})

	return ctx
}

func createSubmenuHandler(subCtx *handler.HandlerContext) handler.MenuStrategy {
	return handler.FuncStrategy{
		Action: func() error {
			for {
				subCtx.ShowMenu()
				var input string
				fmt.Print("Enter your choice: ")
				fmt.Scanln(&input)
				if input == "back" {
					return nil
				}
				if err := subCtx.HandleInput(input); err != nil {
					fmt.Println("Error:", err)
				}
			}
		},
	}
}
