// MEP-1003 Student Recruitment
package model

// ApplicationRound defines the details of an application round.
type ApplicationRound struct {
	RoundID   uint `gorm:"primaryKey"`
	RoundName string
}
