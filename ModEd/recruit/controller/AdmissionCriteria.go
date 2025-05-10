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

	// Check if Quota Category is provided
	if admissionCategory == "" {
		return false
	}

	// Check minimum GPAX
	if applicant.GPAX < 3 {
		return false
	}

	// ✅ Check minimum English score
	if applicant.EnglishGrade < 50 {
		return false
	}

	// ✅ Check minimum Math score
	if applicant.MathGrade < 50 {
		return false
	}

	// ✅ Check minimum Science score
	if applicant.ScienceGrade < 50 {
		return false
	}

	return true
}
