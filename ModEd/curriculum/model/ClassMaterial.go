// MEP-1008
package model

import (
	"ModEd/core"
	"fmt"

	"gorm.io/gorm"
)

type ClassMaterial struct {
	gorm.Model
	ClassId  uint   `gorm:"not null" json:"class_id"`
	Class    Class  `gorm:"foreignKey:ClassId;references:ClassId"`
	FileName string `gorm:"type:varchar(100);not null"`
	FilePath string `gorm:"type:varchar(255);not null"`
	*core.SerializableRecord
}

func (cm *ClassMaterial) GetID() uint {
	return cm.ClassId
}

func (cm *ClassMaterial) ToString() string {
	return fmt.Sprintf("%+v", cm)
}

func (cm *ClassMaterial) Validate() error {
	if cm.ClassId == 0 {
		return fmt.Errorf("Class ID cannot be zero")
	}
	if cm.FileName == "" {
		return fmt.Errorf("File name cannot be empty")
	}
	if cm.FilePath == "" {
		return fmt.Errorf("File path cannot be empty")
	}
	return nil
}
