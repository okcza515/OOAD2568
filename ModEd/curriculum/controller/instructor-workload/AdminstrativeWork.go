package controller

import (
	model "ModEd/curriculum/model/instructor-workload"

	"gorm.io/gorm"
)

type AdminstrativeWorkController struct {
	Connector *gorm.DB
}

func CreateAdminstrativeWorkController(connector *gorm.DB) *AdminstrativeWorkController {
	return &AdminstrativeWorkController{
		Connector: connector,
	}
}

func (a AdminstrativeWorkController) GetAllStudentRequest(StudentAdvisor *model.StudentAdvisor) ([]model.StudentRequest, error) {
	studentRequests := []model.StudentRequest{}
	result := a.Connector.Find(&studentRequests, "advisor_id = ?", StudentAdvisor.ID)
	return studentRequests, result.Error
}

func (a AdminstrativeWorkController) AcceptStudentRequest(StudentRequest *model.StudentRequest) error {
	result := a.Connector.Save(StudentRequest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a AdminstrativeWorkController) RejectStudentRequest(StudentRequest *model.StudentRequest) error {
	// result := a.Connector.Delete(StudentRequest)
	// if result.Error != nil {
	// 	return result.Error
	// }
	// return nil
	return nil
}
