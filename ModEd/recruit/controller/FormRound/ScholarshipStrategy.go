package form

import (
	"ModEd/recruit/model"
	"fmt"
)

type ScholarshipStrategy struct{}

func (s *ScholarshipStrategy) ApplyForm(applicant *model.Applicant) error {
	fmt.Print("Enter Family Yearly Income: ")
	var income float32
	fmt.Scan(&income)
	applicant.FamilyIncome = income
	return nil
}
