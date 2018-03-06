package main

import (
	"tw.ntust.dripmonitor/logger/helpers"
	"tw.ntust.dripmonitor/logger/servers"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"os"
)

var config *helpers.Configuration
var mysqlConn *helpers.MySQLConn

func main() {
	var err error

	// App initialization
	initConsoleLogger()
	helpers.InitializePaths()
	defer log.Infoln("Done.")

	// Load configurations
	config, err = helpers.NewConfiguration()
	if err != nil {
		panic("Failed to load app configuration: " + err.Error())
	}

	// Create database connection
	mysqlConn, err = helpers.NewMySQLConn(config)
	if err != nil {
		panic(err)
	}
	defer mysqlConn.Close()

	// Start servers
	servers.InitializeTCPStream(config, mysqlConn) // tcp stream server
	servers.NewIrisApplication(config, mysqlConn)  // http API server

	log.Infoln("Cleaning up...")
}

func initConsoleLogger() {
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true, FullTimestamp: true, TimestampFormat: helpers.TimeFormat})
	log.SetLevel(logrus.InfoLevel)
	log.SetOutput(os.Stdout)
}
