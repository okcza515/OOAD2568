// MEP-1012
package asset

import (
	model "ModEd/asset/model/asset"
	"gorm.io/gorm"
	"time"
)

type BorrowInstrumentController struct {
	Db *gorm.DB
}

func (c *BorrowInstrumentController) GetAll() (*[]model.BorrowInstrument, error) {
	borrowInstrument := new([]model.BorrowInstrument)
	result := c.Db.Find(&borrowInstrument)
	return borrowInstrument, result.Error
}

func (c *BorrowInstrumentController) GetByID(ID uint) (*model.BorrowInstrument, error) {
	borrowInstrument := new(model.BorrowInstrument)
	result := c.Db.First(&borrowInstrument, "ID = ?", ID)
	return borrowInstrument, result.Error
}

func (c *BorrowInstrumentController) Create(body *model.BorrowInstrument) error {
	result := c.Db.Create(body)
	return result.Error
}

func (c *BorrowInstrumentController) Update(ID uint, body *model.BorrowInstrument) error {
	body.ID = ID
	result := c.Db.Updates(body)
	return result.Error
}

func (c *BorrowInstrumentController) Delete(ID uint) error {
	result := c.Db.Model(&model.BorrowInstrument{}).Where("ID = ?", ID).Update("deleted_at", time.Now())
	return result.Error
}
