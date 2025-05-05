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
	Insert(data model.IndependentStudy) error
	InsertMany(data []model.IndependentStudy) error
	UpdateByID(data model.IndependentStudy) error
	RetrieveByID(id uint, preloads ...string) (*model.IndependentStudy, error)
	List(condition map[string]interface{}) ([]model.IndependentStudy, error)
	ListPagination(condition map[string]interface{}, page int, pageSize int) (*core.PaginationResult, error)
}

func NewIndependentStudyController(connector *gorm.DB) *IndependentStudyController {
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

func (controller IndependentStudyController) ListAllIndependentStudy() ([]model.IndependentStudy, error) {
	var independentStudies []model.IndependentStudy
	result := controller.Connector.
		Preload("WILProject").
		Find(&independentStudies)
	if result.Error != nil {
		return nil, result.Error
	}
	return independentStudies, nil
}

func (controller IndependentStudyController) DeleteIndependentStudiesByID(id uint) error {
	/*
		for implement more complex business logic
	*/
	if err := controller.BaseController.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
