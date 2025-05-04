// MEP-1013
package handler

import(
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
)

type InsertInstrumentManagementStrategy struct {
    controller interface {
         Insert(payload *model.InstrumentManagement) error
    }
}

func NewInsertInstrumentManagementStrategy(
    controller interface {
        Insert(payload *model.InstrumentManagement) error
    },
) *InsertInstrumentManagementStrategy {
    return &InsertInstrumentManagementStrategy{
        controller: controller,
    }
}

func (handler InsertInstrumentManagementStrategy) Execute() error {
    fmt.Println("=== Create New Instrument Management ===")
    
    fmt.Print("Enter Room ID: ")
    var roomID uint
    _, err := fmt.Sscan(util.GetCommandInput(), &roomID)
    if err != nil {
        return fmt.Errorf("invalid room ID: %v", err)
    }

    fmt.Print("Enter Instrument Label: ")
    label := util.GetCommandInput()
    
    instrument := &model.InstrumentManagement{
        RoomID: roomID,
        InstrumentLabel: label,
    }

    if err := handler.controller.Insert(instrument); err != nil {
        return fmt.Errorf("failed to create instrument: %v", err)
    }

    fmt.Println("\nInstrument created successfully!")
    util.PressEnterToContinue()
    return nil
}