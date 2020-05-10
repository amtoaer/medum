package path

import (
	"fmt"
	"medum/text"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func getPath() string {
	path, err := homedir.Dir()
	if err != nil {
		fmt.Printf(text.HomedirError, err)
		os.Exit(1)
	}
	return path
}

//GetConfigPath get the location of config file
func GetConfigPath() string {
	return filepath.Join(getPath(), ".medum/config.json")
}

// GetDataPath get the location of data
func GetDataPath() string {
	return filepath.Join(getPath(), ".medum/data.db")
}
