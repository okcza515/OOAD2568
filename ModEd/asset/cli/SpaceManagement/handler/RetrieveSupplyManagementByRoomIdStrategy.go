// MEP-1013
package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
)

type RetrieveSupplyManagementByRoomIdStrategy struct {
    controller interface {
        RetrieveByRoomId(roomID uint) (*[]model.SupplyManagement, error)
    }
}

func NewGetSupplyManagementByRoomIdStrategy(
	controller interface {
		RetrieveByRoomId(roomID uint) (*[]model.SupplyManagement, error)
	},
) *RetrieveSupplyManagementByRoomIdStrategy {
	return &RetrieveSupplyManagementByRoomIdStrategy{controller: controller}
}

func (handler RetrieveSupplyManagementByRoomIdStrategy) Execute() error {
    fmt.Println("Enter Room ID to search supply managements:")
    input := util.GetCommandInput()

    var roomID uint
    _, err := fmt.Sscan(input, &roomID)
    if err != nil {
        return fmt.Errorf("invalid room ID: %v", err)
    }

    supplyManagements, err := handler.controller.RetrieveByRoomId(roomID)
    if err != nil {
        return fmt.Errorf("failed to retrieve supply managements: %v", err)
    }

    fmt.Printf("\n=== Supply managements in Room %d ===\n", roomID)
    if len(*supplyManagements) == 0 {
        fmt.Println("No supply managements found in this room")
    } else {
        for _, supplyManagement := range *supplyManagements {
            fmt.Printf("ID: %d | Label: %s | Room ID: %d\n",
                supplyManagement.GetID(),
                supplyManagement.SupplyLabel,
                supplyManagement.RoomID)
        }
    }

    util.PressEnterToContinue()
    return nil
}