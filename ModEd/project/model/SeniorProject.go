package model

import "gorm.io/gorm"

type SeniorProject struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey"`
	GroupName string `gorm:"not null"`

	Members []GroupMember `gorm:"foreignKey:SeniorProjectID"`
}
