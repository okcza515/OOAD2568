package controller

type LoginStrategy interface {
	CheckUsername(username string) (bool, error)
}
