package config

import (
	"os"
	"sync"
)

type Config struct {
	Filepath string
}

var (
	once sync.Once
	cfg  *Config
)

// GetConfig returns the config the variables as read from enironment variables
func GetConfig() *Config {

	once.Do(func() {
		cfg = &Config{}
		cfg.SetDefaults()
		cfg.Filepath = os.Getenv("INPUT_FILE_PATH")

		if cfg.Filepath == "" {
			cfg.SetDefaults()
		}
	})

	return cfg
}

// SetDefaults sets the default values for the config
func (c *Config) SetDefaults() {
	c.Filepath = "./input.txt"
}
