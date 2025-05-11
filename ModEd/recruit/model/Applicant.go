// MEP-1003 Student Recruitment
package model

import (
	"encoding/json"
	"time"
)

type Applicant struct {
	ApplicantID        uint      `gorm:"primaryKey" json:"applicant_id" csv:"applicant_id"`
	FirstName          string    `gorm:"not null" json:"first_name" csv:"first_name"`
	LastName           string    `gorm:"not null" json:"last_name" csv:"last_name"`
	Email              string    `gorm:"not null" json:"email" csv:"email"`
	BirthDate          time.Time `gorm:"not null" json:"birth_date" csv:"birth_date"`
	Address            string    `gorm:"not null" json:"address" csv:"address"`
	Phonenumber        string    `gorm:"not null" json:"phone_number" csv:"phonenumber"`
	GPAX               float32   `gorm:"not null" json:"gpax" csv:"gpax"`
	HighSchool_Program string    `gorm:"not null" json:"high_school_program" csv:"high_school_program"`

	TGAT1 float32 `gorm:"not null" json:"tgat1" csv:"tgat1"`
	TGAT2 float32 `gorm:"not null" json:"tgat2" csv:"tgat2"`
	TGAT3 float32 `gorm:"not null" json:"tgat3" csv:"tgat3"`

	TPAT1 float32 `gorm:"not null" json:"tpat1" csv:"tpat1"`
	TPAT2 float32 `gorm:"not null" json:"tpat2" csv:"tpat2"`
	TPAT3 float32 `gorm:"not null" json:"tpat3" csv:"tpat3"`
	TPAT4 float32 `gorm:"not null" json:"tpat4" csv:"tpat4"`
	TPAT5 float32 `gorm:"not null" json:"tpat5" csv:"tpat5"`

	ApplicantRoundInformation string
	PortfolioURL              string  `gorm:"not null" json:"portfolio_url" csv:"portfolio_url"`
	FamilyIncome              float32 `gorm:"not null" json:"family_income" csv:"family_income"`
	MathGrade                 float32 `gorm:"default:0" json:"math_grade" csv:"math_grade"`
	ScienceGrade              float32 `gorm:"default:0" json:"science_grade" csv:"science_grade"`
	EnglishGrade              float32 `gorm:"default:0" json:"english_grade" csv:"english_grade"`
}

func (i *Applicant) GetID() uint {
	return i.ApplicantID
}
func (i *Applicant) FromCSV(csvData string) error {
	return nil
}
func (i *Applicant) ToCSVRow() string {
	return ""
}
func (i *Applicant) FromJSON(jsonData string) error {
	return nil
}
func (i *Applicant) ToJSON() string {
	return ""
}
func (i *Applicant) Validate() error {
	return nil
}
func (i *Applicant) ToString() string {
	return ""
}

func (i *Applicant) SetRoundInfo(data map[string]string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	i.ApplicantRoundInformation = string(jsonData)
	return nil
}

func (i *Applicant) GetRoundInfo() (map[string]string, error) {
	var data map[string]string
	err := json.Unmarshal([]byte(i.ApplicantRoundInformation), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
