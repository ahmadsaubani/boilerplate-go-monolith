package Config

import (
	"fmt"

	logging "github.com/ahmadsaubani/go-logging-lib"
)

type LoggingConfig struct {
	ServiceName    string `yaml:"service_name"`
	LogPath        string `yaml:"log_path"`
	EnableStdout   bool   `yaml:"enable_stdout"`
	EnableFile     bool   `yaml:"enable_file"`
	EnableLoki     bool   `yaml:"enable_loki"`
	EnableRotation bool   `yaml:"enable_rotation"`
}

var AppLogger *logging.Logger

func InitLoggerFromConfig(loggingConfig LoggingConfig) {
	config := &logging.Config{
		ServiceName:    loggingConfig.ServiceName,
		LogPath:        loggingConfig.LogPath,
		EnableStdout:   loggingConfig.EnableStdout,
		EnableFile:     loggingConfig.EnableFile,
		EnableLoki:     loggingConfig.EnableLoki,
		EnableRotation: loggingConfig.EnableRotation,
	}

	var err error
	AppLogger, err = logging.New(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}

	AppLogger.Info("🚀 Logger initialized from YAML configuration")
}
