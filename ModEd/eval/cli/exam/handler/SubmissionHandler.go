package handler

// MEP-1007

import (
	assetUtil "ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"
)

type SubmissionMenuStateHandler struct {
	Manager                          *cli.CLIMenuStateManager
	wrapper                          *controller.ExamModuleWrapper
	SubmissionModuleMenuStateHandler cli.MenuState
	handler                          *handler.HandlerContext
	backhandler                      *handler.ChangeMenuHandlerStrategy
}

func NewSubmissionMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.ExamModuleWrapper, submissionModuleMenuStateHandler cli.MenuState) *SubmissionMenuStateHandler {
	return &SubmissionMenuStateHandler{
		Manager:                          manager,
		wrapper:                          wrapper,
		SubmissionModuleMenuStateHandler: submissionModuleMenuStateHandler,
		handler:                          handler.NewHandlerContext(),
		backhandler:                      handler.NewChangeMenuHandlerStrategy(manager, submissionModuleMenuStateHandler),
	}
}

func (menu *SubmissionMenuStateHandler) Render() {
	menu.handler.SetMenuTitle("\nSubmission management menu:")
	menu.handler.AddHandler("1", "Submit Exam", handler.FuncStrategy{Action: menu.SubmitExam})
	menu.handler.AddHandler("2", "View Score", handler.FuncStrategy{Action: menu.ViewScore})
	menu.handler.AddHandler("3", "List Submission", handler.FuncStrategy{Action: menu.ListSubmission})
	menu.handler.AddHandler("4", "Update Submission", handler.FuncStrategy{Action: menu.UpdateSubmission})
	menu.handler.AddHandler("5", "Delete Submission", handler.FuncStrategy{Action: menu.DeleteSubmission})
	menu.handler.AddBackHandler(menu.backhandler)
	menu.handler.ShowMenu()
}

func (menu *SubmissionMenuStateHandler) HandleUserInput(input string) error {
	return menu.handler.HandleInput(input)
}

func (menu *SubmissionMenuStateHandler) SubmitExam() error {
	var studentID uint
	fmt.Print("Enter Student ID: ")
	fmt.Scanln(&studentID)

	exams, err := menu.wrapper.ExamController.ListActiveExamsByStudentID(studentID)
	if err != nil {
		fmt.Printf("Exam with Student ID: %d not found\n", studentID)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
	}

	if len(exams) == 0 {
		fmt.Println("No exam found for submission.")
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
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
	err = menu.wrapper.SubmissionController.Insert(data)
	if err != nil {
		fmt.Printf("Error While Insert Submission: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
	}

	submission, err := menu.wrapper.SubmissionController.RetrieveByCondition(map[string]interface{}{"student_id": studentID, "exam_id": exams[examNo-1].ID})
	if err != nil {
		fmt.Printf("Error While Retrieve Submission: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
	}

	sections, err := menu.wrapper.ExamSectionController.List(map[string]interface{}{"exam_id": exams[examNo-1].ID})
	if err != nil {
		fmt.Printf("Error While List Sections: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
	}

	for _, section := range sections {
		fmt.Printf("\n===== Section: %d =====\n", section.SectionNo)
		fmt.Printf("Description: %s\n", section.Description)
		fmt.Printf("Number of Question: %d / Score: %f\n", section.NumQuestions, section.Score)
		fmt.Println("--------------------")

		questions, err := menu.wrapper.QuestionController.List(map[string]interface{}{"section_id": section.ID})
		if err != nil {
			fmt.Printf("Error While List Questions: %v\n", err)
			assetUtil.PressEnterToContinue()
			assetUtil.ClearScreen()
			return err
		}

		count = 0
		for _, question := range questions {
			count += 1
			fmt.Printf("Q%d: %s\n", count, question.ActualQuestion)

			if question.QuestionType == "MultipleChoiceQuestion" {
				mcAnswers, err := menu.wrapper.MultipleChoiceAnswerController.List(map[string]interface{}{"question_id": question.ID})
				if err != nil {
					fmt.Printf("Error While List Choices: %v\n", err)
					assetUtil.PressEnterToContinue()
					assetUtil.ClearScreen()
					return err
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
				err = menu.wrapper.MultipleChoiceAnswerSubmissionController.Insert(data)
				if err != nil {
					fmt.Printf("Error While Insert Multiple Choice Answer Submission: %v\n", err)
					assetUtil.PressEnterToContinue()
					assetUtil.ClearScreen()
					return err
				}

			} else if question.QuestionType == "ShortAnswerQuestion" {
				var answer string
				fmt.Printf("\nEnter the Correct Answer: ")
				fmt.Scanln(&answer)

				data := &model.ShortAnswerSubmission{
					QuestionID:    question.ID,
					SubmissionID:  submission.ID,
					StudentAnswer: answer,
				}
				err = menu.wrapper.ShortAnswerSubmissionController.Insert(data)
				if err != nil {
					fmt.Printf("Error While Insert Short Answer Submission: %v\n", err)
					assetUtil.PressEnterToContinue()
					assetUtil.ClearScreen()
					return err
				}

			} else if question.QuestionType == "TrueFalseQuestion" {
				var answer bool
				fmt.Printf("\nEnter the Correct Answer (true/false): ")
				fmt.Scanln(&answer)

				data := &model.TrueFalseAnswerSubmission{
					QuestionID:    question.ID,
					SubmissionID:  submission.ID,
					StudentAnswer: answer,
				}
				err = menu.wrapper.TrueFalseAnswerSubmissionController.Insert(data)
				if err != nil {
					fmt.Printf("Error While Insert True False Answer Submission: %v\n", err)
					assetUtil.PressEnterToContinue()
					assetUtil.ClearScreen()
					return err
				}
			}

			fmt.Println("--------------------")
		}
	}

	_, err = menu.wrapper.SubmissionController.GradingSubmission(submission.ID)
	if err != nil {
		fmt.Printf("Error While Grading Submission: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
	}

	fmt.Println("Exam Submitted")
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *SubmissionMenuStateHandler) ViewScore() error {
	var studentID uint
	fmt.Print("Enter Student ID: ")
	fmt.Scanln(&studentID)

	submissions, err := menu.wrapper.SubmissionController.List(map[string]interface{}{"student_id": studentID})
	if err != nil {
		fmt.Printf("Error While List Submission: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
	}

	if len(submissions) == 0 {
		fmt.Printf("Submission with Student ID: %d not found\n", studentID)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
	}

	fmt.Printf("\n===== Exam Submission (Student ID: %d) =====\n", studentID)
	for _, submission := range submissions {
		perfectScore, err := menu.wrapper.ExamController.GetPerfectScoreByExamID(submission.ExamID)
		if err != nil {
			fmt.Printf("Error While Get Perfect Score: %v\n", err)
			assetUtil.PressEnterToContinue()
			assetUtil.ClearScreen()
			return err
		}

		exam, err := menu.wrapper.ExamController.RetrieveByID(submission.ExamID)
		if err != nil {
			fmt.Printf("Error While Retrieve Exam: %v\n", err)
			assetUtil.PressEnterToContinue()
			assetUtil.ClearScreen()
			return err
		}

		fmt.Printf("Exam: %s\n", exam.ExamName)
		fmt.Printf("Score: %f/%f\n", submission.Score, perfectScore)
		fmt.Println("--------------------")
	}

	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *SubmissionMenuStateHandler) ListSubmission() error {
	var instructorID uint
	fmt.Print("Enter Instructor ID: ")
	fmt.Scanln(&instructorID)

	exams, err := menu.wrapper.ExamController.List(map[string]interface{}{"instructor_id": instructorID})
	if err != nil {
		fmt.Printf("Error While List Exams: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
	}

	if len(exams) == 0 {
		fmt.Printf("Exam with Instructor ID %d not found\n", instructorID)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
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

	submissions, err := menu.wrapper.SubmissionController.List(map[string]interface{}{"exam_id": exams[examNo-1].ID}, "Student")
	if len(submissions) == 0 {
		fmt.Printf("Submission of Exam: %s not found\n", exams[examNo-1].ExamName)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
	}

	perfectScore, err := menu.wrapper.ExamController.GetPerfectScoreByExamID(exams[examNo-1].ID)
	if err != nil {
		fmt.Printf("Error While Get Perfect Score: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return err
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

	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *SubmissionMenuStateHandler) UpdateSubmission() error {
	var instructorID uint
	fmt.Print("Enter Instructor ID: ")
	fmt.Scanln(&instructorID)

	exams, err := menu.wrapper.ExamController.List(map[string]interface{}{"instructor_id": instructorID})
	if err != nil {
		fmt.Printf("Error While Get Perfect Score: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
	}

	if len(exams) == 0 {
		fmt.Printf("Exam with Instructor ID: %d not found\n", instructorID)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
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

	submissions, err := menu.wrapper.SubmissionController.List(map[string]interface{}{"exam_id": exams[examNo-1].ID}, "Student")
	if len(submissions) == 0 {
		fmt.Printf("Submission of Exam: %s not found\n", exams[examNo-1].ExamName)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
	}

	perfectScore, err := menu.wrapper.ExamController.GetPerfectScoreByExamID(exams[examNo-1].ID)
	if err != nil {
		fmt.Printf("Error While Get Perfect Score: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
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
	err = menu.wrapper.SubmissionController.UpdateByID(submissions[submissionNo-1])
	if err != nil {
		fmt.Printf("Error While Update Submission: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
	}

	fmt.Println("Update Successful")
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}

func (menu *SubmissionMenuStateHandler) DeleteSubmission() error {
	var instructorID uint
	fmt.Print("Enter Instructor ID: ")
	fmt.Scanln(&instructorID)

	exams, err := menu.wrapper.ExamController.List(map[string]interface{}{"instructor_id": instructorID})
	if err != nil {
		fmt.Printf("Error While List Exams: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
	}

	if len(exams) == 0 {
		fmt.Printf("Exam with Instructor ID: %d not found\n", instructorID)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
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

	submissions, err := menu.wrapper.SubmissionController.List(map[string]interface{}{"exam_id": exams[examNo-1].ID}, "Student")
	if len(submissions) == 0 {
		fmt.Printf("Submission of Exam: %s not found\n", exams[examNo-1].ExamName)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
	}

	perfectScore, err := menu.wrapper.ExamController.GetPerfectScoreByExamID(exams[examNo-1].ID)
	if err != nil {
		fmt.Printf("Error While Get Perfect Score: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
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

	err = menu.wrapper.SubmissionController.DeleteByID(submissions[submissionNo-1].ID)
	if err != nil {
		fmt.Printf("Error While Delete Submission: %v\n", err)
		assetUtil.PressEnterToContinue()
		assetUtil.ClearScreen()
		return nil
	}

	fmt.Println("Delete Successful")
	assetUtil.PressEnterToContinue()
	assetUtil.ClearScreen()
	return nil
}
