// 65070503445
// MEP-1006
package util

import (
	"ModEd/eval/model"
	"os"

	"github.com/gocarina/gocsv"
)

func LoadEvaluationsFromCSV(path string) ([]*model.Evaluation, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var evaluations []*model.Evaluation
	if err := gocsv.UnmarshalFile(file, &evaluations); err != nil {
		return nil, err
	}
	return evaluations, nil
}
