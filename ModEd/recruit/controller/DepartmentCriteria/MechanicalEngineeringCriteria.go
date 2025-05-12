package departmentCriteria

import "ModEd/recruit/model"

type MechanicalEngineeringCriteria struct{}

const (
	minGPAX_Mechanical  = 3.0
	minTPAT1_Mechanical = 70.0
	minTGAT1_Mechanical = 75.0
	minTGAT2_Mechanical = 80.0
)

func (m *MechanicalEngineeringCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.TPAT1 >= minTPAT1_Mechanical &&
		applicant.GPAX >= minGPAX_Mechanical &&
		applicant.TGAT1 >= minTGAT1_Mechanical &&
		applicant.TGAT2 >= minTGAT2_Mechanical
}
