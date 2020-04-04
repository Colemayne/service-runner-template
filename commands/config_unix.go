package commands

import (
	"os"

	"../helpers"
)

func getDefaultConfigDirectory() string {
	if os.Getuid() == 0 {
		return "/etc/tasker"
	} else if currentDir := helpers.GetCurrentWorkingDirectory(); currentDir != "" {
		return currentDir
	}
	panic("Cannot get default config file location")
}
