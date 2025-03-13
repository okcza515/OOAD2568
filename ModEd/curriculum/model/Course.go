package model

type Course struct {
	Name         string
	Prerequisite []Course
	Optional     bool
}
