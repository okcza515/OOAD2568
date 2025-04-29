package controller

import "gorm.io/gorm"

type LoginController struct {
	Strategy LoginStrategy
}

func NewLoginStrategy(userType string, db *gorm.DB) LoginStrategy {
	switch userType {
	case "admin":
		return &AdminLoginStrategy{DB: db}
	case "instructor":
		return &UserIDLoginStrategy{DB: db}
	case "user":
		return &UserIDLoginStrategy{DB: db}
	default:
		return nil
	}
}

func (c *LoginController) SetStrategy(strategy LoginStrategy) {
	c.Strategy = strategy
}

func (c *LoginController) ExecuteLogin(req LoginRequest, model interface{}) (bool, error) {
	return c.Strategy.ApplyLogin(req, model)
}
