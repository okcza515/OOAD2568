package main

import (
	controller "ModEd/project/controller"
	model "ModEd/project/model"
	"flag"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Project Evaluation")
	var (
		database string
		path     string
	)
	flag.StringVar(&database, "database", "data/ModEd.bin", "Path of SQLite Database.")
	flag.StringVar(&path, "path", "", "Path to CSV or JSON for student registration.")
	flag.Parse()

	// test senior project evaluation
	var testEvaluation = model.Evaluation{
		TaskID:         uuid.New(),
		GroupID:        uuid.New(),
		AssignmentType: "report",
		Score:          0,
		Comment:        "",
	}

	evaluationController := controller.ProjectEvaluationController{}
	score, comment, err := evaluationController.EvaluateTask(&testEvaluation)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println("------Evaluation Result------")
	fmt.Printf("Score: %f\nComment: %s\n", score, comment)
}
