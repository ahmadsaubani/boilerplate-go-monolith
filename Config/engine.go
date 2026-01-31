package Config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

const (
	ENVIRONMENT_PATH = "../Environment/"
	REDIS            = "redis"
	POSTGRES         = "postgres"
)

type envFile []byte

func GetEnvironment(env string) Config {
	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(path.Dir(filename), ENVIRONMENT_PATH+env+".yml")
	fmt.Println(envPath)
	_, err := os.Stat(envPath)
	if err != nil {
		log.Println(err.Error())
		panic(err)
		return nil
	}
	content, err := ioutil.ReadFile(envPath)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	var config envFile = content
	return config
}

func (e envFile) LoadConfig() *ConfigSetting {
	var config EnvironmentConfig

	err := yaml.Unmarshal([]byte(string(e)), &config)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	if config.App.Debug == false {
		log.SetOutput(ioutil.Discard)
	}
	log.Println("Environment Config load successfully!")
	return &ConfigSetting{&config, nil, &config.App, &config}
}

func (e *EnvironmentConfig) BuildConnection() DbConInterface {
	var connectionPool connectionPool = &DatabaseConfig{}
	var dbCon DbCon
	for i := 0; i < len(e.Databases); i++ {
		connectionPool = &e.Databases[i]
		switch e.Databases[i].Engine {
		case POSTGRES:
			con := sql.DB{}
			log.Println("ENGINE " + POSTGRES + " start....")
			connectionPool.PostgresConnectionPool(&con)
			dbCon.setSqlConnection(DbSqlConfigName(e.Databases[i].Connection), &con)
		case REDIS:
			con := redis.Client{}
			log.Println("ENGINE " + REDIS + " start....")
			connectionPool.RedisConnectionPool(&con)
			dbCon.setRedisConnection(dbRedisConfigName(e.Databases[i].Connection), &con)
		}
	}
	return &dbCon
}

func (app *AppConfig) Run(route *gin.Engine) {
	srv := app.buildServer(route)

	go app.startServer(srv)

	app.waitForShutdown(srv)
}

func (app *AppConfig) buildServer(route *gin.Engine) *http.Server {
	address := app.Host + ":" + app.Port

	return &http.Server{
		Addr:         address,
		Handler:      route,
		ReadTimeout:  app.Api.ReadTimeout * time.Second,
		WriteTimeout: app.Api.WriteTimeout * time.Second,
		IdleTimeout:  app.Api.IdleTimeout * time.Second,
	}
}

func (app *AppConfig) startServer(srv *http.Server) {
	if app.Service == "https" {
		app.runWithHttps(srv)
		return
	}
	app.runWithHttp(srv)
}

func (app *AppConfig) runWithHttp(srv *http.Server) {
	log.Println("HTTP service running ...")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server error: %v", err)
	}
}

func (app *AppConfig) runWithHttps(srv *http.Server) {
	log.Println("HTTPS service running ...")

	cert, key := app.getTLSFiles()

	if err := srv.ListenAndServeTLS(cert, key); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTPS server error: %v", err)
	}
}

func (app *AppConfig) getTLSFiles() (string, string) {
	_, filename, _, _ := runtime.Caller(0)
	basePath := path.Dir(filename)

	cert := path.Join(basePath, "../Infrastructures/certificate", app.Certificate)
	key := path.Join(basePath, "../Infrastructures/certificate", app.Pem_key)

	return cert, key
}

func (app *AppConfig) waitForShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
