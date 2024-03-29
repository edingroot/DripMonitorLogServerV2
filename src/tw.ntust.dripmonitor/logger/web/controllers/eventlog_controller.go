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
	const LookupCount21 = 30
	const LookupCount3031 = 10
	response := make(map[string]interface{})
	response["proc_status"] = 1
	response["need_restart"] = false

	// Check count of event21 which message is not null (cached drip device list is empty)
	emptyDripListCount := c.EventLogDAO.GetAdapterEmptyDripListCount(adapterMAC, LookupCount21)
	log.Debugf("[isAdapterNeedRestart] count{21empty}=%d", emptyDripListCount)
	if emptyDripListCount == 30 {
		// Message field of all records are null
		log.Infof("[isAdapterNeedRestart] %s - YES", adapterMAC)
		response["need_restart"] = true
		c.Ctx.JSON(response)
		return
	}

	// Check if event code of all drip connects record is 30
	btConnectRecords := c.EventLogDAO.GetAdapterDripConnectsAfterRestart(adapterMAC, LookupCount3031)
	log.Debugf("[isAdapterNeedRestart] count{30,31}=%d", len(*btConnectRecords))
	if len(*btConnectRecords) < LookupCount3031 {
		response["need_restart"] = false
		c.Ctx.JSON(response)
		return
	} else {
		for _, record := range *btConnectRecords {
			log.Debugf("[isAdapterNeedRestart] seq - %d", record.EventCode)
			if record.EventCode != 30 {
				response["need_restart"] = false
				c.Ctx.JSON(response)
				return
			}
		}

		log.Infof("[isAdapterNeedRestart] %s - YES", adapterMAC)
		response["need_restart"] = true
		c.Ctx.JSON(response)
	}
}
