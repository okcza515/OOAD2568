package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type LoginController struct {
	strategy LoginStrategy
	Base     *core.BaseController[*model.Admin]
	DB       *gorm.DB
}

func CreateLoginController(strategy LoginStrategy) *LoginController {
	return &LoginController{strategy: strategy}
}

func (c *LoginController) CheckUsername(username string) (bool, error) {
	return c.strategy.CheckUsername(username)
}
