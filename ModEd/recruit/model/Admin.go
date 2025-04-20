// MEP-1003 Student Recruitment
package model

type Admin struct {
	AdminID  uint   `gorm:"primaryKey" csv:"admin_id"`
	Username string `csv:"username"`
	Password string `csv:"password"`
}

func (i *Admin) GetID() uint {
	return i.AdminID
}
func (i *Admin) FromCSV(csvData string) error {
	return nil
}
func (i *Admin) ToCSVRow() string {
	return ""
}
func (i *Admin) FromJSON(jsonData string) error {
	return nil
}
func (i *Admin) ToJSON() string {
	return ""
}
func (i *Admin) Validate() error {
	return nil
}
func (i *Admin) ToString() string {
	return ""
}
