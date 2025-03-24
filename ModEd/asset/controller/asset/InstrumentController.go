package asset

import (
	model "ModEd/asset/model/asset"
	"gorm.io/gorm"
	"strings"
)

type InstrumentController struct {
	db *gorm.DB
}

func (c *InstrumentController) Create(instrumentData *[]model.Instrument) error {
	result := c.db.Create(instrumentData)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *InstrumentController) ListAll() ([]string, error) {
	inst := new([]model.Instrument)
	result := c.db.Find(&inst)

	if result.Error != nil {
		return nil, result.Error
	}

	var resultList []string

	for _, i := range *inst {
		list := strings.Join([]string{i.InstrumentCode, i.InstrumentLabel}, "\t")
		resultList = append(resultList, list)
	}

	return resultList, result.Error
}
