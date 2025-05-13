// MEP-1003 Student Recruitment
package roundCriteria

import (
	"ModEd/recruit/model"
	"fmt"
)

type QuotaCriteria struct{}

const (
	MinGpax = 3.0
	Enum    = 3.5
)

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

	if applicant.GPAX < MinGpax {
		return false
	}

	if applicant.EnglishGrade < Enum || applicant.MathGrade < Enum || applicant.ScienceGrade < Enum {
		return false
	}

	return true
}
