package main

import (
	"fmt"
	// "github.com/keiraqz/WWC/config"
	"github.com/keiraqz/WWC/app"
	// "github.com/keiraqz/WWC/logger"
)

// var version string

func main() {
	fmt.Sprintf("Starting WWC application")
	// logger.Info("Starting WWC application")
	// logger.Infof("Running version %s built on %s", version, builtOn)
	// app.ConnectionString = config.DBString()
	// app.MigrationsDirectory = "./db/migrations"
	app.Run()
}
