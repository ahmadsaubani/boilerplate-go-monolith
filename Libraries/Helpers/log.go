package Helpers

import (
	"boilerplate-go/Config"
	"context"

	logging "github.com/ahmadsaubani/go-logging-lib"
)

func LogError(ctx context.Context, err error) {
	if err == nil {
		return
	}

	if Config.AppLogger != nil {
		// Log both human-readable and JSON format
		Config.AppLogger.Error(ctx, err)
		Config.AppLogger.ErrorLoki(ctx, logging.LevelError, err)
	}
}

func LogErrorCritical(ctx context.Context, err error) {
	if err == nil {
		return
	}

	if Config.AppLogger != nil {
		// Log both human-readable and critical level JSON format
		Config.AppLogger.Error(ctx, err)
		Config.AppLogger.ErrorLoki(ctx, logging.LevelCritical, err)
	}
}

func LogInfo(message string) {
	if Config.AppLogger != nil {
		Config.AppLogger.Info(message)
	}
}

func LogWarn(ctx context.Context, err error) {
	if err == nil {
		return
	}

	if Config.AppLogger != nil {
		Config.AppLogger.ErrorLoki(ctx, logging.LevelWarn, err)
	}
}
