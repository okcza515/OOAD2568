package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"time"

	"gorm.io/gorm"
)

type CategoryController struct {
	db *gorm.DB
	*core.BaseController
}

type CategoryControllerInterface interface {
	getAll() ([]model.InstrumentLog, error)
	ListAll() ([]string, error)
	Insert(data core.RecordInterface) error
	RetrieveByID(id uint, preloads ...string) (*core.RecordInterface, error)
}

func (c *CategoryController) GetAll() (*[]model.Category, error) {
	categories := new([]model.Category)
	result := c.db.Find(&categories)
	return categories, result.Error
}

func (c *CategoryController) GetByID(supplyID uint) (*model.Category, error) {
	categories := new(model.Category)
	result := c.db.First(&categories, "ID = ?", supplyID)
	return categories, result.Error
}

func (c *CategoryController) Create(body *model.Category) error {
	result := c.db.Create(body)
	return result.Error
}

func (c *CategoryController) Update(categoryID uint, body *model.Category) error {
	body.ID = categoryID
	result := c.db.Updates(body)
	return result.Error
}

func (c *CategoryController) Delete(categoryID uint) error {
	result := c.db.Model(&model.Category{}).Where("ID = ?", categoryID).Update("deleted_at", time.Now())
	return result.Error
}
