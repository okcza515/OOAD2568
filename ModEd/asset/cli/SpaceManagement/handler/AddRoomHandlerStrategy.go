// MEP-1013
package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
)

type AddRoomHandlerStrategy struct {
	controller interface {
		Insert(dataContext model.Room) error
	}
}

func NewAddRoomHandlerStrategy(
	controller interface {
		Insert(dataContext model.Room) error
	},
) *AddRoomHandlerStrategy {
	return &AddRoomHandlerStrategy{controller: controller}
}

func (handler AddRoomHandlerStrategy) Execute() error {
	fmt.Println("------- Add new Room -------")
	fmt.Println("Please enter the name of the new Room:")
	roomName := util.GetCommandInput()

	fmt.Println("Please enter the room type: (Lecture / Laboratory / Office)")
	roomTypeStr := util.GetCommandInput()
	var roomType model.RoomTypeEnum
	switch roomTypeStr {
	case "Lecture":
		roomType = "ROOM_LECTURE_ROOM"
	case "Laboratory":
		roomType = "ROOM_LAB_ROOM"
	case "Office":
		roomType = "ROOM_MEETING_ROOM"
	default:
		fmt.Println("Invalid room type. Using default type.")
		roomType = "ROOM_LECTURE_ROOM"
	}

	fmt.Println("Please enter the description:")
	description := util.GetCommandInput()

	fmt.Println("Please enter the floor number:")
	floor, errFloor := strconv.Atoi(util.GetCommandInput())

	fmt.Println("Please enter the building name:")
	building := util.GetCommandInput()

	fmt.Println("Please enter the location:")
	location := util.GetCommandInput()

	fmt.Println("Please enter the capacity:")
	capacity, errCapacity := strconv.Atoi(util.GetCommandInput())

	fmt.Println("Is the room out of service? (true/false):")
	isOutOfServiceStr := util.GetCommandInput()
	isOutOfService := isOutOfServiceStr == "true"

	room := &model.Room{
		RoomName:           roomName,
		RoomType:           roomType,
		Description:        description,
		Floor:              floor,
		Building:           building,
		Location:           location,
		Capacity:           capacity,
		IsRoomOutOfService: isOutOfService,
		Instrument:         nil,
		Supply:             nil,
	}

	fmt.Println("-------- Room details --------")
	fmt.Println("Room Name:", room.RoomName)
	fmt.Println("Room Type:", room.RoomType)
	fmt.Println("Description:", room.Description)
	fmt.Println("Floor:", room.Floor)
	fmt.Println("Building:", room.Building)
	fmt.Println("Location:", room.Location)
	fmt.Println("Capacity:", room.Capacity)
	fmt.Println("Is Out Of Service:", room.IsRoomOutOfService)
	fmt.Println("Instruments:", room.Instrument)
	fmt.Println("Supplies:", room.Supply)

	err := handler.controller.Insert(*room)
	if err != nil {
		return err
	}

	if errFloor != nil || errCapacity != nil {
		fmt.Println("Failed to create Room", errFloor, errCapacity)
		util.PressEnterToContinue()
	}
	fmt.Println("Room created successfully")
	util.PressEnterToContinue()

	return nil
}
