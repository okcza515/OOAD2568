package handler

import "ModEd/hr/controller"

type RequestInstructorLeaveStrategy struct {
	leaveInstructorController *controller.LeaveInstructorHRController
}

func NewRequestInstructorLeaveStrategy(leaveInstructorCtrl *controller.LeaveInstructorHRController) *RequestInstructorLeaveStrategy {
	return &RequestInstructorLeaveStrategy{leaveInstructorController: leaveInstructorCtrl}
}

func (handler RequestInstructorLeaveStrategy) Execute() error {
	return handleLeaveRequest("instructor", nil, handler.leaveInstructorController)
}