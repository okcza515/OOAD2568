// MEP-1013
package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core"
	"fmt"

	"gorm.io/gorm"
)

type UpdateInstrumentManagementStrategy struct{
	controller interface {
        UpdateById(payload *model.InstrumentManagement) error
    }
}

func NewUpdateInstrumentManagementStrategy(
    controller interface {
        UpdateById(payload *model.InstrumentManagement) error
    },
) *UpdateInstrumentManagementStrategy {
    return &UpdateInstrumentManagementStrategy{
        controller: controller,
    }
}

func (handler UpdateInstrumentManagementStrategy) Execute() error {
    fmt.Println("=== Update Instrument Management ===")

    // Get ID to update
    fmt.Print("Enter ID to update: ")
    var id uint
    _, err := fmt.Sscan(util.GetCommandInput(), &id)
    if err != nil {
        return fmt.Errorf("invalid ID format: %v", err)
    }

    // Get new room ID
    fmt.Print("Enter new Room ID: ")
    var roomID uint
    _, err = fmt.Sscan(util.GetCommandInput(), &roomID)
    if err != nil {
        return fmt.Errorf("invalid Room ID format: %v", err)
    }

    // Get new instrument ID
    fmt.Print("Enter new Instrument ID: ")
    var instrumentID uint
    _, err = fmt.Sscan(util.GetCommandInput(), &instrumentID)
    if err != nil {
        return fmt.Errorf("invalid Instrument ID format: %v", err)
    }

    // Get new borrow ID
    fmt.Print("Enter new Borrow ID: ")
    var borrowID uint
    _, err = fmt.Sscan(util.GetCommandInput(), &borrowID)
    if err != nil {
        return fmt.Errorf("invalid Borrow ID format: %v", err)
    }

    // Create update payload
    instrumentManagement := &model.InstrumentManagement{
        BaseModel: core.BaseModel{
            Model: gorm.Model{ID: id},
        },
        RoomID:         roomID,
        InstrumentID:   instrumentID,
        BorrowUserID:   borrowID,
    }

    if err := instrumentManagement.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

    // Perform update
    if err := handler.controller.UpdateById(instrumentManagement); err != nil {
        return fmt.Errorf("failed to update instrument management: %v", err)
    }

    fmt.Printf("\nSuccessfully updated instrument management with ID %d!\n", id)
    util.PressEnterToContinue()
    return nil
}