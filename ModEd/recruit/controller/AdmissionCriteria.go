// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"fmt"
)

type AdmissionCriteria struct{}

func (c *AdmissionCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	data, err := applicant.GetRoundInfo()
	if err != nil {
		fmt.Println("Error retrieving Round Information:", err)
		return false
	}

	admissionCategory := ""
	for key, value := range data {
		if key == "Admission Category" {
			admissionCategory = value
			break
		}
	}

	if admissionCategory == "" {
		return false
	}

	if applicant.GPAX < 3 || applicant.EnglishGrade < 50 || applicant.MathGrade < 50 || applicant.ScienceGrade < 50 {
		return false
	}

	return true
}
