// MEP-1013
package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
)

type RetrieveInstrumentManagementByRoomIdStrategy struct {
    controller interface {
        RetrieveByRoomId(roomID uint) (*[]model.InstrumentManagement, error)
    }
}

func NewGetInstrumentManagementByRoomIdStrategy(
	controller interface {
		RetrieveByRoomId(roomID uint) (*[]model.InstrumentManagement, error)
	},
) *RetrieveInstrumentManagementByRoomIdStrategy {
	return &RetrieveInstrumentManagementByRoomIdStrategy{controller: controller}
}

func (handler RetrieveInstrumentManagementByRoomIdStrategy) Execute() error {
    fmt.Println("Enter Room ID to search instruments:")
    input := util.GetCommandInput()

    // Convert input to uint
    var roomID uint
    _, err := fmt.Sscan(input, &roomID)
    if err != nil {
        return fmt.Errorf("invalid room ID: %v", err)
    }

    // Get instruments by room ID
    instruments, err := handler.controller.RetrieveByRoomId(roomID)
    if err != nil {
        return fmt.Errorf("failed to retrieve instruments: %v", err)
    }

    // Display results
    fmt.Printf("\n=== Instruments in Room %d ===\n", roomID)
    if len(*instruments) == 0 {
        fmt.Println("No instruments found in this room")
    } else {
        for _, instrument := range *instruments {
            fmt.Printf("ID: %d | Label: %s | Room ID: %d\n",
                instrument.GetID(),
                instrument.InstrumentLabel,
                instrument.RoomID)
        }
    }

    util.PressEnterToContinue()
    return nil
}