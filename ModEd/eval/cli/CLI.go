package cli

import (
	"fmt"
	"log"
	"time"

	"ModEd/eval/controller"
	evalModel "ModEd/eval/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CLI interface {
	Run()
}

type Command interface {
	Execute()
}

type CLIFactory struct {
	db *gorm.DB
}

func NewCLIFactory(db *gorm.DB) *CLIFactory {
	return &CLIFactory{db: db}
}

func (f *CLIFactory) CreateCLI(cliType string) CLI {
	switch cliType {
	case "assignment":
		return NewAssignmentCLI(controller.NewAssignmentController(f.db))
	case "evaluation":
		return NewEvaluationCLI(f.db)
	case "progress":
		return NewProgressCLI(controller.NewProgressController(f.db))
	case "quiz":
		return createQuizCLI(controller.NewQuizController(f.db))
	case "question":
		return NewQuestionCLI(f.db)
	case "answer":
		return NewAnswerCLI(f.db)
	default:
		log.Fatalf("Unsupported CLI type: %s", cliType)
		return nil
	}
}

type MainCLI struct {
	db         *gorm.DB
	cliFactory *CLIFactory
}

func NewMainCLI() *MainCLI {
	db, err := gorm.Open(sqlite.Open("moded.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = db.AutoMigrate(
		&evalModel.Assignment{},
		&evalModel.AssignmentSubmission{},
		&evalModel.Evaluation{},
		&evalModel.Quiz{},
		&evalModel.Question{},
		&evalModel.Answer{},
	)
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	return &MainCLI{
		db:         db,
		cliFactory: NewCLIFactory(db),
	}
}

func (m *MainCLI) Run() {
	for {
		fmt.Println("\n===== ModEd CLI =====")
		fmt.Println("1. Assignment")
		fmt.Println("2. Evaluation")
		fmt.Println("3. Progress")
		fmt.Println("4. Quiz")
		fmt.Println("5. Question")
		fmt.Println("6. Answer")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			m.cliFactory.CreateCLI("assignment").Run()
		case 2:
			m.cliFactory.CreateCLI("evaluation").Run()
		case 3:
			m.cliFactory.CreateCLI("progress").Run()
		case 4:
			m.cliFactory.CreateCLI("quiz").Run()
		case 5:
			m.cliFactory.CreateCLI("question").Run()
		case 6:
			m.cliFactory.CreateCLI("answer").Run()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

// Main entry point for the CLI
func RunCLI() {
	cli := NewMainCLI()
	cli.Run()
}

// Evaluation-related structures and functions
type EvaluationCLI struct {
	db *gorm.DB
}

type Progress struct {
	StudentCode  string
	Title        string
	AssignmentID uint
	Status       string
	LastUpdate   time.Time
	TotalSubmit  uint
}

type evaluationControllerAdapter struct {
	db *gorm.DB
}

func (ec *evaluationControllerAdapter) SaveEvaluation(studentCode string, instructorCode string, assignmentID, quizID *uint, score float64, comment string) error {
	e := evalModel.Evaluation{
		StudentCode:    studentCode,
		InstructorCode: instructorCode,
		AssignmentID:   assignmentID,
		QuizID:         quizID,
		Score:          uint(score),
		Comment:        comment,
		EvaluatedAt:    time.Now(),
	}
	return ec.db.Create(&e).Error
}

func (ec *evaluationControllerAdapter) ListEvaluations() ([]evalModel.Evaluation, error) {
	var evals []evalModel.Evaluation
	err := ec.db.Find(&evals).Error
	return evals, err
}

func NewEvaluationCLI(db *gorm.DB) *EvaluationCLI {
	return &EvaluationCLI{db: db}
}

func (cli *EvaluationCLI) Run() {
	for {
		fmt.Println("\n===== Evaluation Menu =====")
		fmt.Println("1. Show Evaluation Info")
		fmt.Println("2. Record New Evaluation")
		fmt.Println("3. Back to Main Menu")
		fmt.Print("Choose an option: ")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			cli.listEvaluations()
		case 2:
			cli.recordEvaluation()
		case 3:
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (cli *EvaluationCLI) listEvaluations() {
	evalController := &evaluationControllerAdapter{db: cli.db}

	evals, err := evalController.ListEvaluations()
	if err != nil {
		fmt.Printf("Error retrieving evaluations: %v\n", err)
		return
	}

	fmt.Println("\n===== Evaluations =====")
	for i, e := range evals {
		fmt.Printf("Entry: %d | Student: %s | Score: %d | Date: %s\n",
			i+1, e.StudentCode, e.Score, e.EvaluatedAt.Format("2006-01-02"))
		if e.AssignmentID != nil {
			fmt.Printf("  Assignment ID: %d\n", *e.AssignmentID)
		}
		if e.QuizID != nil {
			fmt.Printf("  Quiz ID: %d\n", *e.QuizID)
		}
		if e.Comment != "" {
			fmt.Printf("  Comment: %s\n", e.Comment)
		}
		fmt.Println("------------------------")
	}
}

func (cli *EvaluationCLI) recordEvaluation() {
	evalController := &evaluationControllerAdapter{db: cli.db}

	var studentCode, instructorCode, comment string
	var score float64
	var evalType, id uint

	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)

	fmt.Print("Enter Instructor Code: ")
	fmt.Scanln(&instructorCode)

	fmt.Println("Evaluation Type: 1. Assignment, 2. Quiz")
	fmt.Print("Enter choice (1/2): ")
	fmt.Scanln(&evalType)

	fmt.Print("Enter ID of Assignment/Quiz: ")
	fmt.Scanln(&id)

	fmt.Print("Enter Score: ")
	fmt.Scanln(&score)

	fmt.Print("Enter Comment: ")
	fmt.Scanln(&comment)

	var assignmentID, quizID *uint

	if evalType == 1 {
		assignmentID = &id
	} else {
		quizID = &id
	}

	err := evalController.SaveEvaluation(studentCode, instructorCode, assignmentID, quizID, score, comment)
	if err != nil {
		fmt.Printf("Error saving evaluation: %v\n", err)
		return
	}

	fmt.Println("Evaluation recorded successfully!")
}

type ProgressCLI struct {
	controller *controller.ProgressController
}

func NewProgressCLI(controller *controller.ProgressController) *ProgressCLI {
	return &ProgressCLI{controller: controller}
}

func (cli *ProgressCLI) Run() {
	for {
		fmt.Println("\n===== Progress Menu =====")
		fmt.Println("1. View All Progress")
		fmt.Println("2. View Progress by Student Code")
		fmt.Println("3. View Progress by Status")
		fmt.Println("4. Back to Main Menu")
		fmt.Print("Choose an option: ")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			cli.viewAllProgress()
		case 2:
			cli.viewProgressByStudent()
		case 3:
			cli.viewProgressByStatus()
		case 4:
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (cli *ProgressCLI) viewAllProgress() {
	var assignmentID uint
	fmt.Print("Enter Assignment ID: ")
	fmt.Scan(&assignmentID)

	progress, err := cli.controller.GetAllProgress(assignmentID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cli.displayProgress(progress)
}

func (cli *ProgressCLI) viewProgressByStudent() {
	var assignmentID uint
	var studentCode string

	fmt.Print("Enter Assignment ID: ")
	fmt.Scan(&assignmentID)
	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)

	progress, err := cli.controller.GetProgressByStudentCode(assignmentID, studentCode)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cli.displayProgress(progress)
}

func (cli *ProgressCLI) viewProgressByStatus() {
	var assignmentID uint
	var status string

	fmt.Print("Enter Assignment ID: ")
	fmt.Scan(&assignmentID)
	fmt.Print("Enter Status (Open/Closed/Submitted): ")
	fmt.Scanln(&status)

	progress, err := cli.controller.GetProgressByStatus(assignmentID, status)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cli.displayProgress(progress)
}

func (cli *ProgressCLI) displayProgress(progressList []controller.Progress) {
	fmt.Println("\n===== Student Progress =====")
	for _, p := range progressList {
		fmt.Printf("ID: %d | Student: %v | Status: %v\n", p.Model.ID, p.StudentCode, p.Status)
		fmt.Printf("  Last Update: %s | Total Submissions: %d\n",
			p.LastUpdate.Format("2006-01-02 15:04"), p.TotalSubmit)
		fmt.Println("------------------------")
	}
}

type quizCLIImpl struct {
	controller controller.QuizController
}

func createQuizCLI(controller controller.QuizController) *quizCLIImpl {
	return &quizCLIImpl{controller: controller}
}

func (cli *quizCLIImpl) Run() {
	for {
		var option int
		fmt.Println("\n===== Quiz Menu =====")
		fmt.Println("1. List Quizzes")
		fmt.Println("2. Create Quiz")
		fmt.Println("3. Update Quiz")
		fmt.Println("4. Delete Quiz")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Choose an option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			cli.listQuizzes()
		case 2:
			cli.createQuiz()
		case 3:
			cli.updateQuiz()
		case 4:
			cli.deleteQuiz()
		case 5:
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (cli *quizCLIImpl) listQuizzes() {
	quizzes, err := cli.controller.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, q := range quizzes {
		fmt.Printf("ID: %d | Title: %s | Released: %t | Status: %s\n",
			q.ID, q.Title, q.Released, q.Status)
		fmt.Printf("  Start: %s | End: %s\n",
			q.QuizStart.Format("2006-01-02 15:04"),
			q.QuizEnd.Format("2006-01-02 15:04"))
		fmt.Println("------------------------")
	}
}

func (cli *quizCLIImpl) createQuiz() {
	var input controller.QuizInput
	fmt.Print("Enter Title: ")
	fmt.Scanln(&input.Title)
	fmt.Print("Enter Description: ")
	fmt.Scanln(&input.Description)

	var released string
	fmt.Print("Released (yes/no): ")
	fmt.Scanln(&released)
	input.Released = (released == "yes")

	input.QuizStart = time.Now()
	input.QuizEnd = time.Now().Add(1 * time.Hour)

	fmt.Print("Enter Status (Scheduled/Active/Completed): ")
	fmt.Scanln(&input.Status)

	quiz, err := cli.controller.Create(input)
	if err != nil {
		fmt.Println("Error creating quiz:", err)
		return
	}
	fmt.Println("Created Quiz:", quiz.ID)
}

func (cli *quizCLIImpl) updateQuiz() {
	var id uint
	fmt.Print("Enter Quiz ID to update: ")
	fmt.Scan(&id)

	var input controller.QuizInput
	fmt.Print("Enter new Title: ")
	fmt.Scanln(&input.Title)
	fmt.Print("Enter new Description: ")
	fmt.Scanln(&input.Description)

	var released string
	fmt.Print("Released (yes/no): ")
	fmt.Scanln(&released)
	input.Released = (released == "yes")

	fmt.Print("Enter Status (Scheduled/Active/Completed): ")
	fmt.Scanln(&input.Status)

	input.QuizStart = time.Now()
	input.QuizEnd = time.Now().Add(1 * time.Hour)

	quiz, err := cli.controller.Update(id, input)
	if err != nil {
		fmt.Println("Error updating quiz:", err)
		return
	}
	fmt.Println("Updated Quiz:", quiz.ID)
}

func (cli *quizCLIImpl) deleteQuiz() {
	var id uint
	fmt.Print("Enter Quiz ID to delete: ")
	fmt.Scan(&id)

	if err := cli.controller.Delete(id); err != nil {
		fmt.Println("Error deleting quiz:", err)
		return
	}
	fmt.Println("Deleted Quiz:", id)
}

type QuestionCLI struct {
	db *gorm.DB
}

func NewQuestionCLI(db *gorm.DB) *QuestionCLI {
	return &QuestionCLI{db: db}
}

func (cli *QuestionCLI) Run() {
	for {
		fmt.Println("\n===== Question Menu =====")
		fmt.Println("1. List Questions by Exam")
		fmt.Println("2. Add Question to Exam")
		fmt.Println("3. Update Question")
		fmt.Println("4. Delete Question")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Choose an option: ")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			cli.listQuestions()
		case 2:
			cli.addQuestion()
		case 3:
			cli.updateQuestion()
		case 4:
			cli.deleteQuestion()
		case 5:
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (cli *QuestionCLI) listQuestions() {
	var examID uint
	fmt.Print("Enter Exam ID: ")
	fmt.Scan(&examID)

	var questions []evalModel.Question
	if err := cli.db.Where("exam_id = ?", examID).Find(&questions).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n===== Questions =====")
	for _, q := range questions {
		fmt.Printf("ID: %d | Detail: %s | Score: %f\n",
			q.ID, q.Question_detail, q.Score)
		fmt.Printf("  Type: %s\n", q.Question_type)
		fmt.Println("------------------------")
	}
}

func (cli *QuestionCLI) addQuestion() {
	var examID uint
	fmt.Print("Enter Exam ID: ")
	fmt.Scan(&examID)

	var detail string
	fmt.Print("Enter Question Detail: ")
	fmt.Scanln(&detail)

	var qType string
	fmt.Print("Enter Question Type (Multiple_choice/Short_answer/True_false/Subjective): ")
	fmt.Scanln(&qType)

	var correctAnswer string
	fmt.Print("Enter Correct Answer: ")
	fmt.Scanln(&correctAnswer)

	var score float64
	fmt.Print("Enter Score: ")
	fmt.Scan(&score)

	question := evalModel.Question{
		Exam_id:         examID,
		Question_detail: detail,
		Question_type:   evalModel.QuestionType(qType),
		Correct_answer:  correctAnswer,
		Score:           score,
	}

	if err := cli.db.Create(&question).Error; err != nil {
		fmt.Println("Error creating question:", err)
		return
	}

	fmt.Println("Created Question:", question.ID)
}

func (cli *QuestionCLI) updateQuestion() {
	var id uint
	fmt.Print("Enter Question ID to update: ")
	fmt.Scan(&id)

	var question evalModel.Question
	if err := cli.db.First(&question, id).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	var detail string
	fmt.Print("Enter new Question Detail: ")
	fmt.Scanln(&detail)

	var qType string
	fmt.Print("Enter new Question Type (Multiple_choice/Short_answer/True_false/Subjective): ")
	fmt.Scanln(&qType)

	var correctAnswer string
	fmt.Print("Enter new Correct Answer: ")
	fmt.Scanln(&correctAnswer)

	var score float64
	fmt.Print("Enter new Score: ")
	fmt.Scan(&score)

	question.Question_detail = detail
	question.Question_type = evalModel.QuestionType(qType)
	question.Correct_answer = correctAnswer
	question.Score = score

	if err := cli.db.Save(&question).Error; err != nil {
		fmt.Println("Error updating question:", err)
		return
	}

	fmt.Println("Updated Question:", question.ID)
}

func (cli *QuestionCLI) deleteQuestion() {
	var id uint
	fmt.Print("Enter Question ID to delete: ")
	fmt.Scan(&id)

	if err := cli.db.Delete(&evalModel.Question{}, id).Error; err != nil {
		fmt.Println("Error deleting question:", err)
		return
	}

	fmt.Println("Deleted Question:", id)
}

// Answer CLI implementation - Strategy Pattern
type AnswerCLI struct {
	db *gorm.DB
}

func NewAnswerCLI(db *gorm.DB) *AnswerCLI {
	return &AnswerCLI{db: db}
}

func (cli *AnswerCLI) Run() {
	for {
		fmt.Println("\n===== Answer Menu =====")
		fmt.Println("1. List Answers by Question")
		fmt.Println("2. Add Answer")
		fmt.Println("3. Update Answer")
		fmt.Println("4. Delete Answer")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Choose an option: ")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			cli.listAnswers()
		case 2:
			cli.addAnswer()
		case 3:
			cli.updateAnswer()
		case 4:
			cli.deleteAnswer()
		case 5:
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (cli *AnswerCLI) listAnswers() {
	var questionID uint
	fmt.Print("Enter Question ID: ")
	fmt.Scan(&questionID)

	var answers []evalModel.Answer
	if err := cli.db.Where("question_id = ?", questionID).Find(&answers).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n===== Answers =====")
	for _, a := range answers {
		fmt.Printf("ID: %d | Answer: %s | Student ID: %d\n",
			a.ID, a.Answer, a.StudentID)
		fmt.Println("------------------------")
	}
}

func (cli *AnswerCLI) addAnswer() {
	var questionID uint
	fmt.Print("Enter Question ID: ")
	fmt.Scan(&questionID)

	var studentID uint
	fmt.Print("Enter Student ID: ")
	fmt.Scan(&studentID)

	var answerText string
	fmt.Print("Enter Answer Text: ")
	fmt.Scanln(&answerText)

	answer := evalModel.Answer{
		QuestionID: questionID,
		StudentID:  studentID,
		Answer:     answerText,
	}

	if err := cli.db.Create(&answer).Error; err != nil {
		fmt.Println("Error creating answer:", err)
		return
	}

	fmt.Println("Created Answer:", answer.ID)
}

func (cli *AnswerCLI) updateAnswer() {
	var id uint
	fmt.Print("Enter Answer ID to update: ")
	fmt.Scan(&id)

	var answer evalModel.Answer
	if err := cli.db.First(&answer, id).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	var answerText string
	fmt.Print("Enter new Answer Text: ")
	fmt.Scanln(&answerText)

	answer.Answer = answerText

	if err := cli.db.Save(&answer).Error; err != nil {
		fmt.Println("Error updating answer:", err)
		return
	}

	fmt.Println("Updated Answer:", answer.ID)
}

func (cli *AnswerCLI) deleteAnswer() {
	var id uint
	fmt.Print("Enter Answer ID to delete: ")
	fmt.Scan(&id)

	if err := cli.db.Delete(&evalModel.Answer{}, id).Error; err != nil {
		fmt.Println("Error deleting answer:", err)
		return
	}

	fmt.Println("Deleted Answer:", id)
}
