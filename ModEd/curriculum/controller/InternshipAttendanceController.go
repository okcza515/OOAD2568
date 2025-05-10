package controller

import (
	"ModEd/curriculum/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type InternshipAttendanceController struct {
	Connector *gorm.DB
}

func NewInternshipAttendanceController(connector *gorm.DB) *InternshipAttendanceController {
	return &InternshipAttendanceController{
		Connector: connector,
	}
}

func (c *InternshipAttendanceController) CreateAttendance(attendance *model.Attendance) error {
	if err := c.Connector.Create(attendance).Error; err != nil {
		return fmt.Errorf("failed to create attendance record: %w", err)
	}
	return nil
}

func (c *InternshipAttendanceController) RetrieveAttendanceByID(id uint) (*model.Attendance, error) {
	var attendance model.Attendance
	if err := c.Connector.First(&attendance, id).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve attendance record with ID %d: %w", id, err)
	}
	return &attendance, nil
}

func (c *InternshipAttendanceController) UpdateAttendance(attendance *model.Attendance) error {
	if err := c.Connector.Save(attendance).Error; err != nil {
		return fmt.Errorf("failed to update attendance record with ID %d: %w", attendance.ID, err)
	}
	return nil
}

func (c *InternshipAttendanceController) DeleteAttendanceByID(id uint) error {
	if err := c.Connector.Delete(&model.Attendance{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete attendance record with ID %d: %w", id, err)
	}
	return nil
}

func (c *InternshipAttendanceController) ListAllAttendances() ([]model.Attendance, error) {
	var attendances []model.Attendance
	if err := c.Connector.Find(&attendances).Error; err != nil {
		return nil, fmt.Errorf("failed to list attendance records: %w", err)
	}
	return attendances, nil
}

func (c *InternshipAttendanceController) ListAttendancesByStudentID(studentID uint) ([]model.Attendance, error) {
	var attendances []model.Attendance
	if err := c.Connector.Where("student_info_id = ?", studentID).Find(&attendances).Error; err != nil {
		return nil, fmt.Errorf("failed to list attendance records for student ID %d: %w", studentID, err)
	}
	return attendances, nil
}

func (c *InternshipAttendanceController) MarkAttendance(studentID uint, date time.Time, checkInTime, checkOutTime time.Time, checkInStatus bool, assignedWork string) error {
	attendance := &model.Attendance{
		Date:          date,
		CheckInTime:   checkInTime,
		CheckOutTime:  checkOutTime,
		CheckInStatus: checkInStatus,
		AssingWork:    assignedWork,
		StudentInfoID: studentID,
	}

	if err := c.Connector.Create(attendance).Error; err != nil {
		return fmt.Errorf("failed to mark attendance for student ID %d: %w", studentID, err)
	}
	return nil
}
