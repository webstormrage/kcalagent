package main

import (
	"ai-kcal-agent/pkg/appContext"
)

func main() {
	err := appContext.Init()
	ctx := appContext.Get()
	if err != nil {
		panic(err)
	}
	if ctx.ServerMode == "enabled" {
		runServer()
	} else {
		runCli()
	}
}
