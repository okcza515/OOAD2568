package cli

import (
	"ModEd/recruit/controller"
	"fmt"

	"gorm.io/gorm"
)

// ReportInterviewDetails แสดงข้อมูลสัมภาษณ์
func ReportInterviewDetails(db *gorm.DB, applicantID uint) {
	// ดึงสถานะจาก interview_status
	status, err := controller.GetApplicationStatus(db, applicantID)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการดึงสถานะการสมัคร:", err)
		return
	}

	if status != "Pass" {
		fmt.Println("คุณไม่มีสิทธิ์เข้าสัมภาษณ์ เนื่องจากสถานะของคุณคือ:", status)
		return
	}

	// ดึงข้อมูลสัมภาษณ์
	interview, err := controller.GetInterviewDetails(db, applicantID)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการดึงรายละเอียดการสัมภาษณ์:", err)
		return
	}

	//  เช็คว่า `InterviewScore` มีค่าหรือไม่
	scoreText := "N/A"
	if interview.InterviewScore != nil {
		scoreText = fmt.Sprintf("%.2f", *interview.InterviewScore) // แสดงทศนิยม 2 ตำแหน่ง
	}

	// แสดงรายละเอียดการสัมภาษณ์
	fmt.Println("--- รายละเอียดการสัมภาษณ์ ---")
	fmt.Println("วันที่:", interview.ScheduledAppointment)
	fmt.Println("คะแนนสัมภาษณ์:", scoreText)
	fmt.Println("สถานะการสัมภาษณ์:", interview.InterviewStatus)
}
