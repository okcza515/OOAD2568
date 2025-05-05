package handler

import (
	"ModEd/curriculum/controller"
	"fmt"
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
	workloadReportMenu.Add("Generate Performance Report", GeneratePerformanceReportHandler{db: w.db})
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
	builder.SetHeader("Daily Workload").WithClasses().WithMeetings().WithProjects().SetDateRange(&startToday, &endToday).Generate()
}

type WeeklyWorkloadHandler struct {
	db *gorm.DB
}

func (w WeeklyWorkloadHandler) Execute() {
	builder := controller.NewWorkloadReportBuilder(w.db, 1)
	startWeek := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday()))
	endWeek := startWeek.Add(7 * 24 * time.Hour).Add(-time.Nanosecond)
	builder.SetHeader("Weekly Workload").WithClasses().WithMeetings().WithProjects().SetDateRange(&startWeek, &endWeek).Generate()
}

type MonthlyWorkloadHandler struct {
	db *gorm.DB
}

func (m MonthlyWorkloadHandler) Execute() {
	builder := controller.NewWorkloadReportBuilder(m.db, 1)
	startMonth := time.Now().Truncate(24*time.Hour).AddDate(0, -int(time.Now().Month()), 0)
	endMonth := startMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)
	builder.SetHeader("Monthly Workload").WithClasses().WithMeetings().WithProjects().SetDateRange(&startMonth, &endMonth).Generate()
}

type GeneratePerformanceReportHandler struct {
	db *gorm.DB
}

func (g GeneratePerformanceReportHandler) Execute() {
	startMonth := time.Now().Truncate(24*time.Hour).AddDate(0, -int(time.Now().Month()), 0)
	endMonth := startMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)
	facade := controller.NewInstructorReportFacade(g.db, 1)
	report, err := facade.GeneratePerformanceReport(1, &startMonth, &endMonth)
	if err != nil {
		fmt.Println("Error generating report:", err)
		return
	}
	fmt.Printf("Instructor ID: %d\n", report.InstructorID)
	fmt.Printf("Total Classes: %d\n", report.TotalClasses)
	fmt.Printf("Total Teaching Hours: %.1f\n", report.TotalTeachingHours)
	fmt.Printf("Total Meetings: %d\n", report.TotalMeetings)
}
