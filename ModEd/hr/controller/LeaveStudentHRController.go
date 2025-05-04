package controller

import (
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"

	"gorm.io/gorm"
	
)
type LeaveStudentHRController struct {
	db *gorm.DB
}

func CreateLeaveStudentHRController(db *gorm.DB) *LeaveStudentHRController {
	db.AutoMigrate(&model.RequestLeaveStudent{})
	return &LeaveStudentHRController{db: db}
}

func (c *LeaveStudentHRController) insert(request *model.RequestLeaveStudent) error {
	return c.db.Create(request).Error
}
func (c *LeaveStudentHRController) update(request *model.RequestLeaveStudent) error {
	return c.db.Save(request).Error
}
func (c *LeaveStudentHRController) delete(request *model.RequestLeaveStudent) error {
	return c.db.Delete(request).Error
}
func (c *LeaveStudentHRController) getAll() ([]model.RequestLeaveStudent, error) {
	var requests []model.RequestLeaveStudent
	err := c.db.Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}
func (c *LeaveStudentHRController) getByID(id uint) (*model.RequestLeaveStudent, error) {
	var request model.RequestLeaveStudent
	err := c.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
func (c *LeaveStudentHRController) getByStudentID(studentID string) ([]model.RequestLeaveStudent, error) {
	var requests []model.RequestLeaveStudent
	err := c.db.Where("student_id = ?", studentID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}


func (c *LeaveStudentHRController) SubmitStudentLeaveRequest(studentID, leaveType, reason, leaveDateStr string) error {

	tm := &util.TransactionManager{DB: c.db}

	return tm.Execute(func(tx *gorm.DB) error {
		leaveController := CreateLeaveStudentHRController(tx) // ใช้ transaction ของตัวนี้เลย
		factory := &model.RequestLeaveFactory{}

		req, err := factory.Create("student", studentID, leaveType, reason, leaveDateStr)
		if err != nil {
			return fmt.Errorf("failed to build leave request: %v", err)
		}

		if err := leaveController.insert(req.(*model.RequestLeaveStudent)); err != nil {
			return fmt.Errorf("failed to submit leave request: %v", err)
		}
		return nil
	})
}
