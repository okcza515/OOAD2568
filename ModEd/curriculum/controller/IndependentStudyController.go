// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type IndependentStudyController struct {
	*core.BaseController[model.IndependentStudy]
	Connector *gorm.DB
}

type IndependentStudyControllerInterface interface {
	Insert(data core.RecordInterface) error
	UpdateByID(data core.RecordInterface) error
	RetrieveByID(id uint, preloads ...string) (*core.RecordInterface, error)
	List(condition map[string]interface{}) ([]core.RecordInterface, error)
	ListPagination(condition map[string]interface{}, page int, pageSize int) (*core.PaginationResult, error)
}

func CreateIndependentStudyController(connector *gorm.DB) *IndependentStudyController {
	return &IndependentStudyController{
		Connector:      connector,
		BaseController: core.NewBaseController[model.IndependentStudy](connector),
	}
}

func (controller IndependentStudyController) CreateIndependentStudy(independentStudies *[]model.IndependentStudy) error {
	for _, independentStudy := range *independentStudies {
		err := controller.BaseController.Insert(independentStudy)
		if err != nil {
			return err
		}
	}
	return nil
}

// func (repo IndependentStudyController) GetIndenpendentStudyByID(id uint) (*model.IndependentStudy, error) {
// 	independentStudy := new(model.IndependentStudy)
// 	result := repo.Connector.First(&independentStudy, "id = ?", id)
// 	return independentStudy, result.Error
// }

// func (repo IndependentStudyController) GetAllIndenpendentStudy() ([]*model.IndependentStudy, error) {
// 	independentStudy := []*model.IndependentStudy{}
// 	result := repo.Connector.Find(&independentStudy, "DeletedAt IS NULL")
// 	return independentStudy, result.Error
// }

// func (repo IndependentStudyController) UpdateIndenpendentStudy(updatedStudy *model.IndependentStudy) error {
// 	result := repo.Connector.Save(updatedStudy)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
