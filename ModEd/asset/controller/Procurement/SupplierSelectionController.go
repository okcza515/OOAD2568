// MEP-1014
package procurement

import (
	model "ModEd/asset/model/Procurement"

	"gorm.io/gorm"
)

type SupplierSelectionController struct {
	Connector *gorm.DB
}

func (quotation SupplierSelectionController) CreateQuotation(s *model.Quotation) error {
	return quotation.Connector.Create(s).Error
}

func (supplier SupplierSelectionController) CreateSupplier(s *model.Supplier) error {
	return supplier.Connector.Create(s).Error
}

func (quotation SupplierSelectionController) ListAllQuotation() ([]model.Quotation, error) {
	quotations := []model.Quotation{}
	result := quotation.Connector.
		Select("QuotationID").Find(&quotations)
	return quotations, result.Error
}

func (supplier SupplierSelectionController) ListAllSupplier() ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	result := supplier.Connector.
		Select("SupplierID").Find(&suppliers)
	return suppliers, result.Error
}

func (quotation SupplierSelectionController) GetByQuotationID(QuotationID uint) (*model.Quotation, error) {
	s := &model.Quotation{}
	result := quotation.Connector.Where("QuotationID = ?", QuotationID).First(s)
	return s, result.Error
}

func (supplier SupplierSelectionController) GetBySupplierID(SupplierID uint) (*model.Supplier, error) {
	s := &model.Supplier{}
	result := supplier.Connector.Where("SupplierID = ?", SupplierID).First(s)
	return s, result.Error
}

func (quotation SupplierSelectionController) UpdateQuotation(QuotationID uint, updatedData map[string]interface{}) error {
	return quotation.Connector.Model(&model.Quotation{}).Where("QuotationID = ?", QuotationID).Updates(updatedData).Error
}

func (supplier SupplierSelectionController) UpdateSupplier(SupplierID uint, updatedData map[string]interface{}) error {
	return supplier.Connector.Model(&model.Supplier{}).Where("SupplierID = ?", SupplierID).Updates(updatedData).Error
}

func (quotation SupplierSelectionController) DeleteByQuotationID(QuotationID uint) error {
	return quotation.Connector.Where("QuotationID = ?", QuotationID).Delete(&model.Quotation{}).Error
}

func (supplier SupplierSelectionController) DeleteBySupplierID(SupplierID uint) error {
	return supplier.Connector.Where("SupplierID = ?", SupplierID).Delete(&model.Supplier{}).Error
}
