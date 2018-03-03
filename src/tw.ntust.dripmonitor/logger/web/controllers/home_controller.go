package controllers

import (
	"github.com/kataras/iris"
)

type HomeController struct{
	Ctx iris.Context
}

func (c *HomeController) Get() {
	c.Ctx.Writef("DripMonitorLogServerV2")
}
