// MEP-1013
package model

import (
	"ModEd/core"
	"fmt"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

type Room struct {
	core.BaseModel
	RoomName           string                 `gorm:"type:varchar(255);not null" json:"room_name" csv:"room_name" validate:"required"`
	RoomType           RoomTypeEnum           `gorm:"type:text;not null" json:"room_type" csv:"room_type" validate:"required"`
	Description        string                 `gorm:"type:text" json:"description,omitempty" csv:"description,omitempty" validate:"required"`
	Floor              int                    `gorm:"type:integer;not null" json:"floor" csv:"floor" validate:"required"`
	Building           string                 `gorm:"type:varchar(255);not null" json:"building" csv:"building" validate:"required"`
	Location           string                 `gorm:"type:text" json:"location,omitempty" csv:"location,omitempty" validate:"required"`
	Capacity           int                    `gorm:"type:integer;not null" json:"capacity" csv:"capacity" validate:"required"`
	IsRoomOutOfService bool                   `gorm:"type:boolean;not null" json:"is_room_out_of_service" csv:"is_room_out_of_service" validate:"required"`
	Instrument         []InstrumentManagement `gorm:"foreignKey:RoomID" json:"instruments,omitempty" csv:"instruments,omitempty" validate:"-"`
	Supply             []SupplyManagement     `gorm:"foreignKey:RoomID" json:"supplies,omitempty" csv:"supplies,omitempty" validate:"-"`
}

func (room Room) ToString() string {
	truncate := func(s string, maxLen int) string {
		if len(s) > maxLen {
			return s[:maxLen-3] + "..."
		}
		return s
	}

	idWidth := 5
	roomNameWidth := 20
	roomTypeWidth := 18
	descriptionWidth := 18
	floorWidth := 5
	buildingWidth := 20
	locationWidth := 18
	capacityWidth := 8
	outOfServiceWidth := 15

	roomName := truncate(room.RoomName, roomNameWidth)
	roomType := truncate(string(room.RoomType), roomTypeWidth)
	description := truncate(room.Description, descriptionWidth)
	building := truncate(room.Building, buildingWidth)
	location := truncate(room.Location, locationWidth)

	headerBorder := "+" +
		strings.Repeat("-", idWidth+2) + "+" +
		strings.Repeat("-", roomNameWidth+2) + "+" +
		strings.Repeat("-", roomTypeWidth+2) + "+" +
		strings.Repeat("-", descriptionWidth+2) + "+" +
		strings.Repeat("-", floorWidth+2) + "+" +
		strings.Repeat("-", buildingWidth+2) + "+" +
		strings.Repeat("-", locationWidth+2) + "+" +
		strings.Repeat("-", capacityWidth+2) + "+" +
		strings.Repeat("-", outOfServiceWidth+2) + "+"

	headerRow := fmt.Sprintf("| %-*s | %-*s | %-*s | %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |",
		idWidth, "ID", roomNameWidth, "Room Name", roomTypeWidth, "Room Type", descriptionWidth, "Description",
		floorWidth, "Floor", buildingWidth, "Building", locationWidth, "Location", capacityWidth, "Capacity", outOfServiceWidth, "Out Of Service")

	dataRow := fmt.Sprintf("| %-*d | %-*s | %-*s | %-*s | %-*d | %-*s | %-*s | %-*d | %-*t |",
		idWidth, room.ID, roomNameWidth, roomName, roomTypeWidth, roomType, descriptionWidth, description,
		floorWidth, room.Floor, buildingWidth, building, locationWidth, location, capacityWidth, room.Capacity, outOfServiceWidth, room.IsRoomOutOfService)

	var result string
	var printHeaderOnce sync.Once

	printHeaderOnce.Do(func() {
		result += headerBorder + "\n"
		result += headerRow + "\n"
		result += headerBorder + "\n"
	})
	result += dataRow + "\n"
	result += headerBorder + "\n"

	return result
}

func (room Room) Validate() error {
	validate := validator.New()
	if err := validate.Struct(room); err != nil {
		return err
	}
	return nil
}
