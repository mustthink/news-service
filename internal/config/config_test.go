package config

import (
	"flag"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// Change working directory to the root of the project
func init() {
	os.Chdir("../..")
}

// cleanFlags resets the command line flags to avoid conflicts between tests
func cleanFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}

func TestFetchConfigPath(t *testing.T) {
	t.Run("config path from command line arguments", func(t *testing.T) {
		os.Args = []string{"cmd", "-config=testdata/config.yml"}
		assert.Equal(t, "testdata/config.yml", fetchConfigPath())
	})

	cleanFlags()
	t.Run("config path from environment variable", func(t *testing.T) {
		os.Args = []string{"cmd"}
		os.Setenv("CONFIG_PATH", "testdata/config.yml")
		assert.Equal(t, "testdata/config.yml", fetchConfigPath())
		os.Unsetenv("CONFIG_PATH")
	})
}

func TestReadConfig(t *testing.T) {
	t.Run("config file exists", func(t *testing.T) {
		cfg, err := readConfig("config/local.yaml")
		assert.NoError(t, err)
		assert.NotNil(t, cfg)
	})

	t.Run("config file does not exist", func(t *testing.T) {
		_, err := readConfig("testdata/nonexistent.yml")
		assert.Error(t, err)
	})
}

func TestMustLoad(t *testing.T) {
	cleanFlags()
	t.Run("config path is not empty", func(t *testing.T) {
		os.Args = []string{"cmd", "-config=config/local.yaml"}
		assert.NotPanics(t, func() { MustLoad() })
	})

	t.Run("config path is empty", func(t *testing.T) {
		os.Args = []string{"cmd"}
		assert.Panics(t, func() { MustLoad() })
	})
}

func TestLogLevel(t *testing.T) {
	t.Run("debug level", func(t *testing.T) {
		cfg := &Config{Env: "local"}
		assert.Equal(t, logrus.DebugLevel, cfg.LogLevel())
	})

	t.Run("info level", func(t *testing.T) {
		cfg := &Config{Env: "prod"}
		assert.Equal(t, logrus.InfoLevel, cfg.LogLevel())
	})
}
