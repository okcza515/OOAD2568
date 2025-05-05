package handler

import (
	"ModEd/curriculum/controller"
	"time"

	"gorm.io/gorm"
)

type WorkloadReportHandler struct {
	db *gorm.DB
}

func NewWorkloadReportHandler(db *gorm.DB) WorkloadReportHandler {
	return WorkloadReportHandler{db: db}
}

func (w WorkloadReportHandler) Execute() {
	workloadReportMenu := NewMenuHandler("Workload Report Menu", true)
	workloadReportMenu.Add("View Today Workload", DailyWorkloadHandler{db: w.db})
	workloadReportMenu.Add("View Weekly Workload", WeeklyWorkloadHandler{db: w.db})
	workloadReportMenu.Add("View Monthly Workload", MonthlyWorkloadHandler{db: w.db})
	workloadReportMenu.Add("View Custom Workload", nil)
	workloadReportMenu.Add("Generate Performance Report", nil)
	workloadReportMenu.SetBackHandler(Back{})
	workloadReportMenu.SetDefaultHandler(UnknownCommand{})
	workloadReportMenu.Execute()
}

type DailyWorkloadHandler struct {
	db *gorm.DB
}

func (d DailyWorkloadHandler) Execute() {
	builder := controller.NewWorkloadReportBuilder(d.db, 1)
	startToday := time.Now().Truncate(24 * time.Hour)
	endToday := startToday.Add(24*time.Hour - time.Nanosecond)
	builder.SetHeader("Daily Workload").WithCoursePlans().WithProjects().SetDateRange(&startToday, &endToday).Generate()
}

type WeeklyWorkloadHandler struct {
	db *gorm.DB
}

func (w WeeklyWorkloadHandler) Execute() {
	builder := controller.NewWorkloadReportBuilder(w.db, 1)
	startWeek := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday()))
	endWeek := startWeek.Add(7 * 24 * time.Hour).Add(-time.Nanosecond)
	builder.SetHeader("Weekly Workload").WithCoursePlans().WithProjects().SetDateRange(&startWeek, &endWeek).Generate()
}

type MonthlyWorkloadHandler struct {
	db *gorm.DB
}

func (m MonthlyWorkloadHandler) Execute() {
	builder := controller.NewWorkloadReportBuilder(m.db, 1)
	startMonth := time.Now().Truncate(24*time.Hour).AddDate(0, -int(time.Now().Month()), 0)
	endMonth := startMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)
	builder.SetHeader("Monthly Workload").WithCoursePlans().WithProjects().SetDateRange(&startMonth, &endMonth).Generate()
}

type CustomWorkloadHandler struct {
	db *gorm.DB
}
