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
	criteriaList, err := f.CriteriaController.ListAllByStudentCode(studentCode)
	if err != nil {
		return fmt.Errorf("failed to retrieve criteria for student %s: %w", studentCode, err)
	}

	if len(criteriaList) == 0 {
		return fmt.Errorf("no criteria found for student %s", studentCode)
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
			return fmt.Errorf("criteria ID %d not found for student %s", criteriaID, studentCode)
		}

		targetCriteria.Score = score
		if err := f.CriteriaController.Update(targetCriteria); err != nil {
			return fmt.Errorf("failed to update criteria ID %d: %w", criteriaID, err)
		}

		totalScore += score
	}

	internshipInfo, err := f.InfoController.GetByStudentCode(studentCode)
	if err != nil {
		return fmt.Errorf("failed to retrieve internship information for student %s: %w", studentCode, err)
	}

	result := &model.InternshipResultEvaluation{
		Comment:                 comment,
		Score:                   totalScore,
		InternshipInformationId: internshipInfo.ID,
	}

	if err := f.ResultController.Create(result); err != nil {
		return fmt.Errorf("failed to create result evaluation: %w", err)
	}

	return nil
}
