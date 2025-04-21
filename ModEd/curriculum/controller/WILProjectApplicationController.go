package controller

import (
	"ModEd/core"
	commonModel "ModEd/common/model"
	model "ModEd/curriculum/model"
	utils "ModEd/curriculum/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type WILProjectApplicationController struct {
	connector *gorm.DB
	*core.BaseController
}

type WILProjectApplicationControllerInterface interface {
	RegisterWILProjectsApplication(projects []core.RecordInterface)
	Insert(data core.RecordInterface) error
	UpdateByID(data core.RecordInterface) error
	RetrieveByID(id uint, preloads ...string) (*core.RecordInterface, error)
	DeleteByID(id uint) error
	ListPagination(condition map[string]interface{}, page int, pageSize int)
}

func CreateWILProjectApplicationController(connector *gorm.DB) *WILProjectApplicationController {
	wil := WILProjectApplicationController{connector: connector, BaseController: core.NewBaseController("WILProjectApplication", connector)}
	return &wil
}

/*
func (repo WILProjectApplicationController) RegisterWILProjectApplication(application *model.WILProjectApplication) {
	repo.connector.Create(application)
}

func (repo WILProjectApplicationController) RegisterWILProjectApplications(applications []*model.WILProjectApplication) {
	for _, application := range applications {
		repo.connector.Create(application)
	}
}

func (repo WILProjectApplicationController) GetAllWILProjectApplications() ([]*model.WILProjectApplication, error) {
	applications := []*model.WILProjectApplication{}
	result := repo.connector.Find(&applications)
	return applications, result.Error
}

func (repo WILProjectApplicationController) GetWILProjectApplicationByID(id uint) (*model.WILProjectApplication, error) {
	application := &model.WILProjectApplication{}
	result := repo.connector.Where("WILProjectApplicationId = ?", id).First(application)
	return application, result.Error
}

func (repo WILProjectApplicationController) UpdateWILProjectApplication(application *model.WILProjectApplication) error {
	result := repo.connector.Save(application)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
*/

func (controller WILProjectApplicationController) RegisterWILProjectsApplication() error {
	WILProjectApplicationModel := model.WILProjectApplication{}

	fmt.Println("\nRegistering WILProjectApplication model")

	numStudents := int(utils.GetUserInputUint("\nHow many students are in the project? 2 or 3: "))
	var StudentsId []string
	for len(StudentsId) < numStudents {
		studentId := utils.GetUserInput("\nEnter Student ID: ")
		for _, id := range StudentsId {
			if id == studentId {
				fmt.Println("\nStudent ID already exists. Please enter a different ID.")
				continue
			}
		}

		// TODO: Check if the student ID is valid
		// If valid, append to the slice
		// Else continue
		StudentsId = append(StudentsId, studentId)
	}
	WILProjectApplicationModel.ProjectName = utils.GetUserInput("\nEnter Project Name: ")
	WILProjectApplicationModel.ProjectDetail = utils.GetUserInput("\nEnter Project Detail: ")
	WILProjectApplicationModel.Semester = utils.GetUserInput("\nEnter Semester: ")
	WILProjectApplicationModel.CompanyId = uint(utils.GetUserInputUint("\nEnter Company Id: "))
	WILProjectApplicationModel.Mentor = utils.GetUserInput("\nEnter Mentor Name: ")
	WILProjectApplicationModel.AdvisorId = utils.GetUserInputUint("\nEnter Advisor Id: ")

	WILProjectApplicationModel.ApplicationStatus = string(model.WIL_APP_PENDING)
	WILProjectApplicationModel.TurninDate = time.Now().Format("2006-01-02 15:04:05")

	result := controller.connector.Create(&WILProjectApplicationModel)
	if result.Error != nil {
		return result.Error
	}

	for _, studentId := range StudentsId {
		WILProjectMemberModel := model.WILProjectMember{
			WILProjectApplicationId: WILProjectApplicationModel.ID,
			StudentId:               studentId,
			Student: commonModel.Student{
				StudentCode: studentId,
			},
		}
		result := controller.connector.Create(&WILProjectMemberModel)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
