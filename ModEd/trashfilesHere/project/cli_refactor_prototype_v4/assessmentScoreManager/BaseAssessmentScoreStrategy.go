package assessmentScoreManager

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
	"strconv"
)

type BaseAssessmentScoreStrategy struct {
	assessmentCtrl     *controller.AssessmentController
	linkCtrl           *controller.AssessmentCriteriaLinkController
	criteriaCtrl       *controller.AssessmentCriteriaController
	advisorScoreCtrl   *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor]
	committeeScoreCtrl *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee]
}

func NewBaseAssessmentScoreStrategy(
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
	criteriaCtrl *controller.AssessmentCriteriaController,
	advisorScoreCtrl *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor],
	committeeScoreCtrl *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee],
) *BaseAssessmentScoreStrategy {
	return &BaseAssessmentScoreStrategy{
		assessmentCtrl:     assessmentCtrl,
		linkCtrl:           linkCtrl,
		criteriaCtrl:       criteriaCtrl,
		advisorScoreCtrl:   advisorScoreCtrl,
		committeeScoreCtrl: committeeScoreCtrl,
	}
}

func (b *BaseAssessmentScoreStrategy) getProjectInput(prompt string) (uint, error) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return b.parseUintInput(input)
}

func (b *BaseAssessmentScoreStrategy) parseUintInput(input string) (uint, error) {
	if input == "-1" {
		return 0, fmt.Errorf("operation cancelled")
	}
	val, err := strconv.ParseUint(input, 10, 32)
	return uint(val), err
}

func (b *BaseAssessmentScoreStrategy) parseFloatInput(input string) (float64, error) {
	if input == "-1" {
		return 0, fmt.Errorf("operation cancelled")
	}
	return strconv.ParseFloat(input, 64)
}
