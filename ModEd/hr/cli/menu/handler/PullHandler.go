package handler

import "fmt"

type ControllerPullFunc func() error

type PullHandler struct{ controllerFunc ControllerPullFunc }

func NewPullHandlerStrategy(controllerFunc ControllerPullFunc) *PullHandler {
	return &PullHandler{controllerFunc: controllerFunc}
}

func (handler PullHandler) Execute() error {
	if err := handler.controllerFunc(); err != nil {
		return fmt.Errorf("failed to pull record: %w", err)
	}

	fmt.Println("Pulled successfully into table.")
	return nil
}
