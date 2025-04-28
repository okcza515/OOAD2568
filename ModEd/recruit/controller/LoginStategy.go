package controller

// LoginStrategy interface defines methods for login strategies
type LoginStrategy interface {
	CheckUsername(username string) (bool, error)
	CheckUsernameAndPassword(username, password string) (bool, error)
}
