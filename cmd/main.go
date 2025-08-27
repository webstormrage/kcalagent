package main

import (
	"ai-kcal-agent/pkg/appContext"
)

func main() {
	err := appContext.Init()
	if err != nil {
		panic(err)
	}
	runServer()
}
