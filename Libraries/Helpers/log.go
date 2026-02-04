package Helpers

import (
	"boilerplate-go/Config"
	"context"

	logging "github.com/ahmadsaubani/go-logging-lib"
	"github.com/gin-gonic/gin"
)

// LogError logs error to error log and Loki (ERROR level)
func LogError(ctx context.Context, err error) {
	if err == nil {
		return
	}

	if Config.AppLogger != nil {
		Config.AppLogger.Error(ctx, err)
		Config.AppLogger.ErrorLoki(ctx, logging.LevelError, err)
	}
}

// LogErrorCritical logs error to error log and Loki (CRITICAL level)
func LogErrorCritical(ctx context.Context, err error) {
	if err == nil {
		return
	}

	if Config.AppLogger != nil {
		Config.AppLogger.Error(ctx, err)
		Config.AppLogger.ErrorLoki(ctx, logging.LevelCritical, err)
	}
}

// LogErrorWithGin logs error and marks it in Gin context to prevent duplicate logging
// The error will be included in Loki log via GinLogger middleware
func LogErrorWithGin(c *gin.Context, err error) {
	if err == nil {
		return
	}

	if Config.AppLogger != nil {
		// Log to error log
		Config.AppLogger.Error(c.Request.Context(), err)
		// Mark error in Gin context for Loki logging (prevents duplicate)
		logging.SetLoggedError(c, err)
	}
}

// LogInfo logs info message to access log
func LogInfo(message string) {
	if Config.AppLogger != nil {
		Config.AppLogger.Info(message)
	}
}

// LogWarn logs warning level error to Loki
func LogWarn(ctx context.Context, err error) {
	if err == nil {
		return
	}

	if Config.AppLogger != nil {
		Config.AppLogger.ErrorLoki(ctx, logging.LevelWarn, err)
	}
}
