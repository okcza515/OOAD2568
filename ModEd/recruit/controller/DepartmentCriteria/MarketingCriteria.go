package departmentCriteria

import "ModEd/recruit/model"

type MarketingCriteria struct{}

const (
	minGPAX_Marketing  = 3
	minTGAT1_Marketing = 50.0
	minTGAT2_Marketing = 40.0
	minTGAT3_Marketing = 45.0
	minTPAT3_Marketing = 50.0
)

func (m *MarketingCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Marketing &&
		applicant.TGAT1 >= minTGAT1_Marketing &&
		applicant.TGAT2 >= minTGAT2_Marketing &&
		applicant.TGAT3 >= minTGAT3_Marketing &&
		applicant.TPAT3 >= minTPAT3_Marketing
}
