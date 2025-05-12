// MEP-1013
package handler

import(
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
)

type InsertSupplyManagementStrategy struct {
    controller interface {
         Insert(payload *model.SupplyManagement) error
    }
}

func NewInsertSupplyManagementStrategy(
    controller interface {
        Insert(payload *model.SupplyManagement) error
    },
) *InsertSupplyManagementStrategy {
    return &InsertSupplyManagementStrategy{
        controller: controller,
    }
}

func (handler InsertSupplyManagementStrategy) Execute() error {
    fmt.Println("=== Create New Supply Management ===")
    
    // Get Room ID
    fmt.Print("Enter Room ID: ")
    var roomID uint
    _, err := fmt.Sscan(util.GetCommandInput(), &roomID)
    if err != nil {
        return fmt.Errorf("invalid room ID: %v", err)
    }

    // Get Supply ID
    fmt.Print("Enter Supply ID: ")
    var supplyID uint
    _, err = fmt.Sscan(util.GetCommandInput(), &supplyID)
    if err != nil {
        return fmt.Errorf("invalid supply ID: %v", err)
    }
    
    // Create supply management record
    supplyManagement := &model.SupplyManagement{
        RoomID:    roomID,
        SupplyID:  supplyID,
    }

    if err := supplyManagement.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

    if err := handler.controller.Insert(supplyManagement); err != nil {
        return fmt.Errorf("failed to create supply management: %v", err)
    }

    fmt.Println("\nSupply Management is created successfully!")
    util.PressEnterToContinue()
    return nil
}