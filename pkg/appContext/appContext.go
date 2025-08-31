package appContext

import (
	"log"
	"os"
)

type Context struct {
	DataSourceName string
	JwtSecret      string
	ServerPort     string
	ServerMode     string
	Logger         *log.Logger
}

var context *Context

func Init() {
	context = &Context{
		DataSourceName: os.Getenv("DATA_SOURCE_NAME"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		ServerPort:     os.Getenv("SERVER_PORT"),
		ServerMode:     os.Getenv("SERVER_MODE"),
		Logger:         log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}
}

func Get() Context {
	return *context
}
