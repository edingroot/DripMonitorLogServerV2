package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/middleware/logger"
	"tw.ntust.dripmonitor/logger/web/controllers"
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

	// Initialize Iris app
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(logger.New())

	registerRoutes(app)

	// Start the web server
	app.Run(
		iris.Addr(fmt.Sprintf("%s:%d", config.ListenHost, config.ListenPort)),
		iris.WithoutVersionChecker, // disables updates
		iris.WithoutServerError(iris.ErrServerClosed), // skip err server closed when CTRL/CMD+C pressed
		iris.WithOptimizations, // enables faster json serialization and more
	)
}

func registerRoutes(app *iris.Application) {
	// Route: /
	mvc.New(app.Party("/")).Handle(new(controllers.HomeController))
}
