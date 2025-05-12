package controller

import (
	"ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type InternshipMentorController struct {
	Connector *gorm.DB
}

func NewInternshipMentorController(connector *gorm.DB) *InternshipMentorController {
	return &InternshipMentorController{
		Connector: connector,
	}
}

func (mc *InternshipMentorController) Create(mentor *model.InternshipMentor) error {
	if err := mc.Connector.Create(mentor).Error; err != nil {
		return fmt.Errorf("failed to create mentor: %w", err)
	}
	return nil
}

func (mc *InternshipMentorController) RetrieveByID(id uint) (*model.InternshipMentor, error) {
	var mentor model.InternshipMentor
	if err := mc.Connector.First(&mentor, id).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve mentor with ID %d: %w", id, err)
	}
	return &mentor, nil
}

func (mc *InternshipMentorController) Update(mentor *model.InternshipMentor) error {
	if err := mc.Connector.Save(mentor).Error; err != nil {
		return fmt.Errorf("failed to update mentor with ID %d: %w", mentor.ID, err)
	}
	return nil
}

func (mc *InternshipMentorController) DeleteByID(id uint) error {
	if err := mc.Connector.Delete(&model.InternshipMentor{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete mentor with ID %d: %w", id, err)
	}
	return nil
}

func (mc *InternshipMentorController) ListAll() ([]model.InternshipMentor, error) {
	var mentors []model.InternshipMentor
	if err := mc.Connector.Find(&mentors).Error; err != nil {
		return nil, fmt.Errorf("failed to list mentors: %w", err)
	}
	return mentors, nil
}
