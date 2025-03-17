package model

import (
	"github.com/google/uuid"
)

type Classroom struct {
	ClassroomID        uuid.UUID         `gorm:"type:text;primaryKey" json:"classroom_id" csv:"classroom_id"`
	ClassroomName      string            `gorm:"type:varchar(255);not null" json:"classroom_name" csv:"classroom_name"`
	Description        *string           `gorm:"type:text" json:"description,omitempty" csv:"description,omitempty"`
	Floor              int               `gorm:"type:integer;not null" json:"floor" csv:"floor"`
	Building           *string           `gorm:"type:varchar(255);not null" json:"building" csv:"building"`
	Location           *string           `gorm:"type:text" json:"location,omitempty" csv:"location,omitempty"`
	Capacity           int               `gorm:"type:integer;not null" json:"capacity" csv:"capacity"`
	IsRoomOutOfService bool              `gorm:"type:boolean;not null" json:"is_room_out_of_service" csv:"is_room_out_of_service"`
	RoomType           RoomTypeEnum      `gorm:"type:text;not null" json:"room_type" csv:"room_type"`
	Assets             []AssetManagement `gorm:"foreignKey:RoomID" json:"assets,omitempty" csv:"assets,omitempty"`
}
