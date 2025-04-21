package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type IndependentStudyController struct {
	*core.BaseController
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
		BaseController: core.NewBaseController("IndependentStudy", connector),
	}
}

func (repo IndependentStudyController) CreateIndependentStudy(independentStudies *[]model.IndependentStudy) {
	for _, independentStudy := range *independentStudies {
		repo.BaseController.Insert(&independentStudy)
	}
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
