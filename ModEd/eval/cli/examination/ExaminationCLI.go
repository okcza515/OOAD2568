package cli

import (
	controllerExamination "ModEd/eval/controller"
	"ModEd/eval/model"
	"ModEd/eval/util"

	"fmt"
	// "log"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RunExaminationCLI(db *gorm.DB) {

	examController := controllerExamination.NewExaminationController(db)

	for{
		fmt.Println("\nExamination CLI")
		fmt.Println("1. Create Examination")
		fmt.Println("2. Display All Examination")
		fmt.Println("3. Update Examination")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			CreateExamination(db,examController)
		case 2:
			DisplayAllExams(db,examController)
		case 3:
			var examID uint
			fmt.Print("Enter Examination ID to update: ")
			fmt.Scan(&examID)
			UpdateExam(db, examController, examID)
			fmt.Println("Examination updated successfully!")
		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func DisplayAllExams(db *gorm.DB, examController *controllerExamination.ExaminationController) {
	exams,err := examController.GetAllExam()
	if err != nil {
		fmt.Println("Error fetching exams:", err)
		return
	}
	if len(exams) == 0 {
		fmt.Println("No exams found")
		return
	}

	fmt.Println("\nList of Examinations:")
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("%-5s | %-20s | %-10s | %-10s | %-10s | %-15s | %-20s | %-10s\n",
		"ID", "Exam Name", "Instructor", "Course", "Curriculum", "Criteria", "Description", "Exam Date")
	fmt.Println("-----------------------------------------------------")

	for _, exam := range exams {
		fmt.Printf("%-5d | %-20s | %-10d | %-10d | %-10d | %-15s | %-20s | %-10s\n",
			exam.ID, exam.Exam_name, exam.Instructor_id, exam.CourseId, exam.CurriculumId,
			exam.Criteria, exam.Description, exam.Exam_date.Format("2006-01-02"))
	}
}

func getExamInput() (*model.Examination, error) {
	examName := util.PromptString("Enter Exam Name: ")

	instructorId, err := util.PromptUint("Enter Instructor ID: ")
	if err != nil {
		return nil, fmt.Errorf("invalid instructor ID: %w", err)
	}

	courseId, err := util.PromptUint("Enter Course ID: ")
	if err != nil {
		return nil, fmt.Errorf("invalid course ID: %w", err)
	}

	curriculumId, err := util.PromptUint("Enter Curriculum ID: ")
	if err != nil {
		return nil, fmt.Errorf("invalid curriculum ID: %w", err)
	}

	criteria := util.PromptString("Enter Criteria: ")
	description := util.PromptString("Enter Description: ")

	examDate, err := util.PromptDate("Enter Exam Date (YYYY-MM-DD): ")
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	return &model.Examination{
		Exam_name:     examName,
		Instructor_id: uint(instructorId),
		CourseId:      uint(courseId),
		CurriculumId:  uint(curriculumId),
		Criteria:      criteria,
		Description:   description,
		Exam_date:     examDate,
	}, nil
}

func CreateExamination(db *gorm.DB, examController *controllerExamination.ExaminationController) {
	exam, err := getExamInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := examController.CreateExam(exam); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Examination created successfully!")
	}
}

func UpdateExam(db *gorm.DB, examController *controllerExamination.ExaminationController, examId uint) {
	exam, err := getExamInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := examController.UpdateExam(examId, exam); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Examination updated successfully!")
	}
}