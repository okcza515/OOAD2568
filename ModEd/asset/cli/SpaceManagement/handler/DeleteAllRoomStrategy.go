// MEP-1013
package handler

import (
	"ModEd/asset/util"
	"fmt"
)

type DeleteAllRoomStrategy struct {
	controller interface {
		DeleteAll() error
	}
}

func NewDeleteAllRoomStrategy(controller interface {
	DeleteAll() error
}) *DeleteAllRoomStrategy {
	return &DeleteAllRoomStrategy{controller: controller}
}

func (handler DeleteAllRoomStrategy) Execute() error {
	err := handler.controller.DeleteAll()
	if err != nil {
		return err
	}
	fmt.Println("All rooms deleted successfully!")
	util.PressEnterToContinue()
	return nil
}
