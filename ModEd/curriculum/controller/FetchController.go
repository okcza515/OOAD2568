// MEP-1009 Student Internship
package controller

import (
    "ModEd/curriculum/model"
    "fmt"
    "gorm.io/gorm"
)

type Fetcher struct {
    Connector *gorm.DB
}

func NewFetcher(connector *gorm.DB) *Fetcher {
    return &Fetcher{Connector: connector}
}

func (rf *Fetcher) FetchIDByStudentID(studentID string) (uint, error) {
    var application model.InternshipApplication
    if err := rf.Connector.Where("student_code = ?", studentID).Last(&application).Error; err != nil {
        return 0, fmt.Errorf("failed to find application for student '%s': %w", studentID, err)
    }
    return application.InternshipReportId, nil
}