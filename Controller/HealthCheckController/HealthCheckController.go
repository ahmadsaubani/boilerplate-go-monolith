package HealthCheckController

import (
	"boilerplate-go/Config/DTO"
	"boilerplate-go/Libraries/Helpers"
	"context"

	logging "github.com/ahmadsaubani/go-logging-lib"
	"github.com/gin-gonic/gin"
)

type HealthCheckInterface interface {
	Ping(g *gin.Context)
}

func NewController(u DTO.Utilities) HealthCheckInterface {
	return &healthCheck{u}
}

func (h healthCheck) Ping(g *gin.Context) {

	mail := h.Email.Set([]string{"gshintas40@gmail.com"}, "testing", "<h1>Hello from Golang 🚀</h1>")
	Helpers.LogInfo("testing info")

	//go h.Email.Send(mail)
	go func(ctx context.Context) {
		if err := h.Email.Send(ctx, mail); err != nil {
			logging.MarkErrorLogged(g)
		}
	}(g.Request.Context())

	data := make(map[string]interface{})
	data["message"] = "pong"

	Helpers.HttpResponseSuccess(g, data)

	return
}
