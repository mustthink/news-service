package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Env string   `yaml:"env" env-default:"local"`
	Srv Server   `yaml:"server"`
	DB  Database `yaml:"db" env-required:"true"`
}

// fetchConfigPath fetches the path to the config file
// from the command line arguments
// or environment variables
func fetchConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "", "path to config path")
	flag.Parse()

	if path == "" {
		return os.Getenv("CONFIG_PATH")
	}

	return path
}

// readConfig reads the config from the file
func readConfig(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file doesn't exist: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("couldn't read config with error: " + err.Error())
	}

	return &cfg, nil
}

// MustLoad read config and panic if it fails
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	cfg, err := readConfig(path)
	if err != nil {
		panic(err)
	}
	return cfg
}

// LogLevel returns the log level based on the environment
func (c *Config) LogLevel() logrus.Level {
	switch c.Env {
	case "local", "dev", "debug", "test":
		return logrus.DebugLevel
	case "prod", "production":
		fallthrough
	default:
		return logrus.InfoLevel
	}
}
