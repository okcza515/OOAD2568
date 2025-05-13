// MEP-1013
package handler

import(
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core"
	"gorm.io/gorm"
	"fmt"
)

type UpdateSupplyManagementStrategy struct{
	controller interface {
        UpdateById(payload *model.SupplyManagement) error
    }
}

func NewUpdateSupplyManagementStrategy(
    controller interface {
        UpdateById(payload *model.SupplyManagement) error
    },
) *UpdateSupplyManagementStrategy {
    return &UpdateSupplyManagementStrategy{
        controller: controller,
    }
}

func (handler UpdateSupplyManagementStrategy) Execute() error {
    fmt.Println("=== Update Supply Management ===")

    fmt.Print("Enter ID to update: ")
    var id uint
    _, err := fmt.Sscan(util.GetCommandInput(), &id)
    if err != nil {
        return fmt.Errorf("invalid ID format: %v", err)
    }

    fmt.Print("Enter new Room ID: ")
    var roomID uint
    _, err = fmt.Sscan(util.GetCommandInput(), &roomID)
    if err != nil {
        return fmt.Errorf("invalid Room ID format: %v", err)
    }

    fmt.Print("Enter new Supply ID: ")
    var supplyID uint
    _, err = fmt.Sscan(util.GetCommandInput(), &supplyID)
    if err != nil {
        return fmt.Errorf("invalid Supply ID format: %v", err)
    }

    supplyManagement := &model.SupplyManagement{
        BaseModel: core.BaseModel{
            Model: gorm.Model{ID: id},
        },
        RoomID:    roomID,
        SupplyID:  supplyID,
    }

    if err := supplyManagement.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

    if err := handler.controller.UpdateById(supplyManagement); err != nil {
        return fmt.Errorf("failed to update supply management: %v", err)
    }

    fmt.Printf("\nSuccessfully updated supply management with ID %d!\n", id)
    util.PressEnterToContinue()
    return nil
}