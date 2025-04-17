// MEP-1013
package handler

import (
	controller "ModEd/asset/controller/spacemanagement"
	model "ModEd/asset/model/spacemanagement"
	"ModEd/asset/util"
	"fmt"
	"strconv"
)

func printOption() {
	fmt.Println("===== Room Management =====")
	fmt.Println("1. Add new Room")
	fmt.Println("2. List all Rooms")
	fmt.Println("3. Get detail of a Room")
	fmt.Println("4. Update a Room")
	fmt.Println("5. Delete a Room")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("===========================")
}

func RoomHandler(facade *controller.SpaceManagementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintBanner()
		printOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Add new Room")
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
			err := facade.Room.CreateRoom(room)
			if errFloor != nil || err != nil || errCapacity != nil {
				fmt.Println("Failed to create Room", err, errFloor, errCapacity)
				util.PressEnterToContinue()
				break
			}
			fmt.Println("Room created successfully")
			util.PressEnterToContinue()
		case "2":
			fmt.Println("List all Room")
			data, err := facade.Room.GetAll()
			if err != nil {
				panic(err)
			}
			if len(*data) == 0 {
				fmt.Println("No Room available")
				util.PressEnterToContinue()
				break
			}
			for _, room := range *data {
				fmt.Println(room)
			}
			util.PressEnterToContinue()

		case "3":
			fmt.Println("Get detail of a Room")
			fmt.Println("Please enter the ID of the Room:")
			roomIdStr := util.GetCommandInput()
			roomId, err := strconv.Atoi(roomIdStr)
			if err != nil {
				fmt.Println("Invalid ID")
				util.PressEnterToContinue()
				break
			}
			data, err := facade.Room.GetById(uint(roomId))
			if err != nil {
				fmt.Println("Failed to get Room detail", err)
				util.PressEnterToContinue()
				break
			}
			fmt.Println("Room detail: ", data)
			util.PressEnterToContinue()
		case "4":
			fmt.Println("Update a Room")
			fmt.Println("Please enter the ID of the Room:")
			roomIdStr := util.GetCommandInput()
			roomId, err := strconv.Atoi(roomIdStr)
			if err != nil {
				fmt.Println("Invalid ID")
				util.PressEnterToContinue()
				break
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
			err = facade.Room.UpdateRoom(uint(roomId), room)
			if err != nil || errFloor != nil || errCapacity != nil {
				fmt.Println("Failed to update Room", err, errFloor, errCapacity)
				util.PressEnterToContinue()
				break
			}
			fmt.Println("Room updated successfully")
			util.PressEnterToContinue()
		case "5":
			fmt.Println("Delete a Room")
			fmt.Println("Please enter the ID of the Room:")
			roomIdStr := util.GetCommandInput()
			roomId, err := strconv.Atoi(roomIdStr)
			if err != nil {
				fmt.Println("Invalid ID / Room not found")
				util.PressEnterToContinue()
				break
			}
			err = facade.Room.DeleteRoom(uint(roomId))
			if err != nil {
				fmt.Println("Failed to delete Room, Room not found")
				util.PressEnterToContinue()
			} else {
				fmt.Println("Room deleted successfully")
				util.PressEnterToContinue()
			}
		default:
			fmt.Println("Invalid Command")
			util.PressEnterToContinue()
		}

		util.ClearScreen()
	}

	util.ClearScreen()
}
