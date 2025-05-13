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

func (f *InternshipEvaluationFacade) EvaluateInternship(informationID uint, criteriaScores map[uint]uint, comment string) error {
	criteriaList, err := f.CriteriaController.ListAllByInformationID(informationID)
	if err != nil {
		return fmt.Errorf("failed to retrieve criteria for InternshipInformation ID %d: %w", informationID, err)
	}

	if len(criteriaList) == 0 {
		return fmt.Errorf("no criteria found for InternshipInformation ID %d", informationID)
	}

	totalScore := uint(0)
	for criteriaID, score := range criteriaScores {
		if score < 1 || score > 5 {
			return fmt.Errorf("invalid score %d for criteria ID %d", score, criteriaID)
		}

		var targetCriteria *model.InternshipCriteria
		for _, criteria := range criteriaList {
			if criteria.ID == criteriaID {
				targetCriteria = &criteria
				break
			}
		}

		if targetCriteria == nil {
			return fmt.Errorf("criteria ID %d not found for InternshipInformation ID %d", criteriaID, informationID)
		}

		targetCriteria.Score = score
		if err := f.CriteriaController.Update(targetCriteria); err != nil {
			return fmt.Errorf("failed to update criteria ID %d: %w", criteriaID, err)
		}

		totalScore += score
	}

	result := &model.InternshipResultEvaluation{
		Comment:                 comment,
		Score:                   totalScore,
		InternshipInformationId: informationID,
	}

	if err := f.ResultController.Create(result); err != nil {
		return fmt.Errorf("failed to create result evaluation: %w", err)
	}

	return nil
}
