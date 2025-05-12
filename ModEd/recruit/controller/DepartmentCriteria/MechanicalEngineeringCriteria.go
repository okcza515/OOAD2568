package departmentCriteria

import "ModEd/recruit/model"

type MechanicalEngineeringCriteria struct{}

const (
	minGPAX_Mechanical  = 2.5
	minTGAT1_Mechanical = 50.0
	minTGAT2_Mechanical = 50.0
	minTGAT3_Mechanical = 50.0
	minTPAT3_Mechanical = 65.0
)

func (m *MechanicalEngineeringCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= minGPAX_Mechanical &&
		applicant.TGAT1 >= minTGAT1_Mechanical &&
		applicant.TGAT2 >= minTGAT2_Mechanical &&
		applicant.TGAT3 >= minTGAT3_Mechanical &&
		applicant.TPAT3 >= minTPAT3_Mechanical
}
