package controller

import (
	model "ModEd/curriculum/model"
	"fmt"
)

type InternshipEvaluationFacade struct {
	InfoController     InternshipInformationController
	CriteriaController InternshipCriteriaController
	ResultController   InternshipResultEvaluationController
}

func NewInternshipEvaluationFacade(
	infoCtrl InternshipInformationController,
	criteriaCtrl InternshipCriteriaController,
	resultCtrl InternshipResultEvaluationController,
) *InternshipEvaluationFacade {
	return &InternshipEvaluationFacade{
		InfoController:     infoCtrl,
		CriteriaController: criteriaCtrl,
		ResultController:   resultCtrl,
	}
}

func (f *InternshipEvaluationFacade) EvaluateInternship(studentCode string, criteriaScores map[uint]uint, comment string) error {

	allInfos, err := f.InfoController.ListAll()
	if err != nil {
		return err
	}

	var targetInfo *model.InternshipInformation
	for _, info := range allInfos {
		if info.StudentCode == studentCode {
			targetInfo = &info
			break
		}
	}

	if targetInfo == nil {
		return fmt.Errorf("student %s not found", studentCode)
	}

	totalScore := uint(0)
	for criteriaID, score := range criteriaScores {
		if score < 1 || score > 5 {
			return fmt.Errorf("invalid score %d for criteria ID %d", score, criteriaID)
		}

		criteria, err := f.CriteriaController.RetrieveByID(criteriaID)
		if err != nil {
			return err
		}

		criteria.Score = score
		if err := f.CriteriaController.Update(criteria); err != nil {
			return err
		}

		totalScore += score
	}

	result := &model.InternshipResultEvaluation{
		Comment:                 comment,
		Score:                   totalScore,
		InternshipInformationId: targetInfo.ID,
	}

	if err := f.ResultController.Create(result); err != nil {
		return err
	}

	return nil
}
