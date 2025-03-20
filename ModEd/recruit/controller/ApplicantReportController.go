// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"fmt"

	"gorm.io/gorm"
)

type ApplicationReportController struct {
	DB *gorm.DB
}

func CreateApplicationReportController(db *gorm.DB) *ApplicationReportController {
	return &ApplicationReportController{DB: db}
}

func (ctrl *ApplicationReportController) SaveApplicationReport(report *model.ApplicationReport) error {
	result := ctrl.DB.Create(report)
	return result.Error
}

func (ctrl *ApplicationReportController) GetApplicantStatus() ([]string, error) {
	var statuses []string

	// ดึงข้อมูลสถานะทั้งหมดจากฐานข้อมูล
	if err := ctrl.DB.Model(&model.ApplicationReport{}).Pluck("application_statuses", &statuses).Error; err != nil {
		return nil, err // ถ้ามีข้อผิดพลาดในฐานข้อมูล ให้คืนค่าผลลัพธ์เป็น error
	}
	fmt.Println(statuses)
	return statuses, nil
}
