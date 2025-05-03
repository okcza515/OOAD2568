// MEP-1013
package menu

import (
	controller "ModEd/asset/controller"
	spaceManagementHandler "ModEd/asset/handler"
	model "ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"

	"gorm.io/gorm"
)

type RoomMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func (menu *RoomMenuState) Render() {
	fmt.Println()
	fmt.Println("===== Room Management =====")
	fmt.Println("1. Add new Room")
	fmt.Println("2. List all Rooms")
	fmt.Println("3. Get detail of a Room")
	fmt.Println("4. Update a Room")
	fmt.Println("5. Delete a Room")
	fmt.Println("6. Delete all Rooms")
	// fmt.Println("7. Seed Rooms Data")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("===========================")
}
func (menu *RoomMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error handling user input", err)
	}
	if input == "back" {

		util.PressEnterToContinue()

	}
	return err
}

func NewRoomMenuState(db *gorm.DB, manager *cli.CLIMenuStateManager, spaceManagementMenu *SpaceManagementState) *RoomMenuState {
	// Check if db is nil
	if db == nil {
		fmt.Println("Error: Database connection is nil")
		return &RoomMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	// Get controller instance and check if it's nil
	controllerFacade := controller.GetSpaceManagementInstance(db)
	if controllerFacade == nil {
		fmt.Println("Error: Space Management Controller Facade is nil")
		return &RoomMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	// Access Room controller safely
	controllerInstance := controllerFacade.Room
	if controllerInstance == nil {
		fmt.Println("Error: Room controller is nil")
		return &RoomMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	handlerContext := handler.NewHandlerContext()

	//Handler is not working yet!
	// insertHandler := handler.NewInsertHandlerStrategy[model.Room](controllerInstance)
	listHandler := handler.NewListHandlerStrategy[model.Room](controllerInstance)
	getHandler := handler.NewRetrieveByIDHandlerStrategy[model.Room](controllerInstance)
	// updateHandler := handler.NewUpdateHandlerStrategy[model.Room](controllerInstance)
	deleteHandler := handler.NewDeleteHandlerStrategy[model.Room](controllerInstance)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, spaceManagementMenu)

	//DIY
	insertHandler := spaceManagementHandler.NewAddRoomHandlerStrategy(controllerInstance)
	updateHandler := spaceManagementHandler.NewUpdateRoomHandlerStrategy(controllerInstance)
	deleteAllHandler := spaceManagementHandler.NewDeleteAllRoomStrategy(controllerInstance)

	handlerContext.AddHandler("1", "Add New Room", insertHandler)
	handlerContext.AddHandler("2", "List all Rooms", listHandler)
	handlerContext.AddHandler("3", "Get full detail of a Room", getHandler)
	handlerContext.AddHandler("4", "Update a Room", updateHandler)
	handlerContext.AddHandler("5", "Delete a Room", deleteHandler)
	handlerContext.AddHandler("6", "Delete all Rooms", deleteAllHandler)
	// handlerContext.AddHandler("7", "Seed Rooms Data", nil)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &RoomMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

// func RoomHandler(facade *controller.SpaceManagementControllerFacade) {
// 	inputBuffer := ""
// 	for inputBuffer != "back" {
// 		util.ClearScreen()
// 		util.PrintSpaceManagementBanner()
// 		fmt.Println()
// 		printRoomMenuOption()
// 		inputBuffer = util.GetCommandInput()
// 		switch inputBuffer {
// 		case "1":
// 			fmt.Println("Add new Room")
// 			fmt.Println("Please enter the name of the new Room:")
// 			roomName := util.GetCommandInput()

// 			fmt.Println("Please enter the room type: (Lecture / Laboratory / Office)")
// 			roomTypeStr := util.GetCommandInput()
// 			var roomType model.RoomTypeEnum
// 			switch roomTypeStr {
// 			case "Lecture":
// 				roomType = "ROOM_LECTURE_ROOM"
// 			case "Laboratory":
// 				roomType = "ROOM_LAB_ROOM"
// 			case "Office":
// 				roomType = "ROOM_MEETING_ROOM"
// 			default:
// 				fmt.Println("Invalid room type. Using default type.")
// 				roomType = "ROOM_LECTURE_ROOM"
// 			}

// 			fmt.Println("Please enter the description:")
// 			description := util.GetCommandInput()

// 			fmt.Println("Please enter the floor number:")
// 			floor, errFloor := strconv.Atoi(util.GetCommandInput())

// 			fmt.Println("Please enter the building name:")
// 			building := util.GetCommandInput()

// 			fmt.Println("Please enter the location:")
// 			location := util.GetCommandInput()

// 			fmt.Println("Please enter the capacity:")
// 			capacity, errCapacity := strconv.Atoi(util.GetCommandInput())

// 			fmt.Println("Is the room out of service? (true/false):")
// 			isOutOfServiceStr := util.GetCommandInput()
// 			isOutOfService := isOutOfServiceStr == "true"

// 			room := &model.Room{
// 				RoomName:           roomName,
// 				RoomType:           roomType,
// 				Description:        description,
// 				Floor:              floor,
// 				Building:           building,
// 				Location:           location,
// 				Capacity:           capacity,
// 				IsRoomOutOfService: isOutOfService,
// 				Instrument:         nil,
// 				Supply:             nil,
// 			}
// 			err := facade.Room.CreateRoom(room)
// 			if errFloor != nil || err != nil || errCapacity != nil {
// 				fmt.Println("Failed to create Room", err, errFloor, errCapacity)
// 				util.PressEnterToContinue()
// 				break
// 			}
// 			fmt.Println("Room created successfully")
// 			util.PressEnterToContinue()
// 		case "2":
// 			fmt.Println("List all Room")
// 			data, err := facade.Room.GetAll()
// 			if err != nil {
// 				panic(err)
// 			}
// 			if len(*data) == 0 {
// 				fmt.Println("No Room available")
// 				util.PressEnterToContinue()
// 				break
// 			}
// 			for _, room := range *data {
// 				fmt.Println(room)
// 			}
// 			util.PressEnterToContinue()

// 		case "3":
// 			fmt.Println("Get detail of a Room")
// 			fmt.Println("Please enter the ID of the Room:")
// 			roomIdStr := util.GetCommandInput()
// 			roomId, err := strconv.Atoi(roomIdStr)
// 			if err != nil {
// 				fmt.Println("Invalid ID")
// 				util.PressEnterToContinue()
// 				break
// 			}
// 			data, err := facade.Room.GetById(uint(roomId))
// 			if err != nil {
// 				fmt.Println("Failed to get Room detail", err)
// 				util.PressEnterToContinue()
// 				break
// 			}
// 			fmt.Println("Room detail: ", data)
// 			util.PressEnterToContinue()
// 		case "4":
// 			fmt.Println("Update a Room")
// 			fmt.Println("Please enter the ID of the Room:")
// 			roomIdStr := util.GetCommandInput()
// 			roomId, err := strconv.Atoi(roomIdStr)
// 			if err != nil {
// 				fmt.Println("Invalid ID")
// 				util.PressEnterToContinue()
// 				break
// 			}
// 			fmt.Println("Please enter the new name of the Room:")
// 			roomName := util.GetCommandInput()
// 			fmt.Println("Please enter the new room type: (Lecture / Laboratory / Office)")
// 			roomTypeStr := util.GetCommandInput()
// 			var roomType model.RoomTypeEnum
// 			switch roomTypeStr {
// 			case "Lecture":
// 				roomType = "ROOM_LECTURE_ROOM"
// 			case "Laboratory":
// 				roomType = "ROOM_LAB_ROOM"
// 			case "Office":
// 				roomType = "ROOM_MEETING_ROOM"
// 			default:
// 				fmt.Println("Invalid room type. Using default type.")
// 				roomType = "ROOM_LECTURE_ROOM"
// 			}
// 			fmt.Println("Please enter the new description:")
// 			description := util.GetCommandInput()
// 			fmt.Println("Please enter the new floor number:")
// 			floor, errFloor := strconv.Atoi(util.GetCommandInput())
// 			fmt.Println("Please enter the new building name:")
// 			building := util.GetCommandInput()
// 			fmt.Println("Please enter the new location:")
// 			location := util.GetCommandInput()
// 			fmt.Println("Please enter the new capacity:")
// 			capacity, errCapacity := strconv.Atoi(util.GetCommandInput())
// 			fmt.Println("Is the room out of service? (true/false):")
// 			isOutOfServiceStr := util.GetCommandInput()
// 			isOutOfService := isOutOfServiceStr == "true"
// 			room := &model.Room{
// 				RoomName:           roomName,
// 				RoomType:           roomType,
// 				Description:        description,
// 				Floor:              floor,
// 				Building:           building,
// 				Location:           location,
// 				Capacity:           capacity,
// 				IsRoomOutOfService: isOutOfService,
// 				Instrument:         nil,
// 				Supply:             nil,
// 			}
// 			err = facade.Room.UpdateRoom(uint(roomId), room)
// 			if err != nil || errFloor != nil || errCapacity != nil {
// 				fmt.Println("Failed to update Room", err, errFloor, errCapacity)
// 				util.PressEnterToContinue()
// 				break
// 			}
// 			fmt.Println("Room updated successfully")
// 			util.PressEnterToContinue()
// 		case "5":
// 			fmt.Println("Delete a Room")
// 			fmt.Println("Please enter the ID of the Room:")
// 			roomIdStr := util.GetCommandInput()
// 			roomId, err := strconv.Atoi(roomIdStr)
// 			if err != nil {
// 				fmt.Println("Invalid ID / Room not found")
// 				util.PressEnterToContinue()
// 				break
// 			}
// 			err = facade.Room.DeleteRoom(uint(roomId))
// 			if err != nil {
// 				fmt.Println("Failed to delete Room, Room not found")
// 				util.PressEnterToContinue()
// 			} else {
// 				fmt.Println("Room deleted successfully")
// 				util.PressEnterToContinue()
// 			}
// 		case "6":
// 			fmt.Println()
// 			err := facade.Room.DeleteAllRooms()
// 			if err != nil {
// 				fmt.Println("Failed to delete all rooms")
// 			} else {
// 				fmt.Println("Deleted all rooms succeed!")
// 			}
// 			util.PressEnterToContinue()
// 		case "7":
// 			fmt.Println("Seed Rooms Data")
// 			data, err := facade.Room.SeedRoomsDatabase("data/asset/Room.json")
// 			if err != nil {
// 				fmt.Println("Failed to seed Room data", err)
// 				util.PressEnterToContinue()
// 			} else {
// 				fmt.Println("Room data seeded successfully")
// 				for _, room := range data {
// 					fmt.Println(room)
// 				}
// 				util.PressEnterToContinue()
// 			}
// 		case "back":
// 		default:
// 			fmt.Println("Invalid Command")
// 			util.PressEnterToContinue()
// 		}
// 		util.ClearScreen()
// 	}
// 	util.ClearScreen()
// }
