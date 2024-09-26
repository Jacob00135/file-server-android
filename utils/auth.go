package utils

import (
	"github.com/Jacob00135/file-server-android/database"
)

func CheckUserFilePermission(username, path string) (bool, error) {
	var user uint = 1
	if username != "" {
		var err error
		user, err = database.DB.GetUserPermission(username)
		if err != nil {
			return false, err
		}
	}

	file, err := database.DB.GetFilePermission(path)
	if err != nil {
		return false, err
	}
	return user >= file, nil
}
