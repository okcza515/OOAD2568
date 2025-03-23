// MEP-1014
package controller

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
		Select("ProcurementApprovalWorkflowID").Find(&procurements)
	return procurements, result.Error
}

func (procurement ProcurementController) GetByApprovalId(ProcurementApprovalWorkflowID string) (*model.Procurement, error) {
	i := &model.Procurement{}
	result := procurement.Connector.Where("ProcurementApprovalWorkflowID = ?", ProcurementApprovalWorkflowID).First(i)
	return i, result.Error
}

func (procurement ProcurementController) Create(i *model.Procurement) error {
	return procurement.Connector.Create(i).Error
}

func (procurement ProcurementController) Update(ProcurementApprovalWorkflowID string, updatedData map[string]any) error {
	return procurement.Connector.Model(&model.Procurement{}).
		Where("ProcurementApprovalWorkflowID = ?", ProcurementApprovalWorkflowID).
		Updates(updatedData).Error
}

func (procurement ProcurementController) DeleteByInstructorId(ProcurementApprovalWorkflowID string) error {
	return procurement.Connector.Where("ProcurementApprovalWorkflowID = ?", ProcurementApprovalWorkflowID).Delete(&model.Procurement{}).Error
}
