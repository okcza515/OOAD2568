package scoreManager

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
)

func BuildAssessmentScoreManagerMenu(
	assessmentCtrl *controller.AssessmentController,
	linkCtrl *controller.AssessmentCriteriaLinkController,
	criteriaCtrl *controller.AssessmentCriteriaController,
	advisorScoreCtrl *controller.ScoreAdvisorController[*model.ScoreAssessmentAdvisor],
	committeeScoreCtrl *controller.ScoreCommitteeController[*model.ScoreAssessmentCommittee],
) *utils.MenuHandler {
	menu := utils.NewMenuHandler("Score Manager")

	menu.AppendComponent("View Project Scores",
		&ViewScoresHandler{
			assessmentCtrl:     assessmentCtrl,
			linkCtrl:           linkCtrl,
			criteriaCtrl:       criteriaCtrl,
			advisorScoreCtrl:   advisorScoreCtrl,
			committeeScoreCtrl: committeeScoreCtrl,
		})

	menu.AppendComponent("Submit Advisor Score",
		&SubmitAdvisorScoreHandler{
			scoreCtrl:      advisorScoreCtrl,
			assessmentCtrl: assessmentCtrl,
			linkCtrl:       linkCtrl,
		})

	menu.AppendComponent("Submit Committee Score",
		&SubmitCommitteeScoreHandler{
			scoreCtrl:      committeeScoreCtrl,
			assessmentCtrl: assessmentCtrl,
			linkCtrl:       linkCtrl,
		})

	return menu
}
