//Package config includes functions related with config file, such as read, write...
package config

import (
	"encoding/json"
	"fmt"
	"medum/path"
	"medum/public"
	"medum/text"
	"os"
)

func writeInitConfig(configPath string) {
	// if ".medum" folder isn't exist, create
	tmp := path.GetPath()
	if _, err := os.Stat(tmp); !os.IsExist(err) {
		os.Mkdir(tmp, os.FileMode(0777))
	}
	file, err := os.Create(configPath)
	defer file.Close()
	if err != nil {
		fmt.Printf(text.CreateConfigError, err)
		os.Exit(1)
	}
	conf := public.Configuration{
		NumberColor: "blue",
		EventColor:  "red",
		TimeColor:   "green",
	}
	encoder := json.NewEncoder(file)
	encoder.Encode(conf)
}

// ReadConfig returns a config struct pointer
func ReadConfig() *public.Configuration {
	configPath := path.GetConfigPath()
	// if there wasn't config file,create one and write init config
	if _, err := os.Stat(configPath); err != nil && !os.IsExist(err) {
		writeInitConfig(configPath)
	}
	file, err := os.Open(configPath)
	defer file.Close()
	if err != nil {
		fmt.Printf(text.OpenConfigError, err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	conf := new(public.Configuration)
	decoder.Decode(conf)
	return conf
}
