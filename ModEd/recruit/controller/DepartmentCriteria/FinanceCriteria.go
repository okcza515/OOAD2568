package departmentCriteria

import "ModEd/recruit/model"

type FinanceCriteria struct{}

const (
	minGPAX_Finance  = 2.5
	minTGAT1_Finance = 55.0
	minTGAT2_Finance = 45.0
	minTGAT3_Finance = 40.0
	minTPAT3_Finance = 40.0
)

func (m *FinanceCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Finance &&
		applicant.TGAT1 >= minTGAT1_Finance &&
		applicant.TGAT2 >= minTGAT2_Finance &&
		applicant.TGAT3 >= minTGAT3_Finance &&
		applicant.TPAT3 >= minTPAT3_Finance
}
