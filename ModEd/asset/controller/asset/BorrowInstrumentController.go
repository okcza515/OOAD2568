// MEP-1012
package asset

import (
	"ModEd/asset/model/asset"
	"gorm.io/gorm"
	"time"
)

type BorrowInstrumentController struct {
	Db *gorm.DB
}

func (c *BorrowInstrumentController) GetAll() (*[]asset.BorrowInstrument, error) {
	borrowInstrument := new([]asset.BorrowInstrument)
	result := c.Db.Find(&borrowInstrument)
	return borrowInstrument, result.Error
}

func (c *BorrowInstrumentController) GetByID(ID uint) (*asset.BorrowInstrument, error) {
	borrowInstrument := new(asset.BorrowInstrument)
	result := c.Db.First(&borrowInstrument, "ID = ?", ID)
	return borrowInstrument, result.Error
}

func (c *BorrowInstrumentController) Create(body *asset.BorrowInstrument) error {
	result := c.Db.Create(body)
	return result.Error
}

func (c *BorrowInstrumentController) Update(ID uint, body *asset.BorrowInstrument) error {
	body.ID = ID
	result := c.Db.Updates(body)
	return result.Error
}

func (c *BorrowInstrumentController) Delete(ID uint) error {
	result := c.Db.Model(&asset.BorrowInstrument{}).Where("ID = ?", ID).Update("deleted_at", time.Now())
	return result.Error
}
