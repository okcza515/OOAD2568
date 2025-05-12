package departmentCriteria

import "ModEd/recruit/model"

type PhysicsCriteria struct{}

const (
	minGPAX_Physics  = 2.5
	minTGAT1_Physics = 40.0
	minTGAT2_Physics = 55.0
	minTGAT3_Physics = 40.0
	minTPAT3_Physics = 50.0
)

func (m *PhysicsCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Physics &&
		applicant.TGAT1 >= minTGAT1_Physics &&
		applicant.TGAT2 >= minTGAT2_Physics &&
		applicant.TGAT3 >= minTGAT3_Physics &&
		applicant.TPAT3 >= minTPAT3_Physics
}
