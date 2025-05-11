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

	if quotaCategory == "" {
		return false
	}

	if applicant.GPAX < 3 {
		return false
	}

	
	if applicant.EnglishGrade < 3.5 || applicant.MathGrade < 3.5 || applicant.ScienceGrade < 3.5  {
		return false
	}

	return true
}
