// MEP-1003 Student Recruitment
package facultyCriteria

import "ModEd/recruit/model"

type EngineeringCriteria struct{}

const (
	minGPAX_Engineer  = 2.5
	minTGAT1_Engineer = 40.0
	minTGAT2_Engineer = 40.0
	minTGAT3_Engineer = 40.0
	minTPAT3_Engineer = 50.0
)

func (c *EngineeringCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Engineer &&
		applicant.TGAT1 >= minTGAT1_Engineer &&
		applicant.TGAT2 >= minTGAT2_Engineer &&
		applicant.TGAT3 >= minTGAT3_Engineer &&
		applicant.TPAT3 >= minTPAT3_Engineer
}
