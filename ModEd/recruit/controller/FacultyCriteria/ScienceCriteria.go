// MEP-1003 Student Recruitment
package facultyCriteria

import "ModEd/recruit/model"

type ScienceCriteria struct{}

const (
	minGPAX_Science  = 2
	minTGAT1_Science = 30.0
	minTGAT2_Science = 30.0
	minTGAT3_Science = 30.0
	minTPAT3_Science = 30.0
)

func (c *ScienceCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Science &&
		applicant.TGAT1 >= minTGAT1_Science &&
		applicant.TGAT2 >= minTGAT2_Science &&
		applicant.TGAT3 >= minTGAT3_Science &&
		applicant.TPAT3 >= minTPAT3_Science
}
