package HealthCheckController

import (
	"boilerplate-go/Config/DTO/ConfigStructs/Utilities"
	"boilerplate-go/Libraries/Helpers"

	"github.com/gin-gonic/gin"
)

type HealthCheckInterface interface {
	Ping(g *gin.Context)
}

func NewController(u ConfigStructs.Utilities) HealthCheckInterface {
	return &healthCheck{u}
}

func (h healthCheck) Ping(g *gin.Context) {
	// example send email
	//mail := h.Email.Set([]string{"yourmail@gmail.com"}, "testing", "<h1>Hello from Golang</h1>")
	//go func(ctx context.Context) {
	//	if err := h.Email.Send(ctx, mail); err != nil {
	//		Config.AppLogger.LogErrorWithMark(g, err)
	//	}
	//}(g.Request.Context())

	data := make(map[string]interface{})
	data["message"] = "pong"

	Helpers.HttpResponseSuccess(g, data)

	return
}
