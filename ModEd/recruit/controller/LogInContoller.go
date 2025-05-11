// MEP-1003 Student Recruitment
package controller

type LoginController struct {
	Strategy LoginStrategy
}

func (c *LoginController) SetStrategy(strategy LoginStrategy) {
	c.Strategy = strategy
}

func (c *LoginController) ExecuteLogin(req LoginRequest, model interface{}) (bool, error) {
	return c.Strategy.ApplyLogin(req, model)
}
