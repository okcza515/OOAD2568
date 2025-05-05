// MEP-1008
package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"

	"time"

	"gorm.io/gorm"
)

type CoursePlanInterface interface {
	CreateCoursePlan(CoursePlan model.CoursePlan) error
	RetrieveByID(id uint, preloads ...string) (*model.CoursePlan, error)
	UpdateByID(data model.CoursePlan) error
	DeleteByID(id uint) error
	ListAllCoursePlans() ([]model.CoursePlan, error)
	ListUpcomingPlan() ([]model.CoursePlan, error)
}

type CoursePlanController struct {
	*core.BaseController[*model.CoursePlan]
	Connector *gorm.DB
}

func NewCoursePlanController(db *gorm.DB) *CoursePlanController {
	return &CoursePlanController{
		BaseController: core.NewBaseController[*model.CoursePlan](db),
		Connector:      db,
	}
}

func (src *CoursePlanController) CreateCoursePlan(coursePlan *model.CoursePlan) (uint, error) {
	if err := src.BaseController.Insert(coursePlan); err != nil {
		return 0, err
	}
	return coursePlan.ID, nil
}

func (src *CoursePlanController) ListAllCoursePlans() ([]model.CoursePlan, error) {
	var coursePlans []model.CoursePlan
	result := src.Connector.Find(&coursePlans)
	return coursePlans, result.Error
}

func (src *CoursePlanController) ListUpcomingPlan() ([]model.CoursePlan, error) {
	var coursePlans []model.CoursePlan
	result := src.Connector.Where("date >= ?", time.Now()).Find(&coursePlans)
	return coursePlans, result.Error
}
