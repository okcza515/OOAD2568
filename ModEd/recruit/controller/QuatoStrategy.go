// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"fmt"
)

type QuotaStrategy struct{}

func (q *QuotaStrategy) ApplyForm(applicant *model.Applicant) error {
	fmt.Println("Enter Math Grade: ")
	fmt.Scan(&applicant.MathGrade)

	fmt.Println("Enter Science Grade: ")
	fmt.Scan(&applicant.ScienceGrade)

	fmt.Println("Enter English Grade: ")
	fmt.Scan(&applicant.EnglishGrade)

	return nil
}
