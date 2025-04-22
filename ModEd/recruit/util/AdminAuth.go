package util

import (
	"ModEd/recruit/model"
)

func ValidateAdminLoginFromCSV(username, password, filePath string) bool {
	admins, err := ReadOnlyFromCSVOrJSON[model.Admin](filePath)
	if err != nil {
		return false
	}

	for _, admin := range admins {
		if admin.Username == username && admin.Password == password {
			return true
		}
	}
	return false
}
