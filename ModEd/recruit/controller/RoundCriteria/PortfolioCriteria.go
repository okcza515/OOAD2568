// MEP-1003 Student Recruitment
package roundCriteria

import (
	"ModEd/recruit/model"
	"fmt"
)

type PortfolioCriteria struct{}

func (c *PortfolioCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	data, err := applicant.GetRoundInfo()
	if err != nil {
		fmt.Println("Error retrieving Round Information :", err)
		return false
	}
	for key, value := range data {
		if key == "Portfolio URL" && value != "" {
			return applicant.GPAX >= 3.0
		}
	}
	return false
}
