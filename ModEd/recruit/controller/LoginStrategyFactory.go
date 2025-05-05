// controller/login_strategy_factory.go

package controller

type LoginStrategyFactory interface {
	CreateStrategy(userType string) LoginStrategy
}
