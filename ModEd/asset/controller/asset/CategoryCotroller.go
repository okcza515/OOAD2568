package asset

import (
	"ModEd/asset/model/asset"
	"time"

	"gorm.io/gorm"
)

type CategoryCotroller struct {
	Db *gorm.DB
}

func (c *CategoryCotroller) GetAll() (*[]asset.Category, error) {
	categorys := new([]asset.Category)
	result := c.Db.Find(&categorys)
	return categorys, result.Error
}

func (c *CategoryCotroller) GetByID(supplyID uint) (*asset.Category, error) {
	categorys := new(asset.Category)
	result := c.Db.First(&categorys, "ID = ?", supplyID)
	return categorys, result.Error
}

func (c *CategoryCotroller) Create(body *asset.Category) error {
	result := c.Db.Create(body)
	return result.Error
}

func (c *CategoryCotroller) Update(categoryID uint, body *asset.Category) error {
	body.ID = categoryID
	result := c.Db.Updates(body)
	return result.Error
}

func (c *CategoryCotroller) Delete(categoryID uint) error {
	result := c.Db.Model(&asset.Category{}).Where("ID = ?", categoryID).Update("deleted_at", time.Now())
	return result.Error
}
