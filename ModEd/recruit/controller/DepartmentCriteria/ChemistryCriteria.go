package departmentCriteria

import "ModEd/recruit/model"

type ChemistryCriteria struct{}

const (
	minGPAX_Chemistry  = 2.5
	minTGAT1_Chemistry = 50.0
	minTGAT2_Chemistry = 50.0
	minTGAT3_Chemistry = 40.0
	minTPAT3_Chemistry = 45.0
)

func (m *ChemistryCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Chemistry &&
		applicant.TGAT1 >= minTGAT1_Chemistry &&
		applicant.TGAT2 >= minTGAT2_Chemistry &&
		applicant.TGAT3 >= minTGAT3_Chemistry &&
		applicant.TPAT3 >= minTPAT3_Chemistry
}
