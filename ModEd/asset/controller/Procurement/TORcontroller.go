// MEP-1014
package controller

import (
	model "ModEd/asset/model/Procurement"
	
	"gorm.io/gorm"
)

type TORController struct {
	Connector *gorm.DB
}

func CreateTORController(connector *gorm.DB) *TORController {
	tor := TORController{Connector: connector}
	connector.AutoMigrate(&model.TOR{})
	return &tor
}

func (tor TORController) ListAll() ([]model.TOR, error) {
	tors := []model.TOR{}
	result := tor.Connector.
		Select("TORID").Find(&tors)
	return tors, result.Error
}

func (tor TORController) GetByID(TORID uint) (*model.TOR, error) {
	t := &model.TOR{}
	result := tor.Connector.Where("TORID = ?", TORID).First(t)
	return t, result.Error
}

func (tor TORController) Create(t *model.TOR) error {
	return tor.Connector.Create(t).Error
}

func (tor TORController) Update(TORID uint, updatedData map[string]interface{}) error {
	return tor.Connector.Model(&model.TOR{}).Where("TORID = ?", TORID).Updates(updatedData).Error
}

func (tor TORController) DeleteByID(TORID uint) error {
	return tor.Connector.Where("TORID = ?", TORID).Delete(&model.TOR{}).Error
}
