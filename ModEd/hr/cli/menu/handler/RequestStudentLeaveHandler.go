package handler

import "ModEd/hr/controller"

type RequestStudentLeaveStrategy struct {
	leaveStudentController *controller.LeaveStudentHRController
}

func NewRequestStudentLeaveStrategy(leaveStudentCtrl *controller.LeaveStudentHRController) *RequestStudentLeaveStrategy {
	return &RequestStudentLeaveStrategy{leaveStudentController: leaveStudentCtrl}
}

func (handler RequestStudentLeaveStrategy) Execute() error {
	return handleLeaveRequest("student", handler.leaveStudentController, nil)
}
