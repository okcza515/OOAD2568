package asset

import (
	"ModEd/asset/model/asset"
	"gorm.io/gorm"
)

type BorrowInstrumentController struct {
	Db *gorm.DB
}

func (c *BorrowInstrumentController) GetAll() (*[]asset.BorrowInstrument, error) {
	borrowInstrument := new([]asset.BorrowInstrument)
	result := c.Db.Find(&borrowInstrument)
	return borrowInstrument, result.Error
}
