// MEP-1003 Student Recruitment
package facultyCriteria

import "ModEd/recruit/model"

type BusinessCriteria struct{}

const (
	minGPAX_Business  = 2
	minTGAT1_Business = 30.0
	minTGAT2_Business = 30.0
	minTGAT3_Business = 30.0
	minTPAT3_Business = 30.0
)

func (c *BusinessCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Business &&
		applicant.TGAT1 >= minTGAT1_Business &&
		applicant.TGAT2 >= minTGAT2_Business &&
		applicant.TGAT3 >= minTGAT3_Business &&
		applicant.TPAT3 >= minTPAT3_Business
}
