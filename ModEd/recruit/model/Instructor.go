package model

type Instructor struct {
	ID         uint `gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Email      string
	Interviews []Interview `gorm:"many2many:instructor_interviews;"`
}
