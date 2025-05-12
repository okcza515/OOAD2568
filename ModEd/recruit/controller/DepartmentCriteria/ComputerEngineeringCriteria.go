// MEP-1003 Student Recruitment
package departmentCriteria

import "ModEd/recruit/model"

type ComputerEngineeringCriteria struct{}

const (
	minGPAX_Computer  = 3.0
	minTGAT1_Computer = 60.0
	minTGAT2_Computer = 60.0
	minTGAT3_Computer = 60.0
	minTPAT3_Computer = 70.0
)

func (c *ComputerEngineeringCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Computer &&
		applicant.TGAT1 >= minTGAT1_Computer &&
		applicant.TGAT2 >= minTGAT2_Computer &&
		applicant.TGAT3 >= minTGAT3_Computer &&
		applicant.TPAT3 >= minTPAT3_Computer
}
