package model

type Audit struct {
	CreatedBy uint `gorm:"foreignKey:CreatedByID;references:ID;not null"`
	UpdatedBy uint `gorm:"foreignKey:UpdatedByID;references:ID;not null"`
	DeletedBy uint `gorm:"foreignKey:DeletedByID;references:ID"`
}
