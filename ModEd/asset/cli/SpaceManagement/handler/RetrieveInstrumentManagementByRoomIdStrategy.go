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
    fmt.Println("Enter Room ID to search instrument managements:")
    input := util.GetCommandInput()

    var roomID uint
    _, err := fmt.Sscan(input, &roomID)
    if err != nil {
        return fmt.Errorf("invalid room ID: %v", err)
    }

    instrumentManagements, err := handler.controller.RetrieveByRoomId(roomID)
    if err != nil {
        return fmt.Errorf("failed to retrieve instrument managements: %v", err)
    }

    fmt.Printf("\n=== Instrument managements in Room %d ===\n", roomID)
    if len(*instrumentManagements) == 0 {
        fmt.Println("No instrument managements found in this room")
    } else {
        for _, instrumentManagement := range *instrumentManagements {
            fmt.Printf("ID: %d | Label: %s | Room ID: %d\n",
                instrumentManagement.GetID(),
                instrumentManagement.InstrumentLabel,
                instrumentManagement.RoomID)
        }
    }

    util.PressEnterToContinue()
    return nil
}