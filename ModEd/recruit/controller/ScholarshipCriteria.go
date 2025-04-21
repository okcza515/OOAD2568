// MEP-1003 Student Recruitment
package controller

import "ModEd/recruit/model"

type ScholarshipCriteria struct{}

func (c *ScholarshipCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= 3.5
}
