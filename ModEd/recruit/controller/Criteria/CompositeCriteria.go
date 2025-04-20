// MEP-1003 Student Recruitment
package criteria

import "ModEd/recruit/model"

type CompositeCriteria struct {
	criteriaList []Criteria
}

func (c *CompositeCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	for _, criteria := range c.criteriaList {
		if !criteria.IsSatisfiedBy(applicant) {
			return false
		}
	}
	return true
}
