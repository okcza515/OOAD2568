package controller

import (
    "errors"
    "fmt"
    "os"

    commonController "ModEd/common/controller"
    commonModel "ModEd/common/model"
    mController "ModEd/hr/controller/Migration"
    sController "ModEd/hr/controller/Student"
    hrModel "ModEd/hr/model"
    "ModEd/hr/util"

    "gorm.io/gorm"
)

type HRController struct {
    db                *gorm.DB
    StudentController *sController.StudentHRController
}

func NewHRController(db *gorm.DB) *HRController {
    return &HRController{
        db:                db,
        StudentController: sController.CreateStudentHRController(db),
    }
}

// ListStudents returns all student information.
func (c *HRController) ListStudents() ([]hrModel.StudentInfo, error) {
    return c.StudentController.GetAll()
}

// UpdateStudent updates student information with provided values.
func (c *HRController) UpdateStudent(sid string, firstName, lastName, gender, citizenID, phoneNumber, email string) error {
    // Get existing student info
    studentInfo, err := c.StudentController.GetById(sid)
    if err != nil {
        return fmt.Errorf("error retrieving student with ID %s: %v", sid, err)
    }

    // Update fields if values provided
    if firstName != "" {
        studentInfo.FirstName = firstName
    }
    if lastName != "" {
        studentInfo.LastName = lastName
    }
    if gender != "" {
        studentInfo.Gender = gender
    }
    if citizenID != "" {
        studentInfo.CitizenID = citizenID
    }
    if phoneNumber != "" {
        studentInfo.PhoneNumber = phoneNumber
    }
    if email != "" {
        studentInfo.Email = email
    }

    return c.StudentController.Update(studentInfo)
}

// CreateStudent creates a new student with the provided information.
func (c *HRController) CreateStudent(studentID, firstName, lastName, gender, citizenID, phoneNumber string) error {
    newStudent := hrModel.StudentInfo{
        Student: commonModel.Student{
            StudentCode: studentID,
            FirstName:   firstName,
            LastName:    lastName,
        },
        Gender:      gender,
        CitizenID:   citizenID,
        PhoneNumber: phoneNumber,
    }

    return c.StudentController.Insert(&newStudent)
}

// ImportStudents imports students from a file and creates or updates their info.
func (c *HRController) ImportStudents(filePath string) error {
    // Check if file exists
    if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
        return fmt.Errorf("file %s does not exist", filePath)
    }

    // Create mapper for the file
    hrMapper, err := util.CreateMapper[hrModel.StudentInfo](filePath)
    if err != nil {
        return fmt.Errorf("failed to create HR mapper: %v", err)
    }

    // Get records from file
    hrRecords := hrMapper.Map()
    commonStudentController := commonController.CreateStudentController(c.db)

    // Process each record
    for _, hrRec := range hrRecords {
        commonStudent, err := commonStudentController.GetByStudentId(hrRec.StudentCode)
        if err != nil {
            fmt.Printf("Failed to retrieve student %s from common data: %v\n", hrRec.StudentCode, err)
            continue
        }

        newStudent := hrModel.StudentInfo{
            Student: *commonStudent,
            Gender:  hrRec.Gender,
        }

        if err := c.StudentController.Upsert(&newStudent); err != nil {
            fmt.Printf("Failed to upsert student %s: %v\n", newStudent.StudentCode, err)
            continue
        }
    }
    return nil
}

// SynchronizeStudents synchronizes student data from common model to HR.
func (c *HRController) SynchronizeStudents() error {
    return mController.SynchronizeStudents(c.db)
}