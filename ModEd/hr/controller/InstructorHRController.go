package controller

import (
	commonModel "ModEd/common/model"
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type InstructorHRController struct {
	db *gorm.DB
}

func NewInstructorHRController(db *gorm.DB) *InstructorHRController {
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
	if err := c.db.Where("instructor_code = ?", id).First(&instructorInfo).Error; err != nil {
		return nil, err
	}
	return &instructorInfo, nil
}

func (c *InstructorHRController) insert(info *model.InstructorInfo) error {
	return c.db.Create(info).Error
}

func (c *InstructorHRController) update(info *model.InstructorInfo) error {
	return c.db.Model(&model.InstructorInfo{}).
		Where("instructor_code = ?", info.InstructorCode).
		Updates(info).Error
}

func (c *InstructorHRController) delete(id string) error {
	return c.db.Where("instructor_code = ?", id).Delete(&model.InstructorInfo{}).Error
}

func (c *InstructorHRController) AddInstructor(
	instructorCode string, firstName string, lastName string, email string, startDate string, department string,
	gender string, citizenID string, phoneNumber string, salary int, academicPos string, departmentPos string,
) error {
	tm := &util.TransactionManager{DB: c.db}
	err := tm.Execute(func(tx *gorm.DB) error {
		startDateParsed, err := time.Parse("02-01-2006", startDate)
		if err != nil {
			return fmt.Errorf("failed to parse start date: %w", err)
		}

		academicPosition, err := model.ParseAcademicPosition(academicPos)
		if err != nil {
			return fmt.Errorf("failed to parse academic position: %w", err)
		}

		departmentPosition, err := model.ParseDepartmentPosition(departmentPos)
		if err != nil {
			return fmt.Errorf("failed to parse department position: %w", err)
		}

		commonInstructor := &commonModel.Instructor{
			InstructorCode: instructorCode,
			FirstName:      firstName,
			LastName:       lastName,
			Email:          email,
			StartDate:      &startDateParsed,
			Department:     &department,
		}

		err = commonModel.ManualAddInstructor(tx, commonInstructor)
		if err != nil {
			return fmt.Errorf("failed to add instructor: %w", err)
		}

		// Migrate here !!

		hrInstructor := model.NewInstructorInfo(*commonInstructor, gender, citizenID, phoneNumber, salary, academicPosition, departmentPosition)
		instructorController := NewInstructorHRController(tx)
		if updateErr := instructorController.update(hrInstructor); updateErr != nil {
			return fmt.Errorf("failed to update instructor HR info: %w", updateErr)
		}

		return nil
	})
	return err
}

func (c *InstructorHRController) ImportInstructors(filePath string) error {
	return nil
}

func (c *InstructorHRController) GetAllInstructors() ([]*model.InstructorInfo, error) {
	instructors, err := c.getAll()
	if err != nil {
		return nil, fmt.Errorf("error retrieving instructors: %v", err)
	}
	return instructors, nil
}

func (c *InstructorHRController) UpdateInstructorInfo(instructorID, field, value string) error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {
		instructorController := NewInstructorHRController(tx)
		
		instructorInfo, err := instructorController.getById(instructorID)
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


		if err := instructorController.update(instructorInfo); err != nil {
            return fmt.Errorf("error updating instructor: %v", err)
        }
		fmt.Println("Instructor updated successfully!")
		return nil
	})
}

func (c *InstructorHRController) MigrateInstructorRecords() error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {
		var commonInstructors []commonModel.Instructor
		if err := tx.Find(&commonInstructors).Error; err != nil {
			return fmt.Errorf("failed to retrieve common instructors: %w", err)
		}

		for _, ci := range commonInstructors {
			commonInstructor := &commonModel.Instructor{
				InstructorCode: ci.InstructorCode,
				FirstName:      ci.FirstName,
				LastName:       ci.LastName,
				Email:          ci.Email,
				StartDate:      ci.StartDate,
				Department:     ci.Department,
			}
			instructorInfo := model.NewInstructorInfo(
				*commonInstructor,
				"",
				"",
				"",
				0,
				model.AcademicPosition(0),
				model.DepartmentPosition(0),
			)

			if err := tx.Where("instructor_code = ?", ci.InstructorCode).
				FirstOrCreate(&instructorInfo).Error; err != nil {
				return fmt.Errorf("failed to migrate instructor %s: %w", ci.InstructorCode, err)
			}
		}

		return nil
	})
}
