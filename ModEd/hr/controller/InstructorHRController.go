package controller

import (
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type InstructorHRController struct {
	db *gorm.DB
}

func CreateInstructorHRController(db *gorm.DB) *InstructorHRController {
	db.AutoMigrate(&model.InstructorInfo{})
	return &InstructorHRController{db: db}
}

func (c *InstructorHRController) getAll() ([]*model.InstructorInfo, error) {
	var infos []*model.InstructorInfo
	err := c.db.Find(&infos).Error
	return infos, err
}

func (c *InstructorHRController) getById(id string) (*model.InstructorInfo, error) {
	var instructorInfo model.InstructorInfo
	if err := c.db.Where("instructor_id = ?", id).First(&instructorInfo).Error; err != nil {
		return nil, err
	}
	return &instructorInfo, nil
}

func (c *InstructorHRController) insert(info *model.InstructorInfo) error {
	return c.db.Create(info).Error
}

func (c *InstructorHRController) update(info *model.InstructorInfo) error {
	return c.db.Model(&model.InstructorInfo{}).
		Where("id = ?", info.ID).
		Updates(info).Error
}

func (c *InstructorHRController) delete(id string) error {
	return c.db.Where("instructor_id = ?", id).Delete(&model.InstructorInfo{}).Error
}

func (c *InstructorHRController) GetAllInstructors(tx *gorm.DB) ([]*model.InstructorInfo, error) {
	instructors, err := c.getAll()
	if err != nil {
		return nil, fmt.Errorf("error retrieving instructors: %v", err)
	}
	return instructors, nil
}

func (c *InstructorHRController) UpdateInstructorInfo(tx *gorm.DB, instructorID, field, value string) error {
	tm := &util.TransactionManager{DB: tx}
	return tm.Execute(func(tx *gorm.DB) error {
		instructorInfo, err := c.getById(instructorID)
		if err != nil {
			return fmt.Errorf("error retrieving instructor with ID %s: %v", instructorID, err)
		}

		switch strings.ToLower(field) {
		case "position", "academicposition", "academic_position":
			parsedPos, err := model.ParseAcademicPosition(value)
			if err != nil {
				return fmt.Errorf("invalid academic position: %v", err)
			}
			instructorInfo.AcademicPosition = parsedPos
		case "department":
			// Uncomment and adjust if InstructorInfo has a Department field.
			// instructorInfo.Department = value
		default:
			return fmt.Errorf("unknown field for instructor update: %s", field)
		}

		if err := c.update(instructorInfo); err != nil {
			return fmt.Errorf("error updating instructor: %v", err)
		}
		fmt.Println("Instructor updated successfully!")
		return nil
	})
}

func (c *InstructorHRController) ImportInstructors(tx *gorm.DB, instructors []*model.InstructorInfo) error {
	for _, instructor := range instructors {
		if instructor.ID == 0 || instructor.FirstName == "" {
			return fmt.Errorf("invalid instructor data: %+v", instructor)
		}

		if err := c.insert(instructor); err != nil {
			return fmt.Errorf("failed to insert instructor %d: %v", instructor.ID, err)
		}
	}
	return nil
}
