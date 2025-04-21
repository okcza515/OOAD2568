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
	LeaveStudent                 *LeaveStudentHRController
	Raise                 *RaiseHRController
	LeaveInstructor 			*LeaveInstructorHRController	
}

func NewHRFacade(db *gorm.DB) *HRFacade {
	return &HRFacade{
		Student:               createStudentHRController(db),
		Instructor:            createInstructorHRController(db),
		ResignationStudent:    createResignationStudentHRController(db),
		ResignationInstructor: createResignationInstructorHRController(db),
		LeaveStudent:                 createLeaveStudentHRController(db),
		LeaveInstructor: 			createLeaveInstructorHRController(db),
		Raise:                 createRaiseHRController(db),
	}
}

// Student-related facade methods

func (f *HRFacade) GetAllStudents() ([]*model.StudentInfo, error) {
	return f.Student.getAll()
}

func (f *HRFacade) GetStudentById(sid string) (*model.StudentInfo, error) {
	return f.Student.getById(sid)
}

func (f *HRFacade) InsertStudent(info *model.StudentInfo) error {
	return f.Student.insert(info)
}

func (f *HRFacade) UpdateStudent(info *model.StudentInfo) error {
	return f.Student.update(info)
}

func (f *HRFacade) DeleteStudent(sid string) error {
	return f.Student.delete(sid)
}

func (f *HRFacade) UpdateStudentStatus(sid string, status commonModel.StudentStatus) error {
	return f.Student.updateStatus(sid, status)
}

func (f *HRFacade) UpsertStudent(info *model.StudentInfo) error {
	return f.Student.upsert(info)
}

// Instructor-related facade methods

func (f *HRFacade) GetAllInstructors() ([]*model.InstructorInfo, error) {
	return f.Instructor.getAll()
}

func (f *HRFacade) GetInstructorById(id string) (*model.InstructorInfo, error) {
	return f.Instructor.getById(id)
}

func (f *HRFacade) InsertInstructor(info *model.InstructorInfo) error {
	return f.Instructor.insert(info)
}

func (f *HRFacade) UpdateInstructor(info *model.InstructorInfo) error {
	return f.Instructor.update(info)
}

func (f *HRFacade) UpsertInstructor(info *model.InstructorInfo) error {
	return f.Instructor.upsert(info)
}

func (f *HRFacade) DeleteInstructor(id string) error {
	return f.Instructor.delete(id)
}

// Resignation-related facade methods

func (f *HRFacade) SubmitResignationStudentRequest(info *model.RequestResignationStudent) error {
	return f.ResignationStudent.insert(info)
}

func (f *HRFacade) SubmitResignationInstructorRequest(info *model.RequestResignationInstructor) error {
	return f.ResignationInstructor.insert(info)
}

// Leave-related facade methods
func (f *HRFacade) SubmitLeaveStudentRequest(info *model.RequestLeaveStudent) error {
	return f.LeaveStudent.insert(info)
}
func (f *HRFacade) SubmitLeaveInstructorRequest(info *model.RequestLeaveInstructor) error {
	return f.LeaveInstructor.insert(info)
}

func (f *HRFacade) UpdateResignationStudentStatus(id string, status string, reason string) error {
	req, err := f.ResignationStudent.getByStudentID(id)
	if err != nil {
		return err
	}
	req.Status = status
	if status == "Rejected" && reason != "" {
		req.Reason = reason
	}
	return f.ResignationStudent.update(req)
}

//Raise-related facade methods

func (f *HRFacade) SubmitRaiseInstructorRequest(request *model.RequestRaise) error {
	return f.Raise.insert(request)
}

func (f *HRFacade) ApproveRaise(id uint) error {
	return f.Raise.updateStatus(id, "Approved")
}

func (f *HRFacade) RejectRaise(id uint) error {
	return f.Raise.updateStatus(id, "Rejected")
}
