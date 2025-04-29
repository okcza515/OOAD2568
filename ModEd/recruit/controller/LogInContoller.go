package controller

import (
	"gorm.io/gorm"
)

type LoginController struct {
	strategy LoginStrategy
	DB       *gorm.DB
}

func NewLoginController(db *gorm.DB) *LoginController {
	return &LoginController{DB: db}
}

func (c *LoginController) SetStrategyByRole(role string) {
	switch role {
	case "user":
		c.strategy = &UserIDLoginStrategy{DB: c.DB}
	case "instructor":
		c.strategy = &InstructorIDLoginStrategy{DB: c.DB}
	case "admin":
		c.strategy = &UsernamePasswordLoginStrategy{DB: c.DB}
	default:
		panic("invalid role")
	}
}

func (c *LoginController) CheckUsername(username string) (bool, error) {
	return c.strategy.CheckUsername(username)
}

func (c *LoginController) CheckUsernameAndPassword(username, password string) (bool, error) {
	return c.strategy.CheckUsernameAndPassword(username, password)
}

func (c *LoginController) CheckID(id string) (bool, error) {
	return c.strategy.CheckID(id)
}
