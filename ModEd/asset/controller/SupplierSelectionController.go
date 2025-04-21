// MEP-1014
package controller

import (
	model "ModEd/asset/model"

	"gorm.io/gorm"
)

type SupplierSelectionController struct {
	db *gorm.DB
}

func CreateSupplierSelectionController(db *gorm.DB) *SupplierSelectionController {
	return &SupplierSelectionController{db: db}
}

func (c *SupplierSelectionController) CreateQuotation(body *model.Quotation) error {
	return c.db.Create(body).Error
}

func (c *SupplierSelectionController) CreateSupplier(body *model.Supplier) error {
	return c.db.Create(body).Error
}

func (c *SupplierSelectionController) ListAllQuotation() ([]model.Quotation, error) {
	var quotations []model.Quotation
	err := c.db.Find(&quotations).Error
	return quotations, err
}

func (c *SupplierSelectionController) ListAllSupplier() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	err := c.db.Find(&suppliers).Error
	return suppliers, err
}

func (c *SupplierSelectionController) GetByQuotationID(id uint) (*model.Quotation, error) {
	var quotation model.Quotation
	err := c.db.First(&quotation, id).Error
	return &quotation, err
}

func (c *SupplierSelectionController) GetBySupplierID(id uint) (*model.Supplier, error) {
	var supplier model.Supplier
	err := c.db.First(&supplier, id).Error
	return &supplier, err
}

func (c *SupplierSelectionController) UpdateQuotation(id uint, updated *model.Quotation) error {
	updated.QuotationID = id
	result := c.db.Model(&model.Quotation{}).Where("QuotationID = ?", id).Updates(updated)
	return result.Error
}

func (c *SupplierSelectionController) UpdateSupplier(id uint, updated *model.Supplier) error {
	updated.SupplierID = id
	result := c.db.Model(&model.Supplier{}).Where("SupplierID = ?", id).Updates(updated)
	return result.Error
}

func (c *SupplierSelectionController) DeleteByQuotationID(id uint) error {
	return c.db.Delete(&model.Quotation{}, id).Error
}

func (c *SupplierSelectionController) DeleteBySupplierID(id uint) error {
	return c.db.Delete(&model.Supplier{}, id).Error
}
