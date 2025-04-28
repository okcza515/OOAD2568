// MEP-1006

package main

import (
	"ModEd/eval/cli"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"
)

func main() {
	// ดึงข้อมูลจาก CSV   //ModEd/eval/cli/evaluation/
	evals, err := model.LoadEvaluationsFromCSV("ModEd/eval/cli/evaluation/EvalTest.csv")
	if err != nil {
		panic(err)
	}

	// สร้าง Evaluation Controller
	evalController := controller.NewEvaluationController(evals, "ModEd/eval/cli/evaluation/EvalTest.csv")

	// เรียก CLI
	cli.RunEvaluationCLI(evalController)
	fmt.Println("Program exited")
}

// package main

// //"ModEd/eval/cli"
// // "gorm.io/gorm"

// func main()
// 	//fmt.Println("Hello from Main")

// 	//cli.PrintAssignment()
// 	//cli.PrintEvaluation()
// 	//cli.RecordEvaluation()

// 	//cli.RunEvaluationCLI()
// }
