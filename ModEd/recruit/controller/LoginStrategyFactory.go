// MEP-1003 Student Recruitment

package controller

type LoginStrategyFactory interface {
	CreateStrategy(userType string) LoginStrategy
}
