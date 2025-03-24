// MEP-1013
package spacemanagement

import (
	model "ModEd/asset/model/SpaeManagement"
	"gorm.io/gorm"
	"time"
)

type PermanentScheduleController struct {
	DB *gorm.DB
}

func 