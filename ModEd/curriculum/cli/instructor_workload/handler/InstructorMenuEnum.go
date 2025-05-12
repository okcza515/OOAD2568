package handler

type WorkloadCommand string

const (
	MENU_LOAD_SEED_DATA   WorkloadCommand = "Load Seed Data"
	MENU_ACADEMIC         WorkloadCommand = "Academic"
	MENU_ADMINISTRATIVE   WorkloadCommand = "Adminstrative"
	MENU_SENIOR_PROJECT   WorkloadCommand = "Senior Project"
	MENU_STUDENT_ADVISOR  WorkloadCommand = "Student Advisor"
	MENU_WORKLOAD_REPORT  WorkloadCommand = "Workload Report"
)
