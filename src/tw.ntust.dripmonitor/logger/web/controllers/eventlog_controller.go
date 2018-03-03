package controllers

import (
	"github.com/kataras/iris"
	"fmt"
	"tw.ntust.dripmonitor/logger/datamodels"
	"tw.ntust.dripmonitor/logger/dao"
)

type EventLogController struct{
	Ctx iris.Context
	EventLogDAO *dao.EventLogDAO
}

func (c *EventLogController) Get() {
	c.Ctx.Writef("EventLogController")
}

func (c *EventLogController) Post() {
	param := datamodels.EventRecord{}

	err := c.Ctx.ReadForm(&param)
	if err != nil {
		fmt.Println(param)
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return
	}

	c.EventLogDAO.InsertRecord(&param)
	c.Ctx.JSON(iris.Map{"proc_status": 1})
}
