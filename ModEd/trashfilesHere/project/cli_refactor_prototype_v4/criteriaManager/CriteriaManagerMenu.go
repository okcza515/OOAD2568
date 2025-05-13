package criteriaManager

import (
	"ModEd/core/handler"
	"ModEd/project/controller"
	"fmt"
)

func BuildAssessmentCriteriaManagerMenu(
	criteriaCtrl *controller.AssessmentCriteriaController,
) *handler.HandlerContext {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Criteria Management")

	base := NewBaseCriteriaStrategy(criteriaCtrl)

	ctx.AddHandler("1", "Define Criteria", &DefineCriteriaHandler{base})
	ctx.AddHandler("2", "List All Criteria", &ListCriteriaHandler{base})
	ctx.AddHandler("3", "Update Criteria", &UpdateCriteriaHandler{base})
	ctx.AddHandler("4", "Delete Criteria", &DeleteCriteriaHandler{base})

	ctx.AddBackHandler(handler.FuncStrategy{
		Action: func() error {
			fmt.Println("Returning to previous menu...")
			return nil
		},
	})

	return ctx
}
