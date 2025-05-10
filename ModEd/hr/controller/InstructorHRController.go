package controller

import (
	commonController "ModEd/common/controller"
	commonModel "ModEd/common/model"
	"ModEd/core"
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"
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

		instructorController := NewInstructorHRController(tx)
		if err := instructorController.MigrateInstructorRecords(); err != nil {
			return fmt.Errorf("failed to migrate instructor to HR module: %w", err)
		}

		hrInstructor := model.NewInstructorInfo(*commonInstructor, gender, citizenID, phoneNumber, salary, academicPosition, departmentPosition)
		if updateErr := instructorController.update(hrInstructor); updateErr != nil {
			return fmt.Errorf("failed to update instructor HR info: %w", updateErr)
		}

		return nil
	})
	return err
}

func (c *InstructorHRController) ImportInstructors(filePath string) error {
	tm := &util.TransactionManager{DB: c.db}
	err := tm.Execute(func(tx *gorm.DB) error {
		hrMapper, err := core.CreateMapper[model.InstructorInfo](filePath)
		if err != nil {
			return fmt.Errorf("failed to create mapper: %w", err)
		}
		hrRecords := hrMapper.Deserialize()
		hrRecordsMap := make(map[string]model.InstructorInfo)
		for _, record := range hrRecords {
			if _, exists := hrRecordsMap[record.InstructorCode]; exists {
				return fmt.Errorf("duplicate instructor code found: %s", record.InstructorCode)
			}
			if record != nil {
				hrRecordsMap[record.InstructorCode] = *record
			} else {
				continue
			}
		}

		instructorController := NewInstructorHRController(tx)
		for instructorCode, record := range hrRecordsMap {
			instructorInfo, err := instructorController.getById(instructorCode)
			if err != nil {
				return fmt.Errorf("error retrieving instructor with ID %s: %v", instructorCode, err)
			}

			importInstructor := model.NewUpdatedInstructorInfo(
				instructorInfo,
				instructorInfo.FirstName,
				instructorInfo.LastName,
				instructorInfo.Email,
				record.Gender,
				record.CitizenID,
				record.PhoneNumber,
				record.AcademicPosition,
				record.DepartmentPosition,
			)
			if err := instructorController.update(importInstructor); err != nil {
				return fmt.Errorf("error updating instructor with ID %s: %v", instructorCode, err)
			}

		}
		return nil
	})
	return err
}

func (c *InstructorHRController) GetAllInstructors() ([]*model.InstructorInfo, error) {
	instructors, err := c.getAll()
	if err != nil {
		return nil, fmt.Errorf("error retrieving instructors: %v", err)
	}
	return instructors, nil
}

func (c *InstructorHRController) UpdateInstructorInfo(instructorID string, firstName string, lastName string, email string,
	gender string, citizenID string, phoneNumber string, academicPosition string, departmentPosition string) error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {
		instructorController := NewInstructorHRController(tx)

		academicPos, err := model.ParseAcademicPosition(academicPosition)
		if err != nil {
			return fmt.Errorf("failed to parse academic position: %w", err)
		}

		departmentPos, err := model.ParseDepartmentPosition(departmentPosition)
		if err != nil {
			return fmt.Errorf("failed to parse department position: %w", err)
		}

		instructorInfo, err := instructorController.getById(instructorID)
		if err != nil {
			return fmt.Errorf("error retrieving instructor with ID %s: %v", instructorID, err)
		}

		instructorData := map[string]any{
			"FirstName": firstName,
			"LastName":  lastName,
			"Email":     email,
		}

		commonInstructorController := commonController.NewInstructorController(tx)
		if err := commonInstructorController.Update(instructorID, instructorData); err != nil {
			return fmt.Errorf("failed to update common instructor data: %v", err)
		}

		if err := instructorController.MigrateInstructorRecords(); err != nil {
			return fmt.Errorf("failed to migrate instructor to HR module: %v", err)
		}

		instructorHRData := model.NewUpdatedInstructorInfo(instructorInfo, firstName, lastName, email, gender, citizenID, phoneNumber, academicPos, departmentPos)
		if err := instructorController.update(instructorHRData); err != nil {
			return fmt.Errorf("failed to update instructor HR info: %v", err)
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
