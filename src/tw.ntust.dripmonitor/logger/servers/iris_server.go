package servers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"fmt"
	"github.com/kataras/iris/mvc"
	"tw.ntust.dripmonitor/logger/web/controllers"
	"tw.ntust.dripmonitor/logger/helpers"
	"tw.ntust.dripmonitor/logger/dao"
)

func NewIrisApplication(config *helpers.Configuration, mysqlConn *helpers.MySQLConn) *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(logger.New())

	registerRoutes(app, mysqlConn)

	// Start the web server
	app.Run(
		iris.Addr(fmt.Sprintf("%s:%d", config.HttpListenHost, config.HttpListenPort)),
		iris.WithoutVersionChecker, // disables updates
		iris.WithoutServerError(iris.ErrServerClosed), // skip err server closed when CTRL/CMD+C pressed
		iris.WithOptimizations, // enables faster json serialization and more
	)

	return app
}

func registerRoutes(app *iris.Application, mysqlConn *helpers.MySQLConn) {
	eventLogDAO := dao.NewEventLogDAO(mysqlConn.DB)

	mvc.New(app.Party("/")).
		Handle(new(controllers.HomeController))

	mvc.New(app.Party("/eventlog")).
		Register(eventLogDAO).
		Handle(new(controllers.EventLogController))
}
