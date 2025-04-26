package core

// Wrote by MEP-12

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	*SerializableRecord
}

func (base BaseModel) GetID() uint {
	return base.ID
}

func (base BaseModel) ToString() string {
	return fmt.Sprintf("%+v", base)
}

func (base BaseModel) Validate() error {
	return errors.New("err: unimplemented validation")
}
