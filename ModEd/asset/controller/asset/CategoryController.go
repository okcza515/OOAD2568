package asset

import (
	"ModEd/asset/model/asset"
	"time"

	"gorm.io/gorm"
)

type CategoryController struct {
	Db *gorm.DB
}

func (c *CategoryController) GetAll() (*[]asset.Category, error) {
	categorys := new([]asset.Category)
	result := c.Db.Find(&categorys)
	return categorys, result.Error
}

func (c *CategoryController) GetByID(supplyID uint) (*asset.Category, error) {
	categorys := new(asset.Category)
	result := c.Db.First(&categorys, "ID = ?", supplyID)
	return categorys, result.Error
}

func (c *CategoryController) Create(body *asset.Category) error {
	result := c.Db.Create(body)
	return result.Error
}

func (c *CategoryController) Update(categoryID uint, body *asset.Category) error {
	body.ID = categoryID
	result := c.Db.Updates(body)
	return result.Error
}

func (c *CategoryController) Delete(categoryID uint) error {
	result := c.Db.Model(&asset.Category{}).Where("ID = ?", categoryID).Update("deleted_at", time.Now())
	return result.Error
}
