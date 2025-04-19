package criteria

import "ModEd/recruit/model"

type ComputerEngineeringCriteria struct{}

func (c *ComputerEngineeringCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.TPAT1 >= 3 && applicant.GPAX >= 3.0 && applicant.TGAT1 >= 3 && applicant.TGAT2 >= 3
}
