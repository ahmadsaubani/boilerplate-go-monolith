package Config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config interface {
	LoadConfig() *ConfigSetting
}

type Db interface {
	BuildConnection() DbConInterface
}

type ConfigSetting struct {
	Database    Db
	Routes      RouteInterface
	HttpEngine  HttpEngine
	Environment *EnvironmentConfig
}

type HttpEngine interface {
	buildServer(route *gin.Engine) *http.Server
	startServer(srv *http.Server)
	runWithHttp(route *http.Server)
	runWithHttps(route *http.Server)
	getTLSFiles() (string, string)
	waitForShutdown(srv *http.Server)
	Run(route *gin.Engine)
}

type RouteInterface interface {
	SetCors()
	CollectRoutes() *gin.Engine
}
