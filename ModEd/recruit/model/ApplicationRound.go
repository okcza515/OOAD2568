// MEP-1003 Student Recruitment
package model

type ApplicationRound struct {
	RoundID   uint   `gorm:"primaryKey" csv:"round_id" json:"round_id"`
	RoundName string `csv:"round_name" json:"round_name"`
}

func (i *ApplicationRound) GetID() uint {
	return i.RoundID
}
func (i *ApplicationRound) FromCSV(csvData string) error {
	return nil
}
func (i *ApplicationRound) ToCSVRow() string {
	return ""
}
func (i *ApplicationRound) FromJSON(jsonData string) error {
	return nil
}
func (i *ApplicationRound) ToJSON() string {
	return ""
}
func (i *ApplicationRound) Validate() error {
	return nil
}
func (i *ApplicationRound) ToString() string {
	return ""
}
