// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"fmt"
)

type QuotaCriteria struct{}

func (c *QuotaCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	data, err := applicant.GetRoundInfo()
	if err != nil {
		fmt.Println("Error retrieving Round Information:", err)
		return false
	}

	quotaCategory := ""
	for key, value := range data {
		if key == "Quota Category" {
			quotaCategory = value
			break
		}
	}

	// Check if Quota Category is provided
	if quotaCategory == "" {
		return false
	}

	// Check minimum GPAX
	if applicant.GPAX < 2.5 {
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
