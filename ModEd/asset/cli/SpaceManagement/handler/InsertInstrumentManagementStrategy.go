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
    
    // Get Room ID
    fmt.Print("Enter Room ID: ")
    var roomID uint
    _, err := fmt.Sscan(util.GetCommandInput(), &roomID)
    if err != nil {
        return fmt.Errorf("invalid room ID: %v", err)
    }

    // Get Instrument ID
    fmt.Print("Enter Instrument ID: ")
    var instrumentID uint
    _, err = fmt.Sscan(util.GetCommandInput(), &instrumentID)
    if err != nil {
        return fmt.Errorf("invalid instrument ID: %v", err)
    }

    // Get Borrow ID (required)
    fmt.Print("Enter Borrow ID: ")
    var borrowID uint
    _, err = fmt.Sscan(util.GetCommandInput(), &borrowID)
    if err != nil {
        return fmt.Errorf("invalid borrow ID: %v", err)
    }
    
    // Create instrument management record
    instrumentManagement := &model.InstrumentManagement{
        RoomID:         roomID,
        InstrumentID:   instrumentID,
        BorrowUserID:   borrowID,
    }

    if err := handler.controller.Insert(instrumentManagement); err != nil {
        return fmt.Errorf("failed to create instrument management: %v", err)
    }

    fmt.Println("\nInstrument Management is created successfully!")
    util.PressEnterToContinue()
    return nil
}