package assessmentManager

import (
	"ModEd/project/controller"
)

type BaseAssessmentStrategy struct {
	criteriaCtrl   *controller.AssessmentCriteriaController
	assessmentCtrl *controller.AssessmentController
	linkCtrl       *controller.AssessmentCriteriaLinkController
}

func NewBaseAssessmentStrategy(
	criteriaCtrl *controller.AssessmentCriteriaController,
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
) *BaseAssessmentStrategy {
	return &BaseAssessmentStrategy{
		criteriaCtrl:   criteriaCtrl,
		assessmentCtrl: assessmentCtrl,
		linkCtrl:       linkCtrl,
	}
}
