package space

import (
	"ModEd/asset/model/asset"
	"gorm.io/gorm"
)

type AssetManagement struct {
	gorm.Model
	AssetManagementID  uint                   `gorm:"type:int;primaryKey"`
	RoomID             uint                   `gorm:"type:int;not null;index" json:"room_id" csv:"room_id"`
	Room               Room                   `gorm:"foreignKey:RoomID;references:RoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room"`
	InstrumentID       uint                   `gorm:"type:text;not null;index" json:"instrument_id" csv:"instrument_id"`
	Instrument         asset.Instrument       `gorm:"foreignKey:InstrumentID;references:InstrumentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"instrument"`
	InstrumentLabel    string                 `gorm:"type:text;not null" json:"instrument_label" csv:"instrument_label"`
	Supply             asset.Supply           `gorm:"foreignKey:SupplyID;references:SupplyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"supply"`
	Quantity           int                    `gorm:"not null" json:"quantity" csv:"quantity"`
}