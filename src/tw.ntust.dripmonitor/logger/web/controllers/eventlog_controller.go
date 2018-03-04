package controllers

import (
	"github.com/kataras/iris"
	"fmt"
	"tw.ntust.dripmonitor/logger/datamodels"
	"tw.ntust.dripmonitor/logger/dao"
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
		fmt.Printf("%s Problem reading form: %s\n", LogTagEC, err.Error())
		fmt.Printf("%s Request payload: %s\n", LogTagEC, form)
		fmt.Printf("%s Continue processing\n", LogTagEC)
	}

	// Check necessary inputs
	if form.Get("mac_adapter") == "" || form.Get("event_code") == "" {
		fmt.Printf("%s Bad request, payload: %s\n", LogTagEC, form)
		c.Ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	c.EventLogDAO.InsertRecord(&param)
	c.Ctx.JSON(iris.Map{"proc_status": 1})
}
