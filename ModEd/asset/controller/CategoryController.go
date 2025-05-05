package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"

	"gorm.io/gorm"
)

type CategoryController struct {
	db *gorm.DB
	*core.BaseController[model.Category]
}

type CategoryControllerInterface interface {
	ListAll() ([]string, error)
	List(condition map[string]interface{}, preloads ...string) ([]model.Category, error)
	RetrieveByID(id uint, preloads ...string) (model.Category, error)
	Insert(data model.Category) error
	UpdateByID(data model.Category) error
	DeleteByID(id uint) error
	InsertMany(data []model.Category) error
}

func NewCategoryController() *CategoryController {
	db := migration.GetInstance().DB
	return &CategoryController{
		db:             db,
		BaseController: core.NewBaseController[model.Category](db),
	}
}
