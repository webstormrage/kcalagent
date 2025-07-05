package appContext

import (
	"github.com/joho/godotenv"
	"os"
)

type Context struct {
	DataSourceName string
	GenAiApiKey string
}

var context *Context

func Init()error{
    err := godotenv.Load()
	if err != nil {
		return err
	}
	context = &Context{
		DataSourceName: os.Getenv("DATA_SOURCE_NAME"),
		GenAiApiKey: os.Getenv("GEN_AI_API_KEY"),
	}
	return nil
}

func Get()Context{
	return *context
}