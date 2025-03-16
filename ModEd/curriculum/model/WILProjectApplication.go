package model

import (
	modelCommon "ModEd/common/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WILProjectApplication struct {
	gorm.Model
	WILProjectApplicationId uuid.UUID
	ProjectName             string
	ProjectDetail           string
	Company                 uuid.UUID
	Mentor                  string
	Students                []modelCommon.Student
	Advisor                 modelCommon.Instructor
	ApplicationStatus       WILApplicationStatusEnum
	TurninDate              time.Time
}
