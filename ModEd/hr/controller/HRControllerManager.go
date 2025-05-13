package controller

import (
	"ModEd/core/migration"
	"errors"
)

type HRControllerManager struct {
	StudentCtrl          *StudentHRController
	InstructorCtrl       *InstructorHRController
	LeaveStudentCtrl     *LeaveStudentHRController
	LeaveInstructorCtrl  *LeaveInstructorHRController
	ResignStudentCtrl    *ResignationStudentHRController
	ResignInstructorCtrl *ResignationInstructorHRController
	RaiseCtrl            *RaiseHRController
}

var hrInstance *HRControllerManager

func GetHRInstance() *HRControllerManager {
	if hrInstance == nil {
		instance, err := newHRControllerManager()
		if err != nil {
			panic(err)
		}
		hrInstance = instance
	}
	return hrInstance
}

func newHRControllerManager() (*HRControllerManager, error) {
	manager := &HRControllerManager{}
	db := migration.GetInstance().DB
	if db == nil {
		return nil, errors.New("err: db not initialized")
	}

	manager.StudentCtrl = NewStudentHRController(db)
	manager.InstructorCtrl = NewInstructorHRController(db)
	manager.LeaveStudentCtrl = NewLeaveStudentHRController(db)
	manager.LeaveInstructorCtrl = NewLeaveInstructorHRController(db)
	manager.ResignStudentCtrl = NewResignationStudentHRController(db)
	manager.ResignInstructorCtrl = NewResignationInstructorHRController(db)
	manager.RaiseCtrl = NewRaiseHRController(db)

	return manager, nil
}
