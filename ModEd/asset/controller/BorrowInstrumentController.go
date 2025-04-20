// MEP-1012
package controller

import (
	"ModEd/asset/model"
	"gorm.io/gorm"
	"time"
)

type BorrowInstrumentController struct {
	db *gorm.DB
}

func (c *BorrowInstrumentController) GetAll() (*[]model.BorrowInstrument, error) {
	borrowInstrument := new([]model.BorrowInstrument)
	result := c.db.Find(&borrowInstrument)
	return borrowInstrument, result.Error
}

func (c *BorrowInstrumentController) GetByID(ID uint) (*model.BorrowInstrument, error) {
	borrowInstrument := new(model.BorrowInstrument)
	result := c.db.First(&borrowInstrument, "ID = ?", ID)
	return borrowInstrument, result.Error
}

func (c *BorrowInstrumentController) Create(body *model.BorrowInstrument) error {
	result := c.db.Create(body)
	return result.Error
}

func (c *BorrowInstrumentController) Update(ID uint, body *model.BorrowInstrument) error {
	body.ID = ID
	result := c.db.Updates(body)
	return result.Error
}

func (c *BorrowInstrumentController) Delete(ID uint) error {
	result := c.db.Model(&model.BorrowInstrument{}).Where("ID = ?", ID).Update("deleted_at", time.Now())
	return result.Error
}
