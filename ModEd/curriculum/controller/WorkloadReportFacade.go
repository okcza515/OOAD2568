package controller

type WorkloadReportControllerFacade struct {
	title                    string
	classMaterialService     ClassMaterialService
	classLectureService      ClassLectureService
	meetingControllerService MeetingControllerService
	projectControllerService ProjectControllerService
	studentWorkloadService   StudentWorkloadService
}

func CreateWorkloadReportFacade(
	classMaterialService ClassMaterialService,
	classLectureService ClassLectureService,
	meetingController MeetingControllerService,
	projectController ProjectControllerService,
	studentWorkload StudentWorkloadService,
) *WorkloadReportControllerFacade {
	return &WorkloadReportControllerFacade{
		title:                    "Workload Report",
		classMaterialService:     classMaterialService,
		classLectureService:      classLectureService,
		meetingControllerService: meetingController,
		projectControllerService: projectController,
		studentWorkloadService:   studentWorkload,
	}
}

func (f *WorkloadReportControllerFacade) GenerateReport(instructorID uint) {
	//
}
