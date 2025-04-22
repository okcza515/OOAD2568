// MEP-1013
package spacemanagement

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomName           string                 `gorm:"type:varchar(255);not null" json:"room_name" csv:"room_name"`
	RoomType           RoomTypeEnum           `gorm:"type:text;not null" json:"room_type" csv:"room_type"`
	Description        string                 `gorm:"type:text" json:"description,omitempty" csv:"description,omitempty"`
	Floor              int                    `gorm:"type:integer;not null" json:"floor" csv:"floor"`
	Building           string                 `gorm:"type:varchar(255);not null" json:"building" csv:"building"`
	Location           string                 `gorm:"type:text" json:"location,omitempty" csv:"location,omitempty"`
	Capacity           int                    `gorm:"type:integer;not null" json:"capacity" csv:"capacity"`
	IsRoomOutOfService bool                   `gorm:"type:boolean;not null" json:"is_room_out_of_service" csv:"is_room_out_of_service"`
	Instrument         []InstrumentManagement `gorm:"foreignKey:RoomID" json:"instruments,omitempty" csv:"instruments,omitempty"`
	Supply             []SupplyManagement     `gorm:"foreignKey:RoomID" json:"supplies,omitempty" csv:"supplies,omitempty"`
}
