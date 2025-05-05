// controller/default_login_strategy_factory.go

package controller

import "gorm.io/gorm"

type DefaultLoginStrategyFactory struct {
	DB *gorm.DB
}

func (f *DefaultLoginStrategyFactory) CreateStrategy(userType string) LoginStrategy {
	switch userType {
	case "admin":
		return &AdminLoginStrategy{DB: f.DB}
	case "instructor":
		return &UserIDLoginStrategy{DB: f.DB}
	case "user":
		return &UserIDLoginStrategy{DB: f.DB}
	default:
		return nil
	}
}
