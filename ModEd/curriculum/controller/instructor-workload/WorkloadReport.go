package controller

import (
	"gorm.io/gorm"
)

type WorkloadReportController struct {
	Connector *gorm.DB
}

func CreateWorkloadReportController(connector *gorm.DB) *WorkloadReportController {
	return &WorkloadReportController{
		Connector: connector,
	}
}
