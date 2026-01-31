package Routes

import (
	"boilerplate-go/Controller"
	"boilerplate-go/Routes/groups"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Controller Controller.Controller
	Gin        *gin.Engine
}

func (app *Routes) CollectRoutes() *gin.Engine {

	appRoute := app.Gin
	//apiGroup := groups.RoutesGroupCollection
	_ = groups.RoutesGroupCollection
	appRoute.GET("/ping", app.Controller.HealthCheck.Ping)
	return appRoute
}
