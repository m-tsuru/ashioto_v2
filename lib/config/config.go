package config

import (
	"errors"
	"github.com/m-tsuru/ashioto-api/structs"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadConfig(filePath *string) (*structs.Config, error) {
	// Load configuration from the specified file
	var config structs.Config
	configFile, err := os.ReadFile(*filePath)
	if err != nil {
		return nil, errors.New("Failed to read configuration file: " + err.Error())
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, errors.New("Failed to parse configuration file: " + err.Error())
	}

	return &config, nil
}
