// MEP-1007
package controller

type IGradingStrategy interface {
	Grade(submissionID uint) (float64, error)
}
