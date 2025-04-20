// MEP-1003 Student Recruitment
package criteria

import "ModEd/recruit/model"

type EngineeringCriteria struct{}

func (c *EngineeringCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.TPAT1 >= 3 && applicant.GPAX >= 3.0 && applicant.TGAT1 >= 3
}
