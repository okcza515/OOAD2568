// MEP-1003 Student Recruitment
package controller

type LoginRequest struct {
	ID       string
	Username string
	Password string
}
type LoginStrategy interface {
	ApplyLogin(req LoginRequest, model interface{}) (bool, error)
}
