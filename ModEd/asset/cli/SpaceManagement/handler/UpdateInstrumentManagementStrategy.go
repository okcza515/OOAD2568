// MEP-1013
package handler

import(
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core"
	"gorm.io/gorm"
	"fmt"
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

    // Get new instrument label
    fmt.Print("Enter new instrument management Label: ")
    label := util.GetCommandInput()

    // Create update payload
    instrument := &model.InstrumentManagement{
        BaseModel: core.BaseModel{
        Model: gorm.Model{ID: id},
    },
        RoomID: roomID,
        InstrumentLabel: label,
    }

    // Perform update
    if err := handler.controller.UpdateById(instrument); err != nil {
        return fmt.Errorf("failed to update instrument management: %v", err)
    }

    fmt.Printf("\nSuccessfully updated instrument management with ID %d!\n", id)
    util.PressEnterToContinue()
    return nil
}