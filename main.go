package main

import (
	"lucassantoss1701/clean/src/infrastructure/config"
	"lucassantoss1701/clean/src/infrastructure/router"
)

var applicationContext *config.ApplicationContext

func init() {
	applicationContext = config.NewApplicationContext()
}

func main() {
	router.SetupHTTP(applicationContext).Run()
}
