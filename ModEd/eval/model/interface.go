// ไฟล์นี้กำหนดอินเตอร์เฟสพื้นฐานสำหรับระบบการส่งงานและการประเมิน
// โดยใช้หลักการ SOLID เพื่อให้ระบบมีความยืดหยุ่นและขยายได้ง่าย

// 65070503466 Warapol Pratumta
// MEP-1006

package model

import (
	"time"
)

// EvalSubmittable คืออินเตอร์เฟสสำหรับงานที่สามารถส่งได้
// กำหนดวิธีพื้นฐานที่จำเป็นสำหรับงานที่สามารถส่งได้
type EvalSubmittable interface {
	GetID() uint             // คืนค่ารหัสของงาน
	GetTitle() string        // คืนค่าชื่องาน
	GetDescription() string  // คืนค่าคำอธิบายงาน
	GetStartDate() time.Time // คืนค่าวันที่เริ่มงาน
	GetDueDate() time.Time   // คืนค่าวันที่ต้องส่งงาน
	GetMaxScore() float64    // คืนค่าคะแนนเต็ม
	GetType() string         // คืนค่าประเภทของงาน (Assignment/Quiz/Exam)
	GetCourseID() string     // คืนค่ารหัสวิชา
}

// EvalSubmission คืออินเตอร์เฟสสำหรับการส่งงานของนักศึกษา
// กำหนดวิธีพื้นฐานที่จำเป็นสำหรับการส่งงาน
type EvalSubmission interface {
	GetID() uint               // คืนค่ารหัสการส่งงาน
	GetStudentID() string      // คืนค่ารหัสนักศึกษา
	GetFirstName() string      // คืนค่าชื่อนักศึกษา
	GetLastName() string       // คืนค่านามสกุลนักศึกษา
	GetSubmittedAt() time.Time // คืนค่าวันที่ส่งงาน
	GetContent() string        // คืนค่าเนื้อหางานที่ส่ง
	IsLate() bool              // ตรวจสอบว่าส่งงานช้าหรือไม่
	GetStatus() string         // คืนค่าสถานะการส่งงาน
	GetScore() *float64        // คืนค่าคะแนนที่ได้รับ
	GetFeedback() string       // คืนค่าข้อเสนอแนะ
}

// EvalEvaluator คืออินเตอร์เฟสสำหรับผู้ประเมินงาน
// กำหนดวิธีพื้นฐานที่จำเป็นสำหรับผู้ประเมิน
type EvalEvaluator interface {
	GetID() string           // คืนค่ารหัสผู้ประเมิน (instructor_id)
	GetFirstName() string    // คืนค่าชื่อผู้ประเมิน
	GetLastName() string     // คืนค่านามสกุลผู้ประเมิน
	GetEmail() string        // คืนค่าemail ผู้ประเมิน
	GetStartDate() time.Time // คืนค่าวันที่เริ่มงาน
	GetFaculty() string      // คืนค่าคณะ
	GetDepartment() string   // คืนค่าภาควิชา
	GetCourseID() string     // คืนค่ารหัสวิชาที่สอน
}

// EvalEvaluation คืออินเตอร์เฟสสำหรับการประเมินงาน
// กำหนดวิธีพื้นฐานที่จำเป็นสำหรับการประเมินงาน
type EvalEvaluation interface {
	GetID() uint               // คืนค่ารหัสการประเมิน
	GetSubmissionID() uint     // คืนค่ารหัสงานที่ถูกประเมิน
	GetEvaluatorID() string    // คืนค่ารหัสผู้ประเมิน
	GetStudentID() string      // คืนค่ารหัสนักศึกษา
	GetScore() float64         // คืนค่าคะแนนที่ได้รับ
	GetFeedback() string       // คืนค่าข้อเสนอแนะ
	GetEvaluatedAt() time.Time // คืนค่าวันที่ประเมิน
	GetStatus() string         // คืนค่าสถานะการประเมิน
	GetGrade() string          // คืนค่าเกรดที่ได้รับ
	GetCourseID() string       // คืนค่ารหัสวิชา
}
