package model

type RequestResignationStudent struct {
	BaseStandardRequest
	StudentCode string `gorm:"type:text;default:'';not null"`
}
