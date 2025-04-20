// model/admin.go
package model

type Admin struct {
	AdminID  uint   `gorm:"primaryKey" csv:"admin_id"`
	Username string `csv:"username"`
	Password string `csv:"password"`
}
