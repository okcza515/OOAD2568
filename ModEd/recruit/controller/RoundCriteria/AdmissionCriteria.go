// MEP-1003 Student Recruitment
package roundCriteria

import (
	"ModEd/recruit/model"
	"fmt"
)

const (
	MinimumGPAX = 3.0
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
		return true
	}

	if applicant.GPAX < MinimumGPAX {
		return false
	}

	return true
}
