// MEP-1007
package handler

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"

	"gorm.io/gorm"
)

type SubmissionHandler struct {
	db *gorm.DB
}

func NewSubmissionHandler(db *gorm.DB) SubmissionHandler {
	return SubmissionHandler{db: db}
}

type Back struct{}

func (b Back) Execute() {
	return
}

type UnknownCommand struct{}

func (u UnknownCommand) Execute() {
	fmt.Println("Unknown command, please try again.")
}

func (s SubmissionHandler) Execute() {
	menu := NewMenuHandler("Submission Options", true)
	menu.Add("Submit Exam", SubmitExam{db: s.db})
	menu.Add("View Score", ViewScore{db: s.db})
	menu.Add("List Exam Submissions", ListSubmission{db: s.db})
	menu.Add("Update Submission Score", UpdateSubmission{db: s.db})
	menu.Add("Delete Submission", DeleteSubmission{db: s.db})

	menu.SetBackHandler(Back{})
	menu.SetDefaultHandler(UnknownCommand{})
	menu.Execute()
}

type SubmitExam struct {
	db *gorm.DB
}

func (s SubmitExam) Execute() {
	ExamCtrl := controller.NewExamController(s.db)
	ExamSectionCtrl := controller.NewExamSectionController(s.db)
	QuestionCtrl := controller.NewQuestionController(s.db)
	MCAnsCtrl := controller.NewMultipleChoiceAnswerController(s.db)

	SubmissionCtrl := controller.NewSubmissionController(s.db)
	MCAnsSubCtrl := controller.NewMultipleChoiceAnswerSubmissionController(s.db)
	ShortAnsSubCtrl := controller.NewShortAnswerSubmissionController(s.db)
	TFAnsSubCtrl := controller.NewTrueFalseAnswerSubmissionController(s.db)

	var studentID uint
	fmt.Print("Enter Student ID: ")
	fmt.Scanln(&studentID)

	exams, err := ExamCtrl.ListActiveExamsByStudentID(studentID)
	if err != nil {
		fmt.Printf("Error: Exam with Student ID: %d not found: %v\n", studentID, err)
		return
	}

	if len(exams) == 0 {
		fmt.Println("No exam found for submission.")
		return
	}

	fmt.Printf("\n===== Exam for Submission (Student ID: %d) =====\n", studentID)

	var count = 0
	for _, exam := range exams {
		count += 1
		fmt.Printf("%d. %s\n", count, exam.ExamName)
		fmt.Printf("Description: %s\n", exam.Description)
		fmt.Printf("Start: %s / End: %s\n", exam.StartDate, exam.EndDate)
		fmt.Println("--------------------")
	}

	var examNo uint
	fmt.Print("Enter Exam No. for Submission: ")
	fmt.Scanln(&examNo)

	data := &model.AnswerSubmission{
		StudentID: studentID,
		ExamID:    exams[examNo-1].ID,
		Score:     0,
	}
	err = SubmissionCtrl.Insert(data)
	if err != nil {
		return
	}

	submission, err := SubmissionCtrl.RetrieveByCondition(map[string]interface{}{"student_id": studentID, "exam_id": exams[examNo-1].ID})
	if err != nil {
		return
	}

	sections, err := ExamSectionCtrl.List(map[string]interface{}{"exam_id": exams[examNo-1].ID})
	if err != nil {
		return
	}

	for _, section := range sections {
		fmt.Printf("\n===== Section: %d =====\n", section.SectionNo)
		fmt.Printf("Description: %s\n", section.Description)
		fmt.Printf("Number of Question: %d / Score: %f\n", section.NumQuestions, section.Score)
		fmt.Println("--------------------")

		questions, err := QuestionCtrl.List(map[string]interface{}{"section_id": section.ID})
		if err != nil {
			return
		}

		count = 0
		for _, question := range questions {
			count += 1
			fmt.Printf("Q%d: %s\n", count, question.ActualQuestion)

			if question.QuestionType == "MultipleChoiceQuestion" {
				mcAnswers, err := MCAnsCtrl.List(map[string]interface{}{"question_id": question.ID})
				if err != nil {
					return
				}

				var choiceCount = 0
				for _, mcAnswer := range mcAnswers {
					choiceCount += 1
					fmt.Printf("%d. %s\n", choiceCount, mcAnswer.AnswerLabel)
				}

				var answer uint
				fmt.Printf("\nEnter the Correct Answer (1-%d): ", choiceCount)
				fmt.Scanln(&answer)

				data := &model.MultipleChoiceAnswerSubmission{
					QuestionID:   question.ID,
					SubmissionID: submission.ID,
					ChoiceID:     mcAnswers[answer-1].ID,
				}
				err = MCAnsSubCtrl.Insert(data)

			} else if question.QuestionType == "ShortAnswerQuestion" {
				var answer string
				fmt.Printf("\nEnter the Correct Answer: ")
				fmt.Scanln(&answer)

				data := &model.ShortAnswerSubmission{
					QuestionID:    question.ID,
					SubmissionID:  submission.ID,
					StudentAnswer: answer,
				}
				err = ShortAnsSubCtrl.Insert(data)

			} else if question.QuestionType == "TrueFalseQuestion" {
				var answer bool
				fmt.Printf("\nEnter the Correct Answer (true/false): ")
				fmt.Scanln(&answer)

				data := &model.TrueFalseAnswerSubmission{
					QuestionID:    question.ID,
					SubmissionID:  submission.ID,
					StudentAnswer: answer,
				}
				err = TFAnsSubCtrl.Insert(data)
			}

			fmt.Println("--------------------")
		}
	}

	_, err = SubmissionCtrl.GradingSubmission(submission.ID)
	if err != nil {
		return
	}

	fmt.Println("Exam Submitted")
	return
}

type ViewScore struct {
	db *gorm.DB
}

func (v ViewScore) Execute() {
	ExamCtrl := controller.NewExamController(v.db)
	SubmissionCtrl := controller.NewSubmissionController(v.db)

	var studentID uint
	fmt.Print("Enter Student ID: ")
	fmt.Scanln(&studentID)

	submissions, err := SubmissionCtrl.List(map[string]interface{}{"student_id": studentID}, "Exam")
	if err != nil {
		fmt.Printf("Error: Exam Submission with Student ID: %d not found: %v\n", studentID, err)
		return
	}

	fmt.Printf("\n===== Exam Submission (Student ID: %d) =====\n", studentID)
	for _, submission := range submissions {
		perfectScore, err := ExamCtrl.GetPerfectScoreByExamID(submission.ExamID)
		if err != nil {
			return
		}

		fmt.Printf("Exam: %s\n", submission.Examination.ExamName)
		fmt.Printf("Score: %f/%f\n", submission.Score, perfectScore)
		fmt.Println("--------------------")
	}

	return
}

type ListSubmission struct {
	db *gorm.DB
}

func (l ListSubmission) Execute() {
	ExamCtrl := controller.NewExamController(l.db)
	SubmissionCtrl := controller.NewSubmissionController(l.db)

	var instructorID uint
	fmt.Print("Enter Instructor ID: ")
	fmt.Scanln(&instructorID)

	exams, err := ExamCtrl.List(map[string]interface{}{"instructor_id": instructorID})
	if err != nil {
		return
	}

	if len(exams) == 0 {
		fmt.Printf("Exam with Instructor ID %d not found\n", instructorID)
		return
	}

	fmt.Printf("\n===== Exams List (Instructor ID: %d) =====\n", instructorID)

	var count = 0
	for _, exam := range exams {
		count += 1
		fmt.Printf("%d. %s\n", count, exam.ExamName)
		fmt.Printf("Description: %s\n", exam.Description)
		fmt.Printf("Start: %s / End: %s\n", exam.StartDate, exam.EndDate)
		fmt.Println("--------------------")
	}

	var examNo uint
	fmt.Print("Enter Exam No. for Submissions List: ")
	fmt.Scanln(&examNo)

	submissions, err := SubmissionCtrl.List(map[string]interface{}{"exam_id": exams[examNo-1].ID}, "Student")
	if len(submissions) == 0 {
		fmt.Printf("Submission of Exam: %s not found\n", exams[examNo-1].ExamName)
		return
	}
	
	perfectScore, err := ExamCtrl.GetPerfectScoreByExamID(exams[examNo-1].ID)
	if err != nil {
		return
	}
	
	fmt.Printf("\n===== Submission List (Exam: %s) =====\n", exams[examNo-1].ExamName)

	count = 0
	for _, submission := range submissions {
		count += 1
		fmt.Printf("Student ID: %d\n", submission.StudentID)
		fmt.Printf("Student Name: %s %s\n", submission.Student.FirstName, submission.Student.LastName)
		fmt.Printf("Score: %f/%f\n", submission.Score, perfectScore)
		fmt.Println("--------------------")
	}

	return
}

type UpdateSubmission struct {
	db *gorm.DB
}

func (l UpdateSubmission) Execute() {
	ExamCtrl := controller.NewExamController(l.db)
	SubmissionCtrl := controller.NewSubmissionController(l.db)

	var instructorID uint
	fmt.Print("Enter Instructor ID: ")
	fmt.Scanln(&instructorID)

	exams, err := ExamCtrl.List(map[string]interface{}{"instructor_id": instructorID})
	if err != nil {
		return
	}

	fmt.Printf("\n===== Exams List (Instructor ID: %d) =====\n", instructorID)

	var count = 0
	for _, exam := range exams {
		count += 1
		fmt.Printf("%d. %s\n", count, exam.ExamName)
		fmt.Printf("Description: %s\n", exam.Description)
		fmt.Printf("Start: %s / End: %s\n", exam.StartDate, exam.EndDate)
		fmt.Println("--------------------")
	}

	if len(exams) == 0 {
		fmt.Printf("Exam with Instructor ID: %d not found\n", instructorID)
		return
	}

	var examNo uint
	fmt.Print("Enter Exam No. for Submissions List: ")
	fmt.Scanln(&examNo)

	submissions, err := SubmissionCtrl.List(map[string]interface{}{"exam_id": exams[examNo-1].ID}, "Student")
	if len(submissions) == 0 {
		fmt.Printf("Submission of Exam: %s not found\n", exams[examNo-1].ExamName)
		return
	}

	perfectScore, err := ExamCtrl.GetPerfectScoreByExamID(exams[examNo-1].ID)
	if err != nil {
		return
	}

	fmt.Printf("\n===== Submission List (Exam: %s) =====\n", exams[examNo-1].ExamName)

	count = 0
	for _, submission := range submissions {
		count += 1
		fmt.Printf("%d. Student ID: %d\n", count, submission.StudentID)
		fmt.Printf("Student Name: %s %s\n", submission.Student.FirstName, submission.Student.LastName)
		fmt.Printf("Score: %f/%f\n", submission.Score, perfectScore)
		fmt.Println("--------------------")
	}

	var submissionNo uint
	fmt.Print("Enter Submission No. for Update Score: ")
	fmt.Scanln(&submissionNo)

	var newScore float64
	fmt.Print("Enter New Score: ")
	fmt.Scanln(&newScore)

	submissions[submissionNo-1].Score = newScore
	err = SubmissionCtrl.UpdateByID(submissions[submissionNo-1])
	if err != nil {
		return
	}

	fmt.Println("Update Successful")
	return
}

type DeleteSubmission struct {
	db *gorm.DB
}

func (l DeleteSubmission) Execute() {
	ExamCtrl := controller.NewExamController(l.db)
	SubmissionCtrl := controller.NewSubmissionController(l.db)

	var instructorID uint
	fmt.Print("Enter Instructor ID: ")
	fmt.Scanln(&instructorID)

	exams, err := ExamCtrl.List(map[string]interface{}{"instructor_id": instructorID})
	if err != nil {
		return
	}

	fmt.Printf("\n===== Exams List (Instructor ID: %d) =====\n", instructorID)

	var count = 0
	for _, exam := range exams {
		count += 1
		fmt.Printf("%d. %s\n", count, exam.ExamName)
		fmt.Printf("Description: %s\n", exam.Description)
		fmt.Printf("Start: %s / End: %s\n", exam.StartDate, exam.EndDate)
		fmt.Println("--------------------")
	}

	if len(exams) == 0 {
		fmt.Printf("Exam with Instructor ID: %d not found\n", instructorID)
		return
	}

	var examNo uint
	fmt.Print("Enter Exam No. for Submissions List: ")
	fmt.Scanln(&examNo)

	submissions, err := SubmissionCtrl.List(map[string]interface{}{"exam_id": exams[examNo-1].ID}, "Student")
	if len(submissions) == 0 {
		fmt.Printf("Submission of Exam: %s not found\n", exams[examNo-1].ExamName)
		return
	}

	perfectScore, err := ExamCtrl.GetPerfectScoreByExamID(exams[examNo-1].ID)
	if err != nil {
		return
	}

	fmt.Printf("\n===== Submission List (Exam: %s) =====\n", exams[examNo-1].ExamName)

	count = 0
	for _, submission := range submissions {
		count += 1
		fmt.Printf("Student ID: %d\n", submission.StudentID)
		fmt.Printf("Student Name: %s %s\n", submission.Student.FirstName, submission.Student.LastName)
		fmt.Printf("Score: %f/%f\n", submission.Score, perfectScore)
		fmt.Println("--------------------")
	}

	var submissionNo uint
	fmt.Print("Enter Submission No. for Delete: ")
	fmt.Scanln(&submissionNo)

	err = SubmissionCtrl.DeleteByID(submissions[submissionNo-1].ID)
	if err != nil {
		return
	}

	fmt.Println("Delete Successful")
	return
}
