// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
)

type AdmissionStrategy struct{}

func (a *AdmissionStrategy) ApplyForm(applicant *model.Applicant) error {
	return nil
}

func (a *AdmissionStrategy) GetForm() []string {
	return []string{}
}

func (a *AdmissionStrategy) Validate(data map[string]string) error {
	return nil
}