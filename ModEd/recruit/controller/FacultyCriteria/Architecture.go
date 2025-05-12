// MEP-1003 Student Recruitment
package facultyCriteria

import "ModEd/recruit/model"

type ArchitectureCriteria struct{}

const (
	minGPAX_Achitechture  = 2
	minTGAT1_Achitechture = 40.0
	minTGAT2_Achitechture = 40.0
	minTGAT3_Achitechture = 40.0
	minTPAT4_Achitechture = 40.0
)

func (c *ArchitectureCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Achitechture &&
		applicant.TGAT1 >= minTGAT1_Achitechture &&
		applicant.TGAT2 >= minTGAT2_Achitechture &&
		applicant.TGAT3 >= minTGAT3_Achitechture &&
		applicant.TPAT4 >= minTPAT4_Achitechture
}
