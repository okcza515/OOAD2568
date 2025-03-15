// Assignment
package model

import (
	"time"

	"ModEd/common/model"

	"gorm.io/gorm"
)

type Assignment struct {
	// Instructor_Name  model.Instructor
	gorm.Model
	Title       string                 //หัวข้อ
	Description string                 //คำอธิบาย
	StartDate   time.Time              //วันที่ assign
	DueDate     time.Time              //วัน deadline
	Status      string                 //สเตตัส assignment
	Submission  []AssignmentSubmission //รายละเอียดการส่ง
	SubmitNum   uint                   //จำนวนคนที่ส่ง
}

type AssignmentSubmission struct { //รายละเอียดการส่ง
	gorm.Model
	SID         model.Student //เลขนักศึกษา
	Content     string        //ส่งอะไร
	SubmittedAt string        //เวลาที่ส่ง
}
