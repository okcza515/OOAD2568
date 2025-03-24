// MEP-1013
package spacemanagement

import (
	"ModEd/asset/model/asset"
	"gorm.io/gorm"
)

type InstrumentManagement struct {
	gorm.Model
	AssetManagementID  uint                   `gorm:"type:integer" json:"asset_management_id" csv:"asset_management_id"`
	BorrowInstrument   asset.BorrowInstrument `gorm:"foreignKey:AssetManagementID;references:AssetManagementID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"borrow_instrument"`
	BorrowUserID       uint       			  `gorm:"not null"`
	RoomID             uint                   `gorm:"type:integer;not null;index" json:"room_id" csv:"room_id"`
	Room               Room                   `gorm:"foreignKey:RoomID;references:RoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room"`
	InstrumentID       uint                   `gorm:"type:integer;not null;index" json:"instrument_id" csv:"instrument_id"`
	Instrument         asset.Instrument       `gorm:"foreignKey:InstrumentID;references:InstrumentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"instrument"`
	InstrumentLabel    string                 `gorm:"type:text;not null" json:"instrument_label" csv:"instrument_label"`
}