package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
)

type UpdateRoomHandlerStrategy struct {
	controller interface {
		UpdateByID(data model.Room) error
	}
}

func NewUpdateRoomHandlerStrategy(controller interface {
	UpdateByID(data model.Room) error
}) *UpdateRoomHandlerStrategy {
	return &UpdateRoomHandlerStrategy{controller: controller}
}
func (handler UpdateRoomHandlerStrategy) Execute() error {
	fmt.Println("------- Update a Room -------")
	fmt.Println("Please enter the ID of the Room:")
	roomIdStr := util.GetCommandInput()
	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		fmt.Println("Invalid ID")
		util.PressEnterToContinue()
	}
	fmt.Println("Please enter the new name of the Room:")
	roomName := util.GetCommandInput()
	fmt.Println("Please enter the new room type: (Lecture / Laboratory / Office)")
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
	fmt.Println("Please enter the new description:")
	description := util.GetCommandInput()
	fmt.Println("Please enter the new floor number:")
	floor, errFloor := strconv.Atoi(util.GetCommandInput())
	fmt.Println("Please enter the new building name:")
	building := util.GetCommandInput()
	fmt.Println("Please enter the new location:")
	location := util.GetCommandInput()
	fmt.Println("Please enter the new capacity:")
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
	room.ID = uint(roomId)
	err = handler.controller.UpdateByID(*room)
	if err != nil || errFloor != nil || errCapacity != nil {
		fmt.Println("Failed to update Room", err, errFloor, errCapacity)
		util.PressEnterToContinue()
	} else {
		fmt.Println("Room updated successfully")
	}
	util.PressEnterToContinue()

	return nil
}
