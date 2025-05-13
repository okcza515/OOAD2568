// MEP-1003 Student Recruitment
package roundCriteria

import (
	"ModEd/recruit/model"
	"fmt"
)

type ScholarshipCriteria struct{}

func (c *ScholarshipCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	data, err := applicant.GetRoundInfo()
	if err != nil {
		fmt.Println("Error retrieving Round Information :", err)
		return false
	}
	for key, value := range data {
		if key == "Family Yearly Income" && value != "" {
			return applicant.GPAX >= 3.5
		}
	}
	return false
}
