package space

import (
	"github.com/google/uuid"
)

type Room struct {
	RoomID             uuid.UUID         `gorm:"type:text;primaryKey" json:"room_id" csv:"room_id"`
	RoomName           string            `gorm:"type:varchar(255);not null" json:"room_name" csv:"room_name"`
	RoomType           RoomTypeEnum      `gorm:"type:text;not null" json:"room_type" csv:"room_type"`
	Description        *string           `gorm:"type:text" json:"description,omitempty" csv:"description,omitempty"`
	Floor              int               `gorm:"type:integer;not null" json:"floor" csv:"floor"`
	Building           *string           `gorm:"type:varchar(255);not null" json:"building" csv:"building"`
	Location           *string           `gorm:"type:text" json:"location,omitempty" csv:"location,omitempty"`
	Capacity           int               `gorm:"type:integer;not null" json:"capacity" csv:"capacity"`
	IsRoomOutOfService bool              `gorm:"type:boolean;not null" json:"is_room_out_of_service" csv:"is_room_out_of_service"`
	Assets             []AssetManagement `gorm:"foreignKey:RoomID" json:"assets,omitempty" csv:"assets,omitempty"`
}
