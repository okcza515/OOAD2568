package cli

import (
	"ModEd/eval/controller"
	"fmt"
	"time"
)

type QuizCLI struct {
	controller controller.QuizController
}

func NewQuizCLI(controller controller.QuizController) *QuizCLI {
	return &QuizCLI{controller: controller}
}

func (cli *QuizCLI) Run() {
	for {
		var option int
		fmt.Println("\nQuiz CLI Menu")
		fmt.Println("1. List Quizzes")
		fmt.Println("2. Create Quiz")
		fmt.Println("3. Update Quiz")
		fmt.Println("4. Delete Quiz")
		fmt.Println("5. Exit")
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

func (cli *QuizCLI) listQuizzes() {
	quizzes, err := cli.controller.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, q := range quizzes {
		fmt.Printf("ID: %d | Title: %s | Released: %t\n", q.ID, q.Title, q.Released)
	}
}

func (cli *QuizCLI) createQuiz() {
	var input controller.QuizInput
	fmt.Print("Enter Title: ")
	fmt.Scanln(&input.Title)
	fmt.Print("Enter Description: ")
	fmt.Scanln(&input.Description)
	input.Released = true
	input.QuizStart = time.Now()
	input.QuizEnd = time.Now().Add(1 * time.Hour)
	input.Status = "Scheduled"

	quiz, err := cli.controller.Create(input)
	if err != nil {
		fmt.Println("Error creating quiz:", err)
		return
	}
	fmt.Println("Created Quiz:", quiz.ID)
}

func (cli *QuizCLI) updateQuiz() {
	var id uint
	fmt.Print("Enter Quiz ID to update: ")
	fmt.Scan(&id)

	var input controller.QuizInput
	fmt.Print("Enter new Title: ")
	fmt.Scanln(&input.Title)
	fmt.Print("Enter new Description: ")
	fmt.Scanln(&input.Description)
	input.Released = true
	input.QuizStart = time.Now()
	input.QuizEnd = time.Now().Add(1 * time.Hour)
	input.Status = "Updated"

	quiz, err := cli.controller.Update(id, input)
	if err != nil {
		fmt.Println("Error updating quiz:", err)
		return
	}
	fmt.Println("Updated Quiz:", quiz.ID)
}

func (cli *QuizCLI) deleteQuiz() {
	var id uint
	fmt.Print("Enter Quiz ID to delete: ")
	fmt.Scan(&id)

	if err := cli.controller.Delete(id); err != nil {
		fmt.Println("Error deleting quiz:", err)
		return
	}
	fmt.Println("Deleted Quiz:", id)
}
