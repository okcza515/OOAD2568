package handler

type WorkloadCommand string
type AcademicMenuCommand string
type AdministrativeMenuCommand string
type SeniorProjectMenuCommand string
type StudentAdvisorMenuCommand string
type WorkloadReportMenuCommand string

const (
	MENU_LOAD_SEED_DATA   WorkloadCommand = "Load Seed Data"
	MENU_ACADEMIC         WorkloadCommand = "Academic"
	MENU_ADMINISTRATIVE   WorkloadCommand = "Adminstrative"
	MENU_SENIOR_PROJECT   WorkloadCommand = "Senior Project"
	MENU_STUDENT_ADVISOR  WorkloadCommand = "Student Advisor"
	MENU_WORKLOAD_REPORT  WorkloadCommand = "Workload Report"
	
	MENU_CURRICULUM    AcademicMenuCommand = "Curriculum"
	MENU_COURSE        AcademicMenuCommand = "Course"
	MENU_CLASS         AcademicMenuCommand = "Class"
	MENU_CLASSMATERIAL AcademicMenuCommand = "Class Material"
	MENU_COURSEPLAN    AcademicMenuCommand = "Course Plan"
	MENU_CREATECURRICULUM AcademicMenuCommand = "Create Curriculum"
	MENU_RETRIEVECURRICULUM AcademicMenuCommand = "Retrieve Curriculum"
	MENU_UPDATECURRICULUM AcademicMenuCommand = "Update Curriculum"
	MENU_DELETECURRICULUM AcademicMenuCommand = "Delete Curriculum"
	MENU_LISTCURRICULUM AcademicMenuCommand = "List Curriculum"
	MENU_CREATECOURSE     AcademicMenuCommand = "Create Course"
	MENU_RETRIEVECOURSE   AcademicMenuCommand = "Retrieve Course"
	MENU_UPDATECOURSE     AcademicMenuCommand = "Update Course"
	MENU_DELETECOURSE     AcademicMenuCommand = "Delete Course"
	MENU_LISTCOURSE       AcademicMenuCommand = "List Course"
	MENU_CREATECLASS      AcademicMenuCommand = "Create Class"
	MENU_RETRIEVECLASS    AcademicMenuCommand = "Retrieve Class"
	MENU_UPDATECLASS      AcademicMenuCommand = "Update Class"
	MENU_DELETECLASS      AcademicMenuCommand = "Delete Class"
	MENU_LISTCLASS        AcademicMenuCommand = "List Class"
	MENU_CREATECLASSMATERIAL AcademicMenuCommand = "Create Class Material"
	MENU_RETRIEVECLASSMATERIAL AcademicMenuCommand = "Retrieve Class Material"
	MENU_UPDATECLASSMATERIAL AcademicMenuCommand = "Update Class Material"
	MENU_DELETECLASSMATERIAL AcademicMenuCommand = "Delete Class Material"
	MENU_LISTCLASSMATERIAL   AcademicMenuCommand = "List Class Material"
	MENU_CREATECOURSEPLAN    AcademicMenuCommand = "Create Course Plan"
	MENU_RETRIEVECOURSEPLAN  AcademicMenuCommand = "Retrieve Course Plan"
	MENU_UPDATECOURSEPLAN    AcademicMenuCommand = "Update Course Plan"
	MENU_LISTCOURSEPLAN      AcademicMenuCommand = "List Course Plan"
	MENU_DELETECOURSEPLAN    AcademicMenuCommand = "Delete Course Plan"
	MENU_LISTUPCOMINGCOURSEPLANS AcademicMenuCommand = "List Upcoming Course Plans"
	
	MENU_MEETING         AdministrativeMenuCommand = "Meeting"
	MENU_STUDENT_REQUEST AdministrativeMenuCommand = "Student Request"
	MENU_MEETING_LIST    AdministrativeMenuCommand = "List All Meeting"
	MENU_MEETING_CREATE  AdministrativeMenuCommand = "Create Meeting"
	MENU_MEETING_CREATE_EXTERNAL AdministrativeMenuCommand = "Create External Meeting"
	MENU_MEETING_CREATE_ONLINE AdministrativeMenuCommand = "Create Online Meeting"
	MENU_MEETING_RETRIEVE AdministrativeMenuCommand = "Retrieve Meeting by Id"
	MENU_MEETING_ADD_ATTENDEE AdministrativeMenuCommand = "Add Attendee"
	MENU_MEETING_UPDATE  AdministrativeMenuCommand = "Update Meeting by Id"
	MENU_MEETING_DELETE  AdministrativeMenuCommand = "Delete Meeting by Id"

	MENU_SENIOR_PROJECT_VIEW_ADVISOR_PROJECT SeniorProjectMenuCommand = "View Advising Project"
	MENU_SENIOR_PROJECT_VIEW_COMMITTEE_PROJECT SeniorProjectMenuCommand = "View Committee Project"
	MENU_SENIOR_PROJECT_EVALUATE_PROJECT SeniorProjectMenuCommand = "Evaluate Project"

	MENU_STUDENT_ADVISOR_VIEW_ADVISOR_PROJECT StudentAdvisorMenuCommand = "View Advising Project"
	MENU_STUDENT_ADVISOR_VIEW_COMMITTEE_PROJECT StudentAdvisorMenuCommand = "View Committee Project"
	MENU_STUDENT_ADVISOR_EVALUATE_PROJECT StudentAdvisorMenuCommand = "Evaluate Project"

	MENU_WORKLOAD_REPORT_TODAY     WorkloadReportMenuCommand = "View Today Workload"
	MENU_WORKLOAD_REPORT_WEEKLY   WorkloadReportMenuCommand = "View Weekly Workload"
	MENU_WORKLOAD_REPORT_MONTHLY  WorkloadReportMenuCommand = "View Monthly Workload"
	MENU_WORKLOAD_REPORT_GENERATE WorkloadReportMenuCommand = "Generate Performance Report"
)
