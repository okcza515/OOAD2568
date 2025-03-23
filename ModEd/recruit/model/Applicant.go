// MEP-1003 Student Recruitment
package model

// Applicant defines an applicant's information.
type Applicant struct {
	ApplicantID uint `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Email       string
	Address     string
	Phonenumber string
	GPAX        float32
	HighSchool_Program  string

	// TGAT Scores
	TGAT1 float32 // การคิดอย่างมีเหตุผล
	TGAT2 float32 // การสื่อสารภาษาอังกฤษ
	TGAT3 float32 // สมรรถนะการทำงาน

	// TPAT Scores (แบ่งตามประเภท)
	TPAT1  float32 // ความถนัดแพทย์
	TPAT2  float32 // ความถนัดสถาปัตย์
	TPAT3  float32 // ความถนัดวิศวะ
	TPAT4  float32 // ความถนัดครู
	TPAT5  float32 // ความถนัดศิลปกรรม
	Status string  // เพิ่มฟิลด์นี้
}
