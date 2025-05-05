package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"

	"gorm.io/gorm"
)

type BorrowInstrumentController struct {
	db *gorm.DB
	*core.BaseController[model.BorrowInstrument]

	observers map[string]AssetObserver[model.BorrowInstrument]
}

type BorrowInstrumentControllerInterface interface {
	ListAll() ([]string, error)
	List(condition map[string]interface{}, preloads ...string) ([]model.BorrowInstrument, error)
	RetrieveByID(id uint, preloads ...string) (model.BorrowInstrument, error)
	Insert(data model.BorrowInstrument) error
	UpdateByID(data model.BorrowInstrument) error
	DeleteByID(id uint) error
	//GetAll() ([]model.BorrowInstrument, error)
	//GetByID(ID uint) (*model.BorrowInstrument, error)
	//Create(body *model.BorrowInstrument) error
	InsertMany(data []model.BorrowInstrument) error
	//Update(ID uint, body *model.BorrowInstrument) error
	//Delete(ID uint) error
}

func NewBorrowInstrumentController() *BorrowInstrumentController {
	db := migration.GetInstance().DB
	return &BorrowInstrumentController{
		db:             db,
		BaseController: core.NewBaseController[model.BorrowInstrument](db),
	}
}

//func (c *BorrowInstrumentController) GetAll() ([]model.BorrowInstrument, error) {
//	var borrowInstruments []model.BorrowInstrument
//	result := c.db.Find(&borrowInstruments)
//	return borrowInstruments, result.Error
//}

//func (c *BorrowInstrumentController) GetByID(ID uint) (*model.BorrowInstrument, error) {
//	borrowInstrument := new(model.BorrowInstrument)
//	result := c.db.First(&borrowInstrument, "ID = ?", ID)
//	return borrowInstrument, result.Error
//}

func (c *BorrowInstrumentController) InsertMany(data []model.BorrowInstrument) error {
	result := c.db.Create(&data)
	return result.Error
}
