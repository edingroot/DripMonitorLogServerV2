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

// Insert a event record
func (c *EventLogController) Post() {
	param := datamodels.EventLog{}
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

// Get: /eventlog/adapter/<AdapterMAC>/<Route>
func (c *EventLogController) GetAdapterBy(adapterMAC string, route string) {
	switch route {
	case "need_restart":
		c.adapterNeedRestart(adapterMAC)

	default:
		c.Ctx.StatusCode(iris.StatusBadRequest)
	}
}

// Get: /eventlog/adapter/<AdapterMAC>/need_restart
func (c *EventLogController) adapterNeedRestart(adapterMAC string) {
	const LookupCount = 10
	response := make(map[string]interface{})
	response["proc_status"] = 1
	response["need_restart"] = false

	btConnectRecords := c.EventLogDAO.GetAdapterConnectsAfterRestart(adapterMAC, LookupCount)

	// Check if event code of all record is 30
	log.Debugf("[isAdapterNeedRestart] %s - count{30,31}=%d", adapterMAC, len(*btConnectRecords))
	if len(*btConnectRecords) < LookupCount {
		c.Ctx.JSON(response)
		return
	} else {
		for _, record := range *btConnectRecords {
			log.Debugf("[isAdapterNeedRestart] %s - %d", adapterMAC, record.EventCode)
			if record.EventCode != 30 {
				c.Ctx.JSON(response)
				return
			}
		}

		response["need_restart"] = true
		c.Ctx.JSON(response)
	}
}
