package controllers

import (
	"github.com/kataras/iris"
	"tw.ntust.dripmonitor/logger/datamodels"
	"tw.ntust.dripmonitor/logger/dao"
	log "github.com/sirupsen/logrus"
	"tw.ntust.dripmonitor/logger/helpers"
)

const LogTagEC = "[EventLogController]"

type EventLogController struct{
	Ctx iris.Context
	EventLogDAO *dao.EventLogDAO
}

func (c *EventLogController) Get() {
	c.Ctx.Writef("EventLogController")
}

func (c *EventLogController) Post() {
	param := datamodels.EventRecord{}
	form := &c.Ctx.Request().Form

	// Read form inputs
	err := c.Ctx.ReadForm(&param) // form (map) got filled after this
	if err != nil {
		log.Warnf("%s Problem reading form: %s", LogTagEC, err.Error())
		log.Warnf("%s Request payload: %s", LogTagEC, form)
		log.Warnf("%s Continue processing", LogTagEC)
	}

	// Check necessary inputs
	if form.Get("mac_adapter") == "" || form.Get("event_code") == "" {
		log.Errorf("%s Bad request, payload: %s", LogTagEC, form)
		c.Ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	// Log IP & port
	param.SrcIP, param.SrcPort = helpers.GetIpPortFromAddr(c.Ctx.Request().RemoteAddr)

	c.EventLogDAO.InsertRecord(&param)
	c.Ctx.JSON(iris.Map{"proc_status": 1})
}
