package main

import (
	"tw.ntust.dripmonitor/logger/helpers"
	"tw.ntust.dripmonitor/logger/servers"
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

	// Start servers
	servers.InitializeTCPStream(config, mysqlConn) // tcp stream server
	servers.NewIrisApplication(config, mysqlConn)  // http API server
}
