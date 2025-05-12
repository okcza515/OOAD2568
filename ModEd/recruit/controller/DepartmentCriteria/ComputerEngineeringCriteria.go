// MEP-1003 Student Recruitment
package departmentCriteria

import "ModEd/recruit/model"

type ComputerEngineeringCriteria struct{}

const (
	minTPAT1_Computer = 90.0
	minGPAX_Computer  = 90.0
	minTGAT1_Computer = 90.0
	minTGAT2_Computer = 90.0
)

func (c *ComputerEngineeringCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.TPAT1 >= minTPAT1_Computer &&
		applicant.GPAX >= minGPAX_Computer &&
		applicant.TGAT1 >= minTGAT1_Computer &&
		applicant.TGAT2 >= minTGAT2_Computer
}
