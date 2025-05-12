package departmentCriteria

import "ModEd/recruit/model"

type InteriaCriteria struct{}

const (
	minGPAX_Interia  = 2.5
	minTGAT1_Interia = 50.0
	minTGAT2_Interia = 50.0
	minTGAT3_Interia = 50.0
	minTPAT4_Interia = 65.0
)

func (m *InteriaCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Interia &&
		applicant.TGAT1 >= minTGAT1_Interia &&
		applicant.TGAT2 >= minTGAT2_Interia &&
		applicant.TGAT3 >= minTGAT3_Interia &&
		applicant.TPAT4 >= minTPAT4_Interia
}
