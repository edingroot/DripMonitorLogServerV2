package main

import (
	"fmt"
	"tw.ntust.dripmonitor/logger/setup/db"
	"tw.ntust.dripmonitor/logger/helpers"
)

var config *helpers.Configuration
var mysqlConn *helpers.MySQLConn

func main() {
	var err error

	// App initialization
	helpers.InitializePaths()

	// Load configurations
	config, err = helpers.NewConfiguration()
	if err != nil {
		panic("Failed to load app configuration: " + err.Error())
	}

	// Create database connection
	mysqlConn, err = helpers.NewMySQLConn(config)
	if err != nil {
		panic("Failed to load app configuration: " + err.Error())
	}
	defer mysqlConn.Close()

	setup()

	fmt.Println("Done.")
}

func setup() {
	// Seed database
	db.SeedDB(mysqlConn.DB)
}
