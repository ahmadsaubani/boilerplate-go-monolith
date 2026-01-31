package Config

import (
	"bytes"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type connectionPool interface {
	PostgresConnectionPool(Connection interface{})
	RedisConnectionPool(Connection interface{})
}

func (env *DatabaseConfig) PostgresConnectionPool(Connection interface{}) {
	Con := Connection.(*sql.DB)
	var buffer bytes.Buffer
	buffer.WriteString("postgres://")
	buffer.WriteString(env.Username + ":" + env.Password)
	buffer.WriteString("@")
	buffer.WriteString(env.Host + ":" + env.Port + "/")
	buffer.WriteString(env.Name)
	buffer.WriteString("?sslmode=disable")
	connection_string := buffer.String()
	Connection, err := sql.Open(POSTGRES, connection_string)
	if err != nil {
		//panic err
		panic(err.Error())
		return
	}
	Connection.(*sql.DB).SetMaxOpenConns(env.Maximum_connection)
	Connection.(*sql.DB).SetConnMaxIdleTime(env.MaximumIdleTime * time.Second)
	*Con = *Connection.(*sql.DB)
	err = Con.Ping()
	if err != nil {
		log.Print(err.Error())
		panic(err.Error())
		return
	}
	return
}

func (env *DatabaseConfig) RedisConnectionPool(Connection interface{}) {
	var buffer bytes.Buffer
	ctx := context.Background()
	Con := Connection.(*redis.Client)
	buffer.WriteString(env.Host + ":" + env.Port)
	connectionString := buffer.String()

	Connection = redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: env.Password,
		DB:       0,
	})

	*Con = *Connection.(*redis.Client)

	_, err := Con.Ping(ctx).Result()
	if err != nil {
		log.Print(err.Error())
		panic(err.Error())
		return
	}
	return
}