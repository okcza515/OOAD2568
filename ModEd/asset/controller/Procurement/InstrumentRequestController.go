// MEP-1014
package procurement

import (
	model "ModEd/asset/model/Procurement"

	"gorm.io/gorm"
)

type InstrumentRequestController struct {
	db *gorm.DB
}

func CreateInstrumentRequestController(db *gorm.DB) *InstrumentRequestController {
	return &InstrumentRequestController{db: db}
}

func (c *InstrumentRequestController) CreateInstrumentRequest(body *model.InstrumentRequest) error {
	return c.db.Create(body).Error
}

func (c *InstrumentRequestController) AddInstrumentToRequest(requestID uint, detail *model.InstrumentDetail) error {
	detail.InstrumentRequestID = requestID
	return c.db.Create(detail).Error
}

func (c *InstrumentRequestController) GetInstrumentRequestByID(id uint) (*model.InstrumentRequest, error) {
	var request model.InstrumentRequest
	err := c.db.First(&request, id).Error
	return &request, err
}

func (c *InstrumentRequestController) ListAllInstrumentRequests() (*[]model.InstrumentRequest, error) {
	var requests []model.InstrumentRequest
	err := c.db.Find(&requests).Error
	return &requests, err
}

func (c *InstrumentRequestController) GetInstrumentRequestWithDetails(id uint) (*model.InstrumentRequest, error) {
	var request model.InstrumentRequest
	err := c.db.Preload("Instruments").Preload("BudgetApproval").First(&request, id).Error
	return &request, err
}

func (c *InstrumentRequestController) GetRequestsByStatus(status model.InstrumentRequestStatus) (*[]model.InstrumentRequest, error) {
	var requests []model.InstrumentRequest
	err := c.db.Where("status = ?", status).Find(&requests).Error
	return &requests, err
}

func (c *InstrumentRequestController) UpdateInstrumentRequest(id uint, updated *model.InstrumentRequest) error {
	updated.InstrumentRequestID = id
	result := c.db.Model(&model.InstrumentRequest{}).Where("instrument_request_id = ?", id).Updates(updated)
	return result.Error
}

func (c *InstrumentRequestController) SubmitForApproval(id uint) error {
	return c.db.Model(&model.InstrumentRequest{}).
		Where("instrument_request_id = ?", id).
		Update("status", model.InstrumentRequestStatusPending).Error // use enum
}

func (c *InstrumentRequestController) DeleteInstrumentRequest(id uint) error {
	return c.db.Delete(&model.InstrumentRequest{}, id).Error
}

func (c *InstrumentRequestController) RemoveInstrumentFromRequest(detailID uint) error {
	return c.db.Delete(&model.InstrumentDetail{}, detailID).Error
}
