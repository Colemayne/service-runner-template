package commands

import (
	"path/filepath"

	"../common"
)

func getDefaultConfigFile() string {
	return filepath.Join(getDefaultConfigDirectory(), "tasker.json")
}

type configOptions struct {
	config *common.Config

	ConfigFile string `short:"c" long:"config" env:"CONFIG_FILE" description:"Config file"`
}
