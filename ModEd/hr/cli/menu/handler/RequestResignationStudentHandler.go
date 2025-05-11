package handler

import "ModEd/hr/controller"

type RequestStudentResignationStrategy struct {
	resignStudentController *controller.ResignationStudentHRController
}

func NewRequestStudentResignationStrategy(resignStudentCtrl *controller.ResignationStudentHRController) *RequestStudentResignationStrategy {
	return &RequestStudentResignationStrategy{resignStudentController: resignStudentCtrl}
}

func (handler RequestStudentResignationStrategy) Execute() error {
	return handleResignationRequest("student", handler.resignStudentController, nil)
}