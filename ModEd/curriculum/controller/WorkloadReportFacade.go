package controller

import (
	controller "ModEd/curriculum/controller"
	"fmt"
)

type WorkloadReportControllerFacade struct {
	title                    string
	classWorkloadService     controller.ClassWorkloadService
	meetingControllerService controller.MeetingControllerService
	projectControllerService controller.ProjectControllerService
	studentWorkloadService   controller.StudentWorkloadService
}

func CreateWorkloadReportFacade(
	classWorkload controller.ClassWorkloadService,
	meetingController controller.MeetingControllerService,
	projectController controller.ProjectControllerService,
	studentWorkload controller.StudentWorkloadService,
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
	lectures := f.classWorkloadService.GetClassLecturesByClassId(1)
	fmt.Println(lectures)

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
	advisingInfo, err := f.studentWorkloadService.GetByInstructorID(instructorID)
	if err != nil {
		fmt.Println("Error getting student advisor info:", err)
	} else {
		fmt.Println(advisingInfo)
	}

	fmt.Println("=====================================")
}
