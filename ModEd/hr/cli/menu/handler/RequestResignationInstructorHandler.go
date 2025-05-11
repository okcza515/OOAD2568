package handler

import "ModEd/hr/controller"

type RequestInstructorResignationStrategy struct {
	resignInstructorController *controller.ResignationInstructorHRController
}

func NewRequestInstructorResignationStrategy(resignInstructorCtrl *controller.ResignationInstructorHRController) *RequestInstructorResignationStrategy {
	return &RequestInstructorResignationStrategy{resignInstructorController: resignInstructorCtrl}
}

func (handler RequestInstructorResignationStrategy) Execute() error {
	return handleResignationRequest("instructor", nil, handler.resignInstructorController)
}
