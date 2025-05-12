package departmentCriteria

import "ModEd/recruit/model"

type UrbanCriteria struct{}

const (
	minGPAX_Urban  = 2.5
	minTGAT1_Urban = 40.0
	minTGAT2_Urban = 40.0
	minTGAT3_Urban = 40.0
	minTPAT4_Urban = 55.0
)

func (m *UrbanCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Urban &&
		applicant.TGAT1 >= minTGAT1_Urban &&
		applicant.TGAT2 >= minTGAT2_Urban &&
		applicant.TGAT3 >= minTGAT3_Urban &&
		applicant.TPAT4 >= minTPAT4_Urban
}
