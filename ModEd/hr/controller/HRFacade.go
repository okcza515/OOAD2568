package controller

import (
	commonModel "ModEd/common/model"
	"ModEd/hr/model"

	"gorm.io/gorm"
)

type HRFacade struct {
	Student               *StudentHRController
	Instructor            *InstructorHRController
	ResignationStudent    *ResignationStudentHRController
	ResignationInstructor *ResignationInstructorHRController
	Leave                 *LeaveHRController
}

func NewHRFacade(db *gorm.DB) *HRFacade {
	return &HRFacade{
		Student:               CreateStudentHRController(db),
		Instructor:            CreateInstructorHRController(db),
		ResignationStudent:    CreateResignationStudentHRController(db),
		ResignationInstructor: CreateResignationInstructorHRController(db),
		Leave:                 CreateLeaveHRController(db),
	}
}

// Student-related facade methods

func (f *HRFacade) GetAllStudents() ([]*model.StudentInfo, error) {
	return f.Student.GetAll()
}

func (f *HRFacade) GetStudentById(sid string) (*model.StudentInfo, error) {
	return f.Student.GetById(sid)
}

func (f *HRFacade) InsertStudent(info *model.StudentInfo) error {
	return f.Student.Insert(info)
}

func (f *HRFacade) UpdateStudent(info *model.StudentInfo) error {
	return f.Student.Update(info)
}

func (f *HRFacade) DeleteStudent(sid string) error {
	return f.Student.Delete(sid)
}

func (f *HRFacade) UpdateStudentStatus(sid string, status commonModel.StudentStatus) error {
	return f.Student.UpdateStatus(sid, status)
}

func (f *HRFacade) UpsertStudent(info *model.StudentInfo) error {
	return f.Student.Upsert(info)
}

// Instructor-related facade methods

func (f *HRFacade) GetAllInstructors() ([]*model.InstructorInfo, error) {
	return f.Instructor.GetAll()
}

func (f *HRFacade) GetInstructorById(id string) (*model.InstructorInfo, error) {
	return f.Instructor.GetById(id)
}

func (f *HRFacade) InsertInstructor(info *model.InstructorInfo) error {
	return f.Instructor.Insert(info)
}

func (f *HRFacade) UpdateInstructor(info *model.InstructorInfo) error {
	return f.Instructor.Update(info)
}

func (f *HRFacade) UpsertInstructor(info *model.InstructorInfo) error {
	return f.Instructor.Upsert(info)
}

func (f *HRFacade) DeleteInstructor(id string) error {
	return f.Instructor.Delete(id)
}

// Resignation-related facade methods

func (f *HRFacade) SubmitResignationStudentRequest(info *model.RequestResignationStudent) error {
	return f.ResignationStudent.Insert(info)
}

func (f *HRFacade) SubmitResignationInstructorRequest(info *model.RequestResignationInstructor) error {
	return f.ResignationInstructor.Insert(info)
}

// Leave-related facade methods
func (f *HRFacade) SubmitLeaveRequest(info *model.RequestLeave) error {
	return f.Leave.Insert(info)
}

func (f *HRFacade) UpdateResignationStudentStatus(id string, status string, reason string) error {
	req, err := f.ResignationStudent.GetByStudentID(id)
	if err != nil {
		return err
	}
	req.Status = status
	if status == "Rejected" && reason != "" {
		req.Reason = reason
	}
	return f.ResignationStudent.Update(req)
}
