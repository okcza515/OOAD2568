//MEP-1009 Student Internship
package controller

import (
    "ModEd/curriculum/model"
    "gorm.io/gorm"
)

type ReportController struct {
    *BaseScoreController[model.InternshipReport]
}

func NewReportController(connector *gorm.DB) *ReportController {
    return &ReportController{
        BaseScoreController: &BaseScoreController[model.InternshipReport]{Connector: connector},
    }
}

func (rc *ReportController) UpdateReportScore(studentID string, score int) error {
    scoreFields := map[string]interface{}{
        "ReportScore": score,
    }

    return rc.UpdateScore(studentID, scoreFields, func(db *gorm.DB, studentID string) (uint, error) {
        var application model.InternshipApplication
        if err := db.Where("student_code = ?", studentID).Last(&application).Error; err != nil {
            return 0, err
        }
        return application.InternshipReportId, nil
    })
}