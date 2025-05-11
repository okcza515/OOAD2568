// MEP-1013
package model

import (
	"ModEd/core"
	"fmt"
	"strings"
	"sync"
)

type Room struct {
	core.BaseModel
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
