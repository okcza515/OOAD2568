// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
	model "ModEd/curriculum/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WILProjectApplicationController struct {
	connector *gorm.DB
	*core.BaseController[model.WILProjectApplication]
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
	return &WILProjectApplicationController{
		connector:      connector,
		BaseController: core.NewBaseController[model.WILProjectApplication](connector),
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

	for _, studentId := range studentIds {
		WILProjectMemberModel := model.WILProjectMember{
			WILProjectApplicationId: wilprojectApplication.ID,
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

func (controller WILProjectApplicationController) ListWILProjectApplication() ([]model.WILProjectApplication, error) {
	var applications []model.WILProjectApplication

	result := controller.connector.
		Preload(clause.Associations).
		Find(&applications)

	if result.Error != nil {
		return nil, result.Error
	}

	return applications, nil
}
