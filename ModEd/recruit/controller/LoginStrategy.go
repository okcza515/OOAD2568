package controller

type LoginStrategy interface {
	CheckUsername(username string) (bool, error)
	CheckUsernameAndPassword(username, password string) (bool, error)
	CheckID(id string) (bool, error)
}
