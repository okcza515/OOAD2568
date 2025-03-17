package controller

import (
	model "ModEd/curriculum/model/instructor-workload"

	"gorm.io/gorm"
)

type WorkloadReportController struct {
	Connector *gorm.DB
}

func CreateWorkloadReportController(connector *gorm.DB) *WorkloadReportController {
	return &WorkloadReportController{
		Connector: connector,
	}
}

func (repo WorkloadReportController) GetTeachingHour(instructorId string) (model.TeacherHourReport, error) {
	teachingHourReport := model.TeacherHourReport{}
	result := repo.Connector.First(&teachingHourReport, "instructor_id = ?", instructorId)
	return teachingHourReport, result.Error
}

func (repo WorkloadReportController) Get(teachingHourReport *model.TeacherHourReport) error {
	result := repo.Connector.Save(teachingHourReport)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
