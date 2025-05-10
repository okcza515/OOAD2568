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

func (handler InsertInstrumentManagementStrategy) Execute() error { //Don't need ID because it is auto-gen by db
    fmt.Println("=== Create New Instrument Management ===")
    
    fmt.Print("Enter Room ID: ")
    var roomID uint
    _, err := fmt.Sscan(util.GetCommandInput(), &roomID)
    if err != nil {
        return fmt.Errorf("invalid room ID: %v", err)
    }

    fmt.Print("Enter Instrument Label: ")
    label := util.GetCommandInput()
    
    instrumentManagement := &model.InstrumentManagement{
        RoomID: roomID,
        InstrumentLabel: label,
    }

    if err := handler.controller.Insert(instrumentManagement); err != nil {
        return fmt.Errorf("failed to create instrument management: %v", err)
    }

    fmt.Println("\nInstrument Management is created successfully!")
    util.PressEnterToContinue()
    return nil
}