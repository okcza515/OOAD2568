// MEP-1008
package controller

import (
	model "ModEd/curriculum/model"

	"time"

	"gorm.io/gorm"
)

type CoursePlanController struct {
	DB *gorm.DB
}

func (src *CoursePlanController) CreateCoursePlan(CoursePlan model.CoursePlan) error {
	return src.DB.Create(&CoursePlan).Error
}

func (src *CoursePlanController) UpdateCoursePlan(course_id uint, body *model.CoursePlan) error {
	body.ID = course_id
	result := src.DB.Updates(body)
	return result.Error
}

func (src *CoursePlanController) DeleteCoursePlan(course_id uint) error {
	result := src.DB.Model(&model.CoursePlan{}).Where("ID = ?", course_id).Update("deleted_at", nil)
	return result.Error
}

func (src *CoursePlanController) ListAllCoursePlans() ([]model.CoursePlan, error) {
	var coursePlans []model.CoursePlan
	result := src.DB.Find(&coursePlans)
	return coursePlans, result.Error
}

func (src *CoursePlanController) ListPlanByCourseID(courseID uint) ([]model.CoursePlan, error) {
	var coursePlans []model.CoursePlan
	result := src.DB.Where("course_id = ?", courseID).Find(&coursePlans)
	return coursePlans, result.Error
}

func (src *CoursePlanController) ListUpcomingPlan() ([]model.CoursePlan, error) {
	var coursePlans []model.CoursePlan
	result := src.DB.Where("date >= ?", time.Now()).Find(&coursePlans)
	return coursePlans, result.Error
}
