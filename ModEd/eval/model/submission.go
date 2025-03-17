// 65070503466 Warapol Pratumta
// MEP-1006

// submission.go เป็นไฟล์ที่จัดการการส่งงานของนักศึกษา โดยมีหน้าที่หลักดังนี้:
// 1. จัดเก็บข้อมูลการส่งงาน (Submission) เช่น เนื้อหา, วันที่ส่ง, สถานะ
// 2. เชื่อมโยงข้อมูลระหว่างนักศึกษา (Student) และอาจารย์ (Instructor)
// 3. จัดการการตรวจให้คะแนนและข้อเสนอแนะจากอาจารย์
// 4. ติดตามสถานะการส่งงาน (ร่าง, ส่งแล้ว, ตรวจแล้ว, ส่งช้า)
// 5. รองรับการทำงานกับฐานข้อมูลผ่าน GORM

package model

import (
	"time"

	"ModEd/common/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SubmissionStatus คือประเภทของสถานะการส่งงาน
type SubmissionStatus string

const ( //สถานะการส่งงาน
	SubmissionDraft     SubmissionStatus = "draft"     // ร่าง
	SubmissionSubmitted SubmissionStatus = "submitted" // ส่งแล้ว
	SubmissionEvaluated SubmissionStatus = "evaluated" // ตรวจแล้ว
	SubmissionLate      SubmissionStatus = "late"      // ส่งช้า
)

// Submission คือโครงสร้างพื้นฐานสำหรับการส่งงาน
type Submission struct {
	gorm.Model
	SubmissionID   uuid.UUID        `gorm:"type:text;primaryKey;not null;unique" json:"submission_id"` // รหัสการส่งงาน
	SID            model.Student    `gorm:"foreignKey:SID"`                                            // รหัสนักศึกษา
	SubmissionDate time.Time        `gorm:"not null" json:"submission_date"`                           // วันที่ส่งงาน
	Status         SubmissionStatus `gorm:"type:text;default:'draft'" json:"status"`                   // สถานะการส่งงาน
	Content        string           `json:"content"`                                                   // เนื้อหาหรืออ้างอิงถึงงานที่ส่ง
	Feedback       string           `json:"feedback,omitempty"`                                        // ข้อเสนอแนะจากอาจารย์
	Score          *float64         `json:"score,omitempty"`                                           // คะแนนที่ได้รับ
	EvaluatedAt    *time.Time       `json:"evaluated_at,omitempty"`                                    // เวลาที่ตรวจงาน
	EvaluatedBy    model.Instructor `gorm:"foreignKey:EvaluatedByID" json:"evaluated_by,omitempty"`    // อาจารย์ที่ตรวจงาน
	EvaluatedByID  string           `gorm:"not null" json:"evaluated_by_id,omitempty"`                 // รหัสอาจารย์ที่ตรวจงาน
	CourseID       string           `gorm:"not null" json:"course_id"`                                 // รหัสวิชา
}

// GetID คืนค่ารหัสการส่งงาน
func (s *Submission) GetID() uint {
	return s.ID // คืนค่า ID จาก gorm.Model
}

// GetStudentID คืนค่ารหัสนักศึกษา
func (s *Submission) GetStudentID() string {
	return s.SID.SID // คืนค่ารหัสนักศึกษาจาก model.Student
}

// GetFirstName คืนค่าชื่อนักศึกษา
func (s *Submission) GetFirstName() string {
	return s.SID.FirstName // คืนค่าชื่อนักศึกษาจาก model.Student
}

// GetLastName คืนค่านามสกุลนักศึกษา
func (s *Submission) GetLastName() string {
	return s.SID.LastName // คืนค่านามสกุลนักศึกษาจาก model.Student
}

// GetSubmittedAt คืนค่าวันที่ส่งงาน
func (s *Submission) GetSubmittedAt() time.Time {
	return s.SubmissionDate // คืนค่าวันที่ส่งงาน
}

// GetContent คืนค่าเนื้อหางานที่ส่ง
func (s *Submission) GetContent() string {
	return s.Content // คืนค่าเนื้อหางานที่ส่ง
}

// IsLate ตรวจสอบว่าส่งงานช้าหรือไม่
func (s *Submission) IsLate() bool {
	return s.Status == SubmissionLate // ตรวจสอบสถานะการส่งงานว่าช้าหรือไม่
}

// GetStatus คืนค่าสถานะการส่งงาน
func (s *Submission) GetStatus() string {
	return string(s.Status) // คืนค่าสถานะการส่งงานในรูปแบบ string
}

// GetScore คืนค่าคะแนนที่ได้รับ
func (s *Submission) GetScore() *float64 {
	return s.Score // คืนค่าคะแนนที่ได้รับ
}

// GetFeedback คืนค่าข้อเสนอแนะ
func (s *Submission) GetFeedback() string {
	return s.Feedback // คืนค่าข้อเสนอแนะจากอาจารย์
}

// GetEvaluator คืนค่าข้อมูลอาจารย์ที่ตรวจงาน
func (s *Submission) GetEvaluator() model.Instructor {
	return s.EvaluatedBy // คืนค่าข้อมูลอาจารย์ที่ตรวจงาน
}

// GetEvaluatorID คืนค่ารหัสอาจารย์ที่ตรวจงาน
func (s *Submission) GetEvaluatorID() string {
	return s.EvaluatedByID // คืนค่ารหัสอาจารย์ที่ตรวจงาน
}

// GetCourseID คืนค่ารหัสวิชา
func (s *Submission) GetCourseID() string {
	return s.CourseID // คืนค่ารหัสวิชา
}

// SetScore กำหนดคะแนนและอัปเดตสถานะการส่งงาน
func (s *Submission) SetScore(score float64, evaluatorID string) {
	s.Score = &score // กำหนดคะแนนที่ได้รับ
	now := time.Now()
	s.EvaluatedAt = &now           // กำหนดเวลาที่ตรวจงาน
	s.EvaluatedByID = evaluatorID  // กำหนดรหัสอาจารย์ที่ตรวจงาน
	s.Status = SubmissionEvaluated // อัปเดตสถานะเป็นตรวจแล้ว
}

// SetFeedback กำหนดข้อเสนอแนะจากอาจารย์
func (s *Submission) SetFeedback(feedback string) {
	s.Feedback = feedback // กำหนดข้อเสนอแนะจากอาจารย์
}

// Submit ทำเครื่องหมายว่าส่งงานแล้วและกำหนดวันที่ส่ง
func (s *Submission) Submit(checkDueDate func(time.Time) bool) {
	s.SubmissionDate = time.Now() // กำหนดวันที่ส่งงานเป็นปัจจุบัน

	// ถ้า checkDueDate คืนค่า true แสดงว่าส่งงานช้า
	if checkDueDate(s.SubmissionDate) {
		s.Status = SubmissionLate // กำหนดสถานะเป็นส่งช้า
	} else {
		s.Status = SubmissionSubmitted // กำหนดสถานะเป็นส่งแล้ว
	}
}

// NewSubmission สร้างการส่งงานใหม่พร้อมรหัส UUID
func NewSubmission(sid model.Student, courseID string) Submission {
	return Submission{
		SubmissionID: uuid.New(),      // สร้างรหัส UUID ใหม่
		SID:          sid,             // กำหนดข้อมูลนักศึกษา
		Status:       SubmissionDraft, // กำหนดสถานะเริ่มต้นเป็นร่าง
		CourseID:     courseID,        // กำหนดรหัสวิชา
	}
}
