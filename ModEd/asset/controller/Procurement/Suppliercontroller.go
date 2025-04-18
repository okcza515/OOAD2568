// MEP-1014
package procurement

import (
	model "ModEd/asset/model/Procurement"
	
	"gorm.io/gorm"
)

type SupplierController struct {
	Connector *gorm.DB
}

func CreateSupplierController(connector *gorm.DB) *SupplierController {
	supplier := SupplierController{Connector: connector}
	connector.AutoMigrate(&model.Supplier{})
	return &supplier
}

func (supplier SupplierController) ListAll() ([]model.Supplier, error) {
	suppliers := []model.Supplier{}
	result := supplier.Connector.
	Select("SupplierID").Find(&suppliers)
	return suppliers, result.Error
}

func (supplier SupplierController) GetByID(SupplierID uint) (*model.Supplier, error) {
	s := &model.Supplier{}
	result := supplier.Connector.Where("SupplierID = ?", SupplierID).First(s)
	return s, result.Error
}

func (supplier SupplierController) Create(s *model.Supplier) error {
	return supplier.Connector.Create(s).Error
}

func (supplier SupplierController) Update(SupplierID uint, updatedData map[string]interface{}) error {
	return supplier.Connector.Model(&model.Supplier{}).Where("SupplierID = ?", SupplierID).Updates(updatedData).Error
}

func (supplier SupplierController) DeleteByID(SupplierID uint) error {
	return supplier.Connector.Where("SupplierID = ?", SupplierID).Delete(&model.Supplier{}).Error
}