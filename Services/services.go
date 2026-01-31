package Services

import (
	"boilerplate-go/Config"
	"boilerplate-go/Config/DTO/ConfigStructs/ModuleConfigs"
	utilities "boilerplate-go/Config/DTO/ConfigStructs/Utilities"
	"boilerplate-go/Controller"
	"boilerplate-go/Middlewares"
	"boilerplate-go/Modules"
	"boilerplate-go/Modules/Articles"
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
	EnvDocker      AppEnv = "Docker"
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

	//env := resolveEnv(*AppEnvFlag)
	//routesConfig = buildGinEngine(env)
	//
	//routesConfig.Use(
	//	Middlewares.Middleware(),
	//	Middlewares.GinLogger(),
	//	Middlewares.HTTPErrorLogger(),
	//	Middlewares.RecoveryWithLogger(),
	//)

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
	case "docker":
		return EnvDocker
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
		gin.SetMode(gin.ReleaseMode)
		return gin.New()
	}
}

func AppInitialization() {
	newConfig := Config.GetEnvironment(*AppEnvFlag).LoadConfig()
	Config.InitLoggerFromConfig(newConfig.Environment.Logging)

	env := resolveEnv(*AppEnvFlag)
	routesConfig = buildGinEngine(env)

	routesConfig.Use(
		Middlewares.Middleware(),
		Middlewares.GinLogger(),
		Middlewares.HTTPErrorLogger(),
		Middlewares.RecoveryWithLogger(),
	)

	connection := newConfig.Database.BuildConnection()
	service := serviceInit(newConfig.Environment)
	_ = Repositories.InitRepo(connection)
	ModulesConfig := ConfigStructs.ModuleConfigs{
		Repo: Repositories.InitRepo(connection),
	}
	newsConfig := newConfig.Environment.News
	utilities := utilities.Utilities{
		Email: service.Email,
		Modules: Modules.Modules{
			ArticleModule: Articles.NewModule(ModulesConfig, newsConfig),
		},
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
