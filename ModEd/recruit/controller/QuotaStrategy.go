package controller

import (
	"ModEd/recruit/model"
	"errors"
	"fmt"
)

type QuotaStrategy struct{}

func (q *QuotaStrategy) ApplyForm(applicant *model.Applicant) error {
	fmt.Println("Enter Math Grade: ")
	fmt.Scan(&applicant.MathGrade)

	fmt.Println("Enter Science Grade: ")
	fmt.Scan(&applicant.ScienceGrade)

	fmt.Println("Enter English Grade: ")
	fmt.Scan(&applicant.EnglishGrade)

	return nil
}

// Implements FormRound
func (q *QuotaStrategy) GetForm() []string {
	return []string{"MathGrade", "ScienceGrade", "EnglishGrade"}
}

func (q *QuotaStrategy) Validate(data map[string]string) error {
	requiredFields := []string{"MathGrade", "ScienceGrade", "EnglishGrade"}
	for _, field := range requiredFields {
		if data[field] == "" {
			return errors.New("missing field: " + field)
		}
	}
	return nil
}
