// MEP-1013
package model

import (
	"ModEd/core"
	"fmt"
	"strings"
    "github.com/go-playground/validator/v10"
)

type Booking struct {
	core.BaseModel
	TimeTableID uint      `gorm:"type:integer;not null" json:"time_table_id" csv:"time_table_id" validate:"-"`
	TimeTable   TimeTable `gorm:"foreignKey:TimeTableID;references:ID" json:"time_table" csv:"time_table" validate:"-"`
	UserID      uint      `gorm:"type:integer;not null" json:"user_id" csv:"user_id" validate:"required"`
	UserRole    Role      `gorm:"type:text;not null" json:"user_role" csv:"user_role" validate:"required"`
	EventName   string    `gorm:"type:text;not null" json:"event_name" csv:"event_name" validate:"required"`
}

func (booking Booking) Validate() error {
    validate := validator.New()
    if err := validate.Struct(booking); err != nil {
        return err
    }
    return nil
}

func (booking Booking) ToString() string {
    return fmt.Sprintf("==================================================================\n"+
        " BOOKING #%-43d \n"+
        "------------------------------------------------------------------\n"+
        " Created:                      %-25s \n"+
        " Updated:                      %-25s \n"+
        "------------------------------------------------------------------\n"+
        " Event Name:                   %-25s \n"+
        " User ID:                      %-25d \n"+
        " User Role:                    %-25s \n"+
        "------------------------------------------------------------------\n"+
        " Room:                         %-25s \n"+
        " Building:                     %-25s \n"+
        " Floor:                        %-25d \n"+
        "------------------------------------------------------------------\n"+
        " Start:                        %-25s \n"+
        " End:                          %-25s \n"+
        " Booking Type:                 %-25s \n"+
        "==================================================================",
        booking.ID,
        booking.CreatedAt.Format("2006-01-02 15:04:05"),
        booking.UpdatedAt.Format("2006-01-02 15:04:05"),
        booking.EventName,
        booking.UserID,
        booking.UserRole,
        booking.TimeTable.Room.RoomName,
        booking.TimeTable.Room.Building,
        booking.TimeTable.Room.Floor,
        booking.TimeTable.StartDate.Format("2006-01-02 15:04"),
        booking.TimeTable.EndDate.Format("2006-01-02 15:04"),
        booking.TimeTable.BookingType)
}

func (booking Booking) ToShortString() string {
    return fmt.Sprintf("| %-3d | %-10s | %-8s | %-10s | %-7s | %-20s |",
        booking.ID,
        booking.TimeTable.Room.RoomName,
        fmt.Sprintf("%d/%s",
            booking.TimeTable.Room.Floor,
            booking.TimeTable.Room.Building),
        booking.TimeTable.StartDate.Format("15:04"),
        string(booking.UserRole),
        truncateString(booking.EventName, 20))
}

func truncateString(str string, length int) string {
    if len(str) <= length {
        return str + strings.Repeat(" ", length-len(str))
    }
    return str[:length-3] + "..."
}