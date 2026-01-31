package Controller

import (
	"boilerplate-go/Config/DTO/ConfigStructs/Utilities"
	"boilerplate-go/Controller/Articles"
	"boilerplate-go/Controller/HealthCheckController"
)

type Controller struct {
	HealthCheck HealthCheckController.HealthCheckInterface
	Article     Articles.ArticleControllerInterface
}

func InitControllerApi(u ConfigStructs.Utilities) Controller {
	return Controller{
		HealthCheck: HealthCheckController.NewController(u),
		Article:     Articles.NewController(u),
	}
}
