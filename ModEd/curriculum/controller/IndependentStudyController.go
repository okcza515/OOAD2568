// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"

	"time"

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

func (controller IndependentStudyController) CreateIndependentStudy(independentStudy *model.IndependentStudy, turnInDate string) error {

	if turnInDate != "" {
		t, err := time.Parse("2006-01-02 15:04:05", turnInDate)
		if err != nil {
			return err
		}
		independentStudy.TurnInDate = &t
	}

	independentStudy.CreatedAt = time.Now()
	independentStudy.UpdatedAt = time.Now()
	if err := controller.BaseController.Insert(*independentStudy); err != nil {
		return err
	}
	return nil
}
