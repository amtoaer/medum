package path

import (
	"fmt"
	"medum/text"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

//GetPath return the location of ".medum" folder
func GetPath() string {
	path, err := homedir.Dir()
	if err != nil {
		fmt.Printf(text.HomedirError, err)
		os.Exit(1)
	}
	return filepath.Join(path, ".medum")
}

//GetConfigPath get the location of config file
func GetConfigPath() string {
	return filepath.Join(GetPath(), "config.json")
}

// GetDataPath get the location of data
func GetDataPath() string {
	return filepath.Join(GetPath(), "data.db")
}
