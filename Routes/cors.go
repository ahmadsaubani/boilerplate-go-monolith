package Routes

import "github.com/gin-contrib/cors"

func (app *Routes) SetCors() {
	route := app.Gin

	cors := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-Requested-With"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	route.Use(cors)
}
