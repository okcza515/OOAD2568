// MEP-1003 Student Recruitment
package model

// ApplicationRound defines the details of an application round.
type ApplicationRound struct {
	RoundID   uint   `gorm:"primaryKey" csv:"round_id" json:"round_id"`
	RoundName string `csv:"round_name" json:"round_name"`
}
