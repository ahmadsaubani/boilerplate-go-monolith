package Routes

import (
	"boilerplate-go/Controller"
	"boilerplate-go/Libraries/Helpers"
	"boilerplate-go/Routes/groups"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Controller Controller.Controller
	Gin        *gin.Engine
}

func (app *Routes) CollectRoutes() *gin.Engine {

	appRoute := app.Gin
	apiGroup := groups.RoutesGroupCollection
	appRoute.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		Helpers.HttpResponseError(
			c,
			"internal server error",
			http.StatusInternalServerError,
		)
	}))

	// ✅ 404 handler
	appRoute.NoRoute(func(c *gin.Context) {
		Helpers.HttpResponseError(
			c,
			"route not found",
			http.StatusNotFound,
		)
	})
	appRoute.GET("/ping", app.Controller.HealthCheck.Ping)

	articleGroup := appRoute.Group("/api/v1")
	apiGroup.ArticleGroup.ArticleApiGroup(articleGroup, app.Controller.Article)
	return appRoute
}
