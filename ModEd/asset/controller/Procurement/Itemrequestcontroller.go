// MEP-1014
package controller

import (
	model "ModEd/asset/model/Procurement"
	
	"gorm.io/gorm"
)

type ItemRequestController struct {
	Connector *gorm.DB
}

func CreateItemRequestController(connector *gorm.DB) *ItemRequestController {
	itemRequest := ItemRequestController{Connector: connector}
	connector.AutoMigrate(&model.ItemRequest{})
	return &itemRequest
}

func (itemRequest ItemRequestController) ListAll() ([]model.ItemRequest, error) {
	requests := []model.ItemRequest{}
	result := itemRequest.Connector.
		Select("ItemRequestID").Find(&requests)
	return requests, result.Error
}

func (itemRequest ItemRequestController) GetByID(ItemRequestID uint) (*model.ItemRequest, error) {
	r := &model.ItemRequest{}
	result := itemRequest.Connector.Where("ItemRequestID = ?", ItemRequestID).First(r)
	return r, result.Error
}

func (itemRequest ItemRequestController) Create(r *model.ItemRequest) error {
	return itemRequest.Connector.Create(r).Error
}

func (itemRequest ItemRequestController) Update(ItemRequestID uint, updatedData map[string]interface{}) error {
	return itemRequest.Connector.Model(&model.ItemRequest{}).Where("ItemRequestID = ?", ItemRequestID).Updates(updatedData).Error
}

func (itemRequest ItemRequestController) DeleteByID(ItemRequestID uint) error {
	return itemRequest.Connector.Where("ItemRequestID = ?", ItemRequestID).Delete(&model.ItemRequest{}).Error
}
