// model/applicant.go
package model

type Applicant struct {
	ApplicantID        uint    `gorm:"primaryKey" json:"applicant_id" csv:"applicant_id"`
	FirstName          string  `json:"first_name" csv:"first_name"`
	LastName           string  `json:"last_name" csv:"last_name"`
	Email              string  `json:"email" csv:"email"`
	Address            string  `json:"address" csv:"address"`
	Phonenumber        string  `json:"phone_number" csv:"phone_number"`
	GPAX               float32 `json:"gpax" csv:"gpax"`
	HighSchool_Program string  `json:"high_school_program" csv:"high_school_program"`

	TGAT1 float32 `json:"tgat1" csv:"tgat1"`
	TGAT2 float32 `json:"tgat2" csv:"tgat2"`
	TGAT3 float32 `json:"tgat3" csv:"tgat3"`

	TPAT1 float32 `json:"tpat1" csv:"tpat1"`
	TPAT2 float32 `json:"tpat2" csv:"tpat2"`
	TPAT3 float32 `json:"tpat3" csv:"tpat3"`
	TPAT4 float32 `json:"tpat4" csv:"tpat4"`
	TPAT5 float32 `json:"tpat5" csv:"tpat5"`

	PortfolioURL string  `json:"portfolio_url" csv:"portfolio_url"`
	FamilyIncome float32 `json:"family_income" csv:"family_income"`
}
