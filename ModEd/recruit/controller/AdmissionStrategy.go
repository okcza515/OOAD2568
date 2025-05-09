// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
)

type AdmissionStrategy struct{}

func (a *AdmissionStrategy) ApplyForm(applicant *model.Applicant) error {
	// No additional input required for admission
	return nil
}
