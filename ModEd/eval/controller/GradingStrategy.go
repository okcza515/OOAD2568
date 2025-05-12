// MEP-1007
package controller

type GradingStrategy interface {
	Grade(submissionID uint) (float64, error)
}
