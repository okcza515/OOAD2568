// MEP-1014
package procurement

import (
	model "ModEd/asset/model/Procurement"

	"gorm.io/gorm"
)

type ItemRequestController struct {
	db *gorm.DB
}

func CreateItemRequestController(db *gorm.DB) *ItemRequestController {
	return &ItemRequestController{db: db}
}

func (c *ItemRequestController) CreateItemRequest(body *model.ItemRequest) error {
	result := c.db.Create(body)
	return result.Error
}

func (c *ItemRequestController) AddItemToRequest(itemRequestID uint, detail *model.ItemDetail) error {
	detail.ItemRequestID = itemRequestID
	result := c.db.Create(detail)
	return result.Error
}

func (c *ItemRequestController) GetItemRequestByID(id uint) (*model.ItemRequest, error) {
	var request model.ItemRequest
	err := c.db.First(&request, id).Error
	return &request, err
}

func (c *ItemRequestController) ListAllItemRequests() (*[]model.ItemRequest, error) {
	var requests []model.ItemRequest
	err := c.db.Find(&requests).Error
	return &requests, err
}

func (c *ItemRequestController) GetItemRequestWithDetails(id uint) (*model.ItemRequest, error) {
	var request model.ItemRequest
	err := c.db.Preload("Items").
		Preload("BudgetApproval").
		First(&request, id).Error
	return &request, err
}

func (c *ItemRequestController) UpdateItemRequest(id uint, updated *model.ItemRequest) error {
	updated.ID = id
	result := c.db.Model(&model.ItemRequest{}).Where("id = ?", id).Updates(updated)
	return result.Error
}

func (c *ItemRequestController) SubmitForApproval(id uint) error {
	result := c.db.Model(&model.ItemRequest{}).
		Where("id = ?", id).
		Update("status", "Pending")
	return result.Error
}

func (c *ItemRequestController) DeleteItemRequest(id uint) error {
	result := c.db.Delete(&model.ItemRequest{}, id)
	return result.Error
}

func (c *ItemRequestController) RemoveItemFromRequest(detailID uint) error {
	result := c.db.Delete(&model.ItemDetail{}, detailID)
	return result.Error
}
