package datamodels

import (
	"time"
	"database/sql"
	"tw.ntust.dripmonitor/logger/helpers"
)

type EventLog struct {
	SN            int32     `json:"sn"`
	EventCode     int       `json:"event_code" form:"event_code"`
	Message       string    `json:"message" form:"message"`
	AdapterMAC    string    `json:"mac_adapter" form:"mac_adapter"`
	DripMAC       string    `json:"mac_drip" form:"mac_drip"`
	SrcIP         string    `json:"src_ip"`
	SrcPort       int32     `json:"src_port"`
	CreatedAt     time.Time `json:"created_at"`
}

type EventLogSQL struct {
	SN            int32
	EventCode     int
	Message       sql.NullString
	AdapterMAC    sql.NullString
	DripMAC       sql.NullString
	SrcIP         sql.NullString
	SrcPort       sql.NullInt64
	CreatedAt     time.Time
}

func (r *EventLog) SQLForm() *EventLogSQL {
	return &EventLogSQL{
		r.SN,
		r.EventCode,
		helpers.StringToNullString(r.Message),
		helpers.StringToNullString(r.AdapterMAC),
		helpers.StringToNullString(r.DripMAC),
		helpers.StringToNullString(r.SrcIP),
		helpers.Int64ToNullInt64(int64(r.SrcPort)),
		r.CreatedAt,
	}
}

