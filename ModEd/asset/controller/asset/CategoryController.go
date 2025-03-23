package asset

import (
	"ModEd/asset/model/asset"
	"time"

	"gorm.io/gorm"
)

type CategoryController struct {
	db *gorm.DB
}

func (c *CategoryController) GetAll() (*[]asset.Category, error) {
	categories := new([]asset.Category)
	result := c.db.Find(&categories)
	return categories, result.Error
}

func (c *CategoryController) GetByID(supplyID uint) (*asset.Category, error) {
	categories := new(asset.Category)
	result := c.db.First(&categories, "ID = ?", supplyID)
	return categories, result.Error
}

func (c *CategoryController) Create(body *asset.Category) error {
	result := c.db.Create(body)
	return result.Error
}

func (c *CategoryController) Update(categoryID uint, body *asset.Category) error {
	body.ID = categoryID
	result := c.db.Updates(body)
	return result.Error
}

func (c *CategoryController) Delete(categoryID uint) error {
	result := c.db.Model(&asset.Category{}).Where("ID = ?", categoryID).Update("deleted_at", time.Now())
	return result.Error
}
