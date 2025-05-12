package handlerr

// MEP-1012 Asset
import (
	"ModEd/core"
	"fmt"
	"strconv"
	"time"
)

type ReturnHandlerStrategy[T core.RecordInterface] struct {
	controller interface {
		RetrieveByID(id uint, preloads ...string) (T, error)
		UpdateByID(data T) error
		List(condition map[string]interface{}, preloads ...string) ([]T, error)
	}
}

func NewReturnHandlerStrategy[T core.RecordInterface](
	controller interface {
		RetrieveByID(id uint, preloads ...string) (T, error)
		UpdateByID(data T) error
		List(condition map[string]interface{}, preloads ...string) ([]T, error)
	},
) *ReturnHandlerStrategy[T] {
	return &ReturnHandlerStrategy[T]{controller: controller}
}

func (handler ReturnHandlerStrategy[T]) Execute() error {
	if handler.controller == nil {
		return fmt.Errorf("controller is nil")
	}
	records, err := handler.controller.List(nil)
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return fmt.Errorf("No records found.")
	}

	result := fmt.Sprintf("Total %v record(s)\n\n", len(records))
	for _, record := range records {
		result += fmt.Sprintf("%v\n", record.ToString())
	}
	result += "\nEnter ID to return: "
	fmt.Print(result)

	var inputBuffer string
	fmt.Scanln(&inputBuffer)

	idNum, err := strconv.Atoi(inputBuffer)
	if err != nil {
		return fmt.Errorf("invalid ID format. Please enter a valid number")
	}

	record, err := handler.controller.RetrieveByID(uint(idNum))
	if err != nil {
		return fmt.Errorf("record not found with ID %d", idNum)
	}

	now := time.Now()

	if r, ok := any(&record).(interface{ SetReturnDate(*time.Time) }); ok {
		r.SetReturnDate(&now)
	} else {
		return fmt.Errorf("record does not implement SetReturnDate")
	}

	err = handler.controller.UpdateByID(record)
	if err != nil {
		return fmt.Errorf("error updating return status: %v", err)
	}

	fmt.Printf("Borrow record ID %d successfully returned at %s", record.GetID(), now.Format("2006-01-02 15:04"))
	return nil
}
