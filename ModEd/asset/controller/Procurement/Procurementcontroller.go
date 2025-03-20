// MEP-1014
package controller

import (
	"ModEd/asset/model"

	"gorm.io/gorm"
)

type ProcurementController struct {
	Connector *gorm.DB
}

func CreateProcurementController(connector *gorm.DB) *ProcurementController {
	procurement := ProcurementController{Connector: connector}
	//TODO connector.AutoMigrate(&model.Procurement{})
	return &procurement
}

func (procurement ProcurementController) ListAll() ([]model.Procurement, error) {
	Approval := []*model.ItemApprovalWorkFlow{}
	result := procurement.Connector.Find(&Approval)
	return Approval, result.Error
}

func (procurement ProcurementController) GetByApprovalId(ItemApprovalStatusID string) (*model.Procurement, error) {
	i := &model.Procurement{}
	result := procurement.Connector.Where("ItemApprovalStatusID = ?", ItemApprovalStatusID).First(i)
	return i, result.Error
}

func (procurement ProcurementController) Create(i *model.Procurement) error {
	return procurement.Connector.Create(i).Error
}

func (procurement ProcurementController) Update(ItemApprovalStatusID string, updatedData map[string]any) error {
	return procurement.Connector.Model(&model.Procurement{}).
		Where("ItemApprovalStatusID = ?", ItemApprovalStatusID).
		Updates(updatedData).Error
}

func (procurement ProcurementController) DeleteByInstructorId(ItemApprovalStatusID string) error {
	return procurement.Connector.Where("ItemApprovalStatusID = ?", ItemApprovalStatusID).Delete(&model.Procurement{}).Error
}
