package Controller

import (
	"boilerplate-go/Config/DTO"
	"boilerplate-go/Controller/HealthCheckController"
)

type Controller struct {
	HealthCheck HealthCheckController.HealthCheckInterface
}

func InitControllerApi(u DTO.Utilities) Controller {
	return Controller{
		HealthCheck: HealthCheckController.NewController(u),
	}
}
