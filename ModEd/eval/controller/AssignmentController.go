// MEP-1006

package controller

import (
	"ModEd/core"
	"ModEd/eval/model"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	// "time"

	"gorm.io/gorm"
)

type AssignmentController struct {
	db   *gorm.DB
	core *core.BaseController[*model.Assignment]
}

type AssignmentControllerInterface interface {
	CreateAssignment(assignment *model.Assignment) (assignmentId uint, err error)
	GetAssignment(assignmentId uint, preload ...string) (assignment *model.Assignment, err error)
	GetAssignments(preload ...string) (assignments []*model.Assignment, err error)
	GetAssignmentsByClass(classId uint, preload ...string) (assignments []*model.Assignment, err error)
	GetAssignmentsByInstructor(instructorCode string, preload ...string) (assignments []*model.Assignment, err error)
	UpdateAssignment(updatedAssignment *model.Assignment) (*model.Assignment, error)
	DeleteAssignment(assignmentId uint) (assignment *model.Assignment, err error)
	UpdateAssignmentStatus(assignmentID uint, newStatus model.AssignmentStatus) error
	SaveAssignmentsToCSV() error
}

func NewAssignmentController(db *gorm.DB) *AssignmentController {
	return &AssignmentController{
		db:   db,
		core: core.NewBaseController[*model.Assignment](db),
	}
}

func (c *AssignmentController) CreateAssignment(assignment *model.Assignment) (assignmentId uint, err error) {
	// หา assignmentId ล่าสุดเพื่อกำหนดค่าใหม่ที่ไม่ซ้ำกัน
	var maxAssignmentId sql.NullInt64
	err = c.db.Model(&model.Assignment{}).Select("max(assignment_id)").Scan(&maxAssignmentId).Error
	if err != nil {
		return 0, fmt.Errorf("failed to get max assignment ID: %v", err)
	}

	// กำหนดค่า assignmentId ใหม่เป็นค่าสูงสุดที่มีอยู่ + 1
	// If maxAssignmentId is NULL (no records yet), use 1 as the first ID
	if maxAssignmentId.Valid {
		assignment.AssignmentId = uint(maxAssignmentId.Int64) + 1
	} else {
		assignment.AssignmentId = 1
	}

	fmt.Printf("Creating new assignment with ID: %d\n", assignment.AssignmentId)

	if err := c.core.Insert(assignment); err != nil {
		return 0, err
	}

	// Save to CSV after creating in database
	if err := c.SaveAssignmentsToCSV(); err != nil {
		fmt.Printf("Warning: Failed to save to CSV: %v\n", err)
	}

	return assignment.AssignmentId, nil
}

func (c *AssignmentController) GetAssignment(assignmentId uint, preload ...string) (assignment *model.Assignment, err error) {
	assignment, err = c.core.RetrieveByCondition(map[string]interface{}{"assignment_id": assignmentId}, preload...)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

func (c *AssignmentController) GetAssignments(preload ...string) (assignments []*model.Assignment, err error) {
	assignments, err = c.core.List(nil, preload...)
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

func (c *AssignmentController) GetAssignmentsByClass(classId uint, preload ...string) (assignments []*model.Assignment, err error) {
	condition := map[string]interface{}{"class_id": classId}
	assignments, err = c.core.List(condition, preload...)
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

func (c *AssignmentController) GetAssignmentsByInstructor(instructorCode string, preload ...string) (assignments []*model.Assignment, err error) {
	condition := map[string]interface{}{"instructor_code": instructorCode}
	assignments, err = c.core.List(condition, preload...)
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

func (c *AssignmentController) UpdateAssignment(updatedAssignment *model.Assignment) (assignment *model.Assignment, err error) {
	assignment, err = c.core.RetrieveByCondition(map[string]interface{}{"assignment_id": updatedAssignment.AssignmentId})
	if err != nil {
		return nil, err
	}

	assignment.Title = updatedAssignment.Title
	assignment.Description = updatedAssignment.Description
	assignment.PublishDate = updatedAssignment.PublishDate
	assignment.DueDate = updatedAssignment.DueDate
	assignment.Status = updatedAssignment.Status
	assignment.ClassId = updatedAssignment.ClassId
	assignment.InstructorCode = updatedAssignment.InstructorCode

	if err := c.core.UpdateByCondition(map[string]interface{}{"assignment_id": updatedAssignment.AssignmentId}, assignment); err != nil {
		return nil, err
	}

	// Save to CSV after updating in database
	if err := c.SaveAssignmentsToCSV(); err != nil {
		fmt.Printf("Warning: Failed to save to CSV: %v\n", err)
	}

	return assignment, nil
}

func (c *AssignmentController) DeleteAssignment(assignmentId uint) (assignment *model.Assignment, err error) {
	assignment, err = c.core.RetrieveByCondition(map[string]interface{}{"assignment_id": assignmentId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"assignment_id": assignmentId}); err != nil {
		return nil, err
	}

	// Save to CSV after deleting from database
	if err := c.SaveAssignmentsToCSV(); err != nil {
		fmt.Printf("Warning: Failed to save to CSV: %v\n", err)
	}

	return assignment, nil
}

func (c *AssignmentController) UpdateAssignmentStatus(assignmentID uint, newStatus model.AssignmentStatus) error {
	assignment, err := c.GetAssignment(assignmentID)
	if err != nil {
		return err
	}

	if assignment.State == nil {
		// Initialize state based on current status if it's nil
		switch assignment.Status {
		case model.StatusDraft:
			assignment.State = &DraftState{}
		case model.StatusPublished:
			assignment.State = &PublishedState{}
		case model.StatusClosed:
			assignment.State = &ClosedState{}
		default:
			assignment.State = &DraftState{}
		}
	}

	if err := assignment.State.HandleStatusChange(assignment, newStatus); err != nil {
		return err
	}

	// Save the updated assignment state to the database
	if err := c.db.Save(assignment).Error; err != nil {
		return fmt.Errorf("failed to save updated assignment status: %v", err)
	}

	// Save to CSV after updating status
	if err := c.SaveAssignmentsToCSV(); err != nil {
		fmt.Printf("Warning: Failed to save to CSV: %v\n", err)
	}

	return nil
}

func (c *AssignmentController) SaveAssignmentsToCSV() error {
	assignments, err := c.GetAssignments()
	if err != nil {
		return fmt.Errorf("failed to retrieve assignments: %v", err)
	}

	file, err := os.Create("../../data/quiz/Assignment.csv")
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{
		"assignment_id",
		"title",
		"description",
		"publish_date",
		"due_date",
		"status",
		"class_id",
		"instructor_code",
	}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header to CSV: %v", err)
	}

	// Write assignments
	for _, assignment := range assignments {
		record := []string{
			strconv.FormatUint(uint64(assignment.AssignmentId), 10),
			assignment.Title,
			assignment.Description,
			assignment.PublishDate.Format("2006-01-02"),
			assignment.DueDate.Format("2006-01-02"),
			string(assignment.Status),
			strconv.FormatUint(uint64(assignment.ClassId), 10),
			assignment.InstructorCode,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write assignment to CSV: %v", err)
		}
	}

	return nil
}
