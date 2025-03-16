package controller

import (
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type WIL struct {
	Connector *gorm.DB
}

func CreateWIL(connector *gorm.DB) *WIL {
	wil := WIL{Connector: connector}
	connector.AutoMigrate(&model.WILProjectApplication{}, &model.WILProject{}, &model.IndependentStudy{})
	return &wil
}
