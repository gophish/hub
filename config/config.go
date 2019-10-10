package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// LoggingConfig represents configuration details for Hub logging.
type LoggingConfig struct {
	Filename string `json:"filename"`
}

type ServerConfig struct {
	ListenAddress string `yaml:"listenAddress"`
}

type Config struct {
	Repositories []string       `yaml:"repositories"`
	Logging      *LoggingConfig `yaml:"logging"`
	Server       *ServerConfig  `yaml:"server"`
}

// LoadConfig loads the configuration from the specified filepath
func LoadConfig(filepath string) (*Config, error) {
	// Get the config file
	configFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
