// MEP-1014
package procurement

import (
	model "ModEd/asset/model/Procurement"

	"gorm.io/gorm"
)

type ProcurementController struct {
	Connector *gorm.DB
}

func CreateProcurementController(connector *gorm.DB) *ProcurementController {
	procurement := ProcurementController{Connector: connector}
	connector.AutoMigrate(&model.Procurement{})
	return &procurement
}

func (procurement ProcurementController) ListAll() ([]model.Procurement, error) {
	procurements := []model.Procurement{}
	result := procurement.Connector.
		Select("ProcurementApprovalID").Find(&procurements)
	return procurements, result.Error
}

func (procurement ProcurementController) GetByID(ProcurementApprovalID uint) (*model.Procurement, error) {
	i := &model.Procurement{}
	result := procurement.Connector.Where("ProcurementApprovalID = ?", ProcurementApprovalID).First(i)
	return i, result.Error
}

func (procurement ProcurementController) Create(i *model.Procurement) error {
	return procurement.Connector.Create(i).Error
}

func (procurement ProcurementController) Update(ProcurementApprovalID uint, updatedData map[string]any) error {
	return procurement.Connector.Model(&model.Procurement{}).
		Where("ProcurementApprovalID = ?", ProcurementApprovalID).
		Updates(updatedData).Error
}

func (procurement ProcurementController) DeleteByInstructorID(ProcurementApprovalID uint) error {
	return procurement.Connector.Where("ProcurementApprovalID = ?", ProcurementApprovalID).Delete(&model.Procurement{}).Error
}
