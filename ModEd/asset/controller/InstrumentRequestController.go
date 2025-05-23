// MEP-1014
package controller

import (
	model "ModEd/asset/model"

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

func (c *InstrumentRequestController) GetInstrumentRequestByName(name string) (*model.InstrumentRequest, error) {
	var request model.InstrumentRequest
	err := c.db.Joins("Department").Where("departments.name = ?", name).First(&request).Error
	return &request, err
}

func (c *InstrumentRequestController) ListAllInstrumentRequests() (*[]model.InstrumentRequest, error) {
	var requests []model.InstrumentRequest
	err := c.db.Where("is_linked_to_tor = ?", false).Find(&requests).Error
	return &requests, err
}

func (c *InstrumentRequestController) GetInstrumentRequestWithDetails(id uint) (*model.InstrumentRequest, error) {
	var request model.InstrumentRequest
	err := c.db.Preload("Instruments").First(&request, id).Error
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

func (c *InstrumentRequestController) UpdateInstrumentDetail(detailID uint, updated *model.InstrumentDetail) error {
	updated.InstrumentDetailID = detailID
	result := c.db.Model(&model.InstrumentDetail{}).
		Where("instrument_detail_id = ?", detailID).
		Updates(updated)

	return result.Error
}

func (c *InstrumentRequestController) RequestApprove(id uint) error {
	return c.db.Model(&model.InstrumentRequest{}).
		Where("instrument_request_id = ?", id).
		Update("status", model.InstrumentRequestStatusApproved).Error // use enum
}

func (c *InstrumentRequestController) DeleteInstrumentRequest(id uint) error {
	return c.db.Delete(&model.InstrumentRequest{}, id).Error
}

func (c *InstrumentRequestController) RemoveInstrumentFromRequest(detailID uint) error {
	return c.db.Delete(&model.InstrumentDetail{}, detailID).Error
}

func (c *InstrumentRequestController) MarkAsUsed(id uint) error {
	return c.db.Model(&model.InstrumentRequest{}).
		Where("instrument_request_id = ?", id).
		Update("is_linked_to_tor", true).Error
}

func (c *InstrumentRequestController) UpdateTotalEstimatedPrice(id uint) error {
	var total float64

	err := c.db.Table("instrument_details").
		Where("instrument_request_id = ?", id).
		Select("SUM(estimated_price * quantity) as total").
		Scan(&total).Error

	if err != nil {
		return err
	}

	return c.db.Model(&model.InstrumentRequest{}).
		Where("instrument_request_id = ?", id).
		Update("total_estimated_price", total).Error
}

func (c *InstrumentRequestController) UpdateInstrumentRequestStatus(id uint, status model.InstrumentRequestStatus) error {
	return c.db.Model(&model.InstrumentRequest{}).
		Where("instrument_request_id = ?", id).
		Update("status", status).Error
}
