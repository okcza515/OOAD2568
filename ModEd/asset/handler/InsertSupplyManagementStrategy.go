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

    // Get Supply Label
    fmt.Print("Enter Supply Label: ")
    label := util.GetCommandInput()
    
    // Get Quantity
    fmt.Print("Enter Quantity: ")
    var quantity int
    _, err = fmt.Sscan(util.GetCommandInput(), &quantity)
    if err != nil {
        return fmt.Errorf("invalid quantity: %v", err)
    }
    
    // Create supply management record
    supplyManagement := &model.SupplyManagement{
        RoomID: roomID,
        SupplyLabel: label,
        Quantity: quantity,
    }

    if err := handler.controller.Insert(supplyManagement); err != nil {
        return fmt.Errorf("failed to create supply management: %v", err)
    }

    fmt.Println("\nSupply Management is created successfully!")
    util.PressEnterToContinue()
    return nil
}