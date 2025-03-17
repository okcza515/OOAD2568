package controller

import (
	"ModEd/common/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	Connector *gorm.DB
}

func CreateDepartmentController(connector *gorm.DB) *Department {
    department := Department{Connector: connector}
	connector.AutoMigrate(&model.Department{})
	return &department
}

func (department *Department) GetAll() ([]*model.Department, error) {
	departments := []*model.Department{}
	result := department.Connector.Find(&departments)
	return departments, result.Error
}

func (department *Department) GetByDepartmentId(departmentId uuid.UUID) (*model.Department, error) {
	d := &model.Department{}
	result := department.Connector.Where("department_id = ?", departmentId).First(d)
	return d, result.Error
}

func (department *Department) SetBudget(departmentName string, newBudget int) error {
    return department.Connector.Model(&model.Department{}).
        Where("name = ?", departmentName).
        Update("budget", newBudget).Error
}

//use for both decrement and increment
func (department *Department) UpdateBudget(departmentName string, updateAmount int) error {
	if (updateAmount >= 0) {
		return department.Connector.Model(&model.Department{}).
        	Where("name = ?", departmentName).
        	Update("budget", gorm.Expr("budget + ?", updateAmount)).Error
	} else { // ensure the budget won't go below 0
		return department.Connector.Model(&model.Department{}).
			Where("name = ?", departmentName).
			Where("budget + ? >= 0", updateAmount).
        	Update("budget", gorm.Expr("budget + ?", updateAmount)).Error
	}
	
}
