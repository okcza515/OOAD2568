package assessmentScoreManager

import (
	"ModEd/core/handler"
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
)

func BuildAssessmentScoreManagerMenu(
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
	criteriaCtrl *controller.AssessmentCriteriaController,
	advisorScoreCtrl *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor],
	committeeScoreCtrl *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee],
) *handler.HandlerContext {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Score Manager")

	base := NewBaseAssessmentScoreStrategy(
		assessmentCtrl,
		linkCtrl,
		criteriaCtrl,
		advisorScoreCtrl,
		committeeScoreCtrl,
	)

	ctx.AddHandler("1", "View Project Scores", &ViewScoresHandler{base})
	ctx.AddHandler("2", "Submit Advisor Score", &SubmitAdvisorScoreHandler{base})
	ctx.AddHandler("3", "Submit Committee Score", &SubmitCommitteeScoreHandler{base})

	ctx.AddBackHandler(handler.FuncStrategy{
		Action: func() error {
			fmt.Println("Returning to main menu...")
			return nil
		},
	})

	return ctx
}
