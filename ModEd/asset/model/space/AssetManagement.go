package space

import (
	"ModEd/asset/model/asset"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssetManagement struct {
	gorm.Model
	AssetManagementID  uuid.UUID              `gorm:"type:text;primaryKey"`
	RoomID             uuid.UUID              `gorm:"type:text;not null;index" json:"room_id" csv:"room_id"`
	Room               Room                   `gorm:"foreignKey:RoomID;references:RoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room"`
	InstrumentID       uuid.UUID              `gorm:"type:text;not null;index" json:"instrument_id" csv:"instrument_id"`
	Instrument         asset.Instrument       `gorm:"foreignKey:InstrumentID;references:InstrumentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"instrument"`
	InstrumentLabel    string                 `gorm:"type:text;not null" json:"instrument_label" csv:"instrument_label"`
	SupplyID           uuid.UUID              `gorm:"type:text;not null;index" json:"supply_id" csv:"supply_id"`
	Supply             asset.Supply           `gorm:"foreignKey:SupplyID;references:SupplyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"supply"`
	Quantity           int                    `gorm:"not null" json:"quantity" csv:"quantity"`
	BorrowStatus       int                    `gorm:"type:integer;not null" json:"borrow_status" csv:"borrow_status"`
	BorrowID           uuid.UUID              `gorm:"type:text;not null" json:"borrow_id" csv:"borrow_id"`
	BorrowInstrument   asset.BorrowInstrument `gorm:"foreignKey:BorrowID;references:BorrowID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"borrow_instrument"`
	BorrowDate         time.Time              `gorm:"type:timestamp" json:"borrow_date"`
	ReturnDate         time.Time              `gorm:"type:timestamp" json:"return_date"`
	ExpectedReturnDate time.Time              `gorm:"type:timestamp" json:"expected_return_date"`
	IsLate             bool                   `json:"is_late"`
}
