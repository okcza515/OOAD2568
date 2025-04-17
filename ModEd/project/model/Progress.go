package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Progress struct {
	gorm.Model
	AssignmentId uint   `gorm:"not null;index"`
	Name         string `gorm:"not null"`
	IsCompleted  bool   `gorm:"not null"`
}

func (p *Progress) GetID() uint {
	return p.ID
}
func (p *Progress) FromCSV(csvData string) error {
	return nil
}
func (p *Progress) ToCSVRow() string {
	return fmt.Sprintf("%d,%d,%s,%t", p.ID, p.AssignmentId, p.Name, p.IsCompleted)
}
func (p *Progress) FromJSON(jsonData string) error {
	return nil
}
func (p *Progress) ToJSON() string {
	return ""
}
func (p *Progress) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}
func (p *Progress) ToString() string {
	return fmt.Sprintf("Progress[ID=%d, AssignmentId=%d, Name=%s, IsCompleted=%t]", p.ID, p.AssignmentId, p.Name, p.IsCompleted)
}
