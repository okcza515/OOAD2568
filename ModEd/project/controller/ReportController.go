package controller

import (
	"ModEd/project/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IReportController interface {
	ListAllReports() ([]model.Report, error)
	RetrieveReport(id uuid.UUID) (*model.Report, error)
	InsertReport(report *model.Report) error
	UpdateReport(report *model.Report) error
	DeleteReport(id uuid.UUID) error
}

type ReportController struct {
	db *gorm.DB
}

func NewReportController(db *gorm.DB) IReportController {
	return &ReportController{db: db}
}

func (c *ReportController) ListAllReports() ([]model.Report, error) {
	var reports []model.Report
	err := c.db.Find(&reports).Error
	return reports, err
}

func (c *ReportController) RetrieveReport(id uuid.UUID) (*model.Report, error) {
	var report model.Report
	if err := c.db.Where("report_id = ?", id).First(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (c *ReportController) InsertReport(report *model.Report) error {
	return c.db.Create(report).Error
}

func (c *ReportController) UpdateReport(report *model.Report) error {
	return c.db.Save(report).Error
}

func (c *ReportController) DeleteReport(id uuid.UUID) error {
	return c.db.Where("report_id = ?", id).Delete(&model.Report{}).Error
}
