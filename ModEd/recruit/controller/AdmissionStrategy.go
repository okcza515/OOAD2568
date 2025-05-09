// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	
)

type AdmissionStrategy struct{}

// ApplyForm does not require any additional input for admission
func (a *AdmissionStrategy) ApplyForm(applicant *model.Applicant) error {
	return nil
}

// GetForm returns an empty list since no extra form fields are needed
func (a *AdmissionStrategy) GetForm() []string {
	return []string{}
}

// Validate returns nil as there's nothing to validate in this strategy
func (a *AdmissionStrategy) Validate(data map[string]string) error {
	return nil
}