package Middlewares

import (
	"boilerplate-go/Config"

	ginmiddleware "github.com/ahmadsaubani/go-logging-lib/middleware"
	"github.com/gin-gonic/gin"
)

// Middleware returns the main middleware that injects request metadata
func Middleware() gin.HandlerFunc {
	return ginmiddleware.GinMiddleware(Config.AppLogger)
}

// GinLogger returns the access logger middleware
func GinLogger() gin.HandlerFunc {
	return ginmiddleware.GinLogger(Config.AppLogger)
}

// HTTPErrorLogger returns the HTTP error logger middleware
func HTTPErrorLogger() gin.HandlerFunc {
	return ginmiddleware.GinHTTPErrorLogger(Config.AppLogger)
}

// RecoveryWithLogger returns the panic recovery middleware with logging
func RecoveryWithLogger() gin.HandlerFunc {
	return ginmiddleware.GinRecovery(Config.AppLogger)
}
