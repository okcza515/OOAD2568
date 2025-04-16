package model

// Applicant defines an applicant's information.
type Applicant struct {
	ApplicantID        uint    `gorm:"primaryKey" json:"applicant_id" csv:"applicant_id"`
	FirstName          string  `json:"first_name" csv:"first_name"`
	LastName           string  `json:"last_name" csv:"last_name"`
	Email              string  `json:"email" csv:"email"`
	Address            string  `json:"address" csv:"address"`
	Phonenumber        string  `json:"phone_number" csv:"phone_number"`
	GPAX               float32 `json:"gpax" csv:"gpax"`
	HighSchool_Program string  `json:"high_school_program" csv:"high_school_program"`

	// TGAT Scores
	TGAT1 float32 `json:"tgat1" csv:"tgat1"` // การคิดอย่างมีเหตุผล
	TGAT2 float32 `json:"tgat2" csv:"tgat2"` // การสื่อสารภาษาอังกฤษ
	TGAT3 float32 `json:"tgat3" csv:"tgat3"` // สมรรถนะการทำงาน

	// TPAT Scores (แบ่งตามประเภท)
	TPAT1 float32 `json:"tpat1" csv:"tpat1"` // ความถนัดแพทย์
	TPAT2 float32 `json:"tpat2" csv:"tpat2"` // ความถนัดสถาปัตย์
	TPAT3 float32 `json:"tpat3" csv:"tpat3"` // ความถนัดวิศวะ
	TPAT4 float32 `json:"tpat4" csv:"tpat4"` // ความถนัดครู
	TPAT5 float32 `json:"tpat5" csv:"tpat5"` // ความถนัดศิลปกรรม

	Status string `json:"status" csv:"status"`
}
