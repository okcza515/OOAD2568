package controller

import (
	model "ModEd/curriculum/model/Internship"
	"fmt"

	"gorm.io/gorm"
)

type ReportController struct {
	DB *gorm.DB
}

func (rc *ReportController) GetReportByStudentAndReportID(studentID string) error {
	var student model.InternStudent
	if err := rc.DB.Where("student_id = ?", studentID).First(&student); err != nil {
		return err.Error
	}
	fmt.Println("Found Student:", student)
	return nil
}
