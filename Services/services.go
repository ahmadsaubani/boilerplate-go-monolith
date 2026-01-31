package Services

import (
	"boilerplate-go/Config"
	"boilerplate-go/Config/DTO"
	"boilerplate-go/Controller"
	"boilerplate-go/Middlewares"
	"boilerplate-go/Repositories"
	"boilerplate-go/Routes"
	"boilerplate-go/Services/Emails"
	"flag"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	EnvDevelopment AppEnv = "Development"
	EnvProduction  AppEnv = "Production"
	EnvStaging     AppEnv = "Staging"
	EnvTest        AppEnv = "Test"
)

type AppEnv string

var routesConfig *gin.Engine

var AppEnvFlag = flag.String(
	"env",
	"",
	"define environment (Development|Staging|Production|Test)",
)

func init() {
	flag.Parse()

	env := resolveEnv(*AppEnvFlag)
	routesConfig = buildGinEngine(env)

	routesConfig.Use(
		Middlewares.Middleware(),
		Middlewares.GinLogger(),
		Middlewares.HTTPErrorLogger(),
		Middlewares.RecoveryWithLogger(),
	)

}

func resolveEnv(input string) AppEnv {
	env := strings.ToLower(strings.TrimSpace(input))

	if env == "" {
		return EnvDevelopment
	}

	switch env {
	case "prod", "production":
		return EnvProduction
	case "staging", "stage":
		return EnvStaging
	case "test", "testing":
		return EnvTest
	default:
		return EnvDevelopment
	}
}

func buildGinEngine(env AppEnv) *gin.Engine {
	switch env {
	case EnvProduction:
		gin.SetMode(gin.ReleaseMode)
		return gin.New()

	case EnvTest:
		gin.SetMode(gin.TestMode)
		return gin.New()

	default:
		//gin.SetMode(gin.DebugMode)
		gin.SetMode(gin.ReleaseMode)
		//return gin.Default()
		return gin.New()
	}
}

func AppInitialization() {
	newConfig := Config.GetEnvironment(*AppEnvFlag).LoadConfig()
	Config.InitLoggerFromConfig(newConfig.Environment.Logging)

	connection := newConfig.Database.BuildConnection()
	service := serviceInit(newConfig.Environment)
	_ = Repositories.InitRepo(connection)
	_ = DTO.ModuleConfig{
		Repositories.InitRepo(connection),
	}

	utilities := DTO.Utilities{
		Email: service.Email,
	}
	newConfig.Routes = &Routes.Routes{
		Controller: Controller.InitControllerApi(utilities),
		Gin:        routesConfig,
	}

	newConfig.Routes.SetCors()
	routes := newConfig.Routes.CollectRoutes()
	newConfig.HttpEngine.Run(routes)

}

func serviceInit(Env *Config.EnvironmentConfig) service {
	serv := service{
		Email: Emails.EmailSetting{Config: &Env.Email},
	}
	return serv
}
