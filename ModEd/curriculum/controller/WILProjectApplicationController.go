// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
	model "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type WILProjectApplicationController struct {
	connector                  *gorm.DB
	wilProjectMemberController *WILProjectMemberController

	*core.BaseController[model.WILProjectApplication]
}

type WILProjectApplicationControllerInterface interface {
	RegisterWILProjectsApplication(projects []model.WILProjectApplication)
	Insert(data model.WILProjectApplication) error
	UpdateByID(data model.WILProjectApplication) error
	RetrieveByID(id uint, preloads ...string) (*model.WILProjectApplication, error)
	DeleteByID(id uint) error
	ListPagination(condition map[string]interface{}, page int, pageSize int)
}

func NewWILProjectApplicationController(connector *gorm.DB) *WILProjectApplicationController {
	return &WILProjectApplicationController{
		connector:                  connector,
		wilProjectMemberController: NewWILProjectMemberController(connector),
		BaseController:             core.NewBaseController[model.WILProjectApplication](connector),
	}
}

func (controller WILProjectApplicationController) RegisterWILProjectsApplication(
	wilprojectApplication model.WILProjectApplication,
	studentIds []string,
) error {

	resultError := controller.Insert(wilprojectApplication)
	if resultError != nil {
		return resultError
	}

	var wilMembers []model.WILProjectMember
	for _, studentId := range studentIds {
		member := model.WILProjectMember{
			WILProjectApplicationId: wilprojectApplication.ID,
			StudentId:               studentId,
			Student: commonModel.Student{
				StudentCode: studentId,
			},
		}
		wilMembers = append(wilMembers, member)

	}

	err := controller.wilProjectMemberController.InsertMany(wilMembers)

	if err != nil {
		return err
	}

	return nil
}

func (controller WILProjectApplicationController) ListWILProjectApplication() ([]model.WILProjectApplication, error) {
	var applications []model.WILProjectApplication

	result := controller.connector.
		Preload("Advisor").
		Preload("Students").
		Find(&applications)

	if result.Error != nil {
		return nil, result.Error
	}

	return applications, nil
}
