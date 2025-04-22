package controller

import (
	"fmt"
)

type WorkloadReportControllerFacade struct {
	title                    string
	classWorkloadService     ClassWorkloadService
	meetingControllerService MeetingControllerService
	projectControllerService ProjectControllerService
	studentWorkloadService   StudentWorkloadService
}

func CreateWorkloadReportFacade(
	classWorkload ClassWorkloadService,
	meetingController MeetingControllerService,
	projectController ProjectControllerService,
	studentWorkload StudentWorkloadService,
) *WorkloadReportControllerFacade {
	return &WorkloadReportControllerFacade{
		title:                    "Workload Report",
		classWorkloadService:     classWorkload,
		meetingControllerService: meetingController,
		projectControllerService: projectController,
		studentWorkloadService:   studentWorkload,
	}
}

func (f *WorkloadReportControllerFacade) GenerateReport(instructorID uint) {
	fmt.Println("==========", f.title, "==========")

	fmt.Println("\n[1] Academic Class Workload:")
	lectures, err := f.classWorkloadService.GetClassLecturesByClassId(1)
	if err != nil {
		fmt.Println("Error getting lectures:", err)
	} else {
		fmt.Println(lectures)
	}

	fmt.Println("\n[2] Administrative Meetings:")
	meetings, err := f.meetingControllerService.GetAll()
	if err != nil {
		fmt.Println("Error getting meetings:", err)
	} else {
		for _, m := range *meetings {
			fmt.Printf("Meeting: %s, Location: %s\n", m.Title, m.Location)
		}
	}

	fmt.Println("\n[3] Senior Project Involvement:")
	projects, err := f.projectControllerService.GetProjectByAdvisorID(instructorID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(projects)
	}

	fmt.Println("\n[4] Student Advisor Info:")
	advisingInfo, err := f.studentWorkloadService.GetStudentRequestsByInstructorId(instructorID)
	if err != nil {
		fmt.Println("Error getting student advisor info:", err)
	} else {
		fmt.Println(advisingInfo)
	}

	fmt.Println("=====================================")
}
