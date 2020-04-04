package helpers

import "os"

// GetCurrentWorkingDirectory says it in the name
func GetCurrentWorkingDirectory() string {
	dir, err := os.Getwd()
	if err == nil {
		return dir
	}
	return ""
}
