package datamodels

import (
	"time"
	"database/sql"
)

type EventRecord struct {
	SN            int64     `json:"sn"`
	EventCode     int       `json:"event_code" form:"event_code"`
	Message       string    `json:"message" form:"message"`
	AdapterMAC    string    `json:"mac_adapter" form:"mac_adapter"`
	DripMAC       string    `json:"mac_drip" form:"mac_drip"`
	SrcIP         string    `json:"src_ip"`
	SrcPort       int64     `json:"src_port"`
	CreatedAt     time.Time `json:"created_at"`
}

type EventRecordSQL struct {
	SN            int64
	EventCode     int
	Message       sql.NullString
	AdapterMAC    string
	DripMAC       sql.NullString
	SrcIP         sql.NullString
	SrcPort       sql.NullInt64
	CreatedAt     time.Time
}

func (r *EventRecord) SQLForm() *EventRecordSQL {
	rs := &EventRecordSQL{
		r.SN,
		r.EventCode,
		sql.NullString{r.Message, true},
		r.AdapterMAC,
		sql.NullString{r.DripMAC, true},
		sql.NullString{r.SrcIP, true},
		sql.NullInt64{r.SrcPort, true},
		r.CreatedAt,
	}

	// Handle null values
	if len(rs.Message.String) == 0 {
		rs.Message.Valid = false
	}
	if len(rs.DripMAC.String) == 0 {
		rs.DripMAC.Valid = false
	}
	if len(rs.SrcIP.String) == 0 {
		rs.SrcIP.Valid = false
	}
	if rs.SrcPort.Int64 == 0 {
		rs.SrcPort.Valid = false
	}

	return rs
}

// IsValid can do some very very simple "low-level" data validations.
//func (u User) IsValid() bool {
//	return u.ID > 0
//}


