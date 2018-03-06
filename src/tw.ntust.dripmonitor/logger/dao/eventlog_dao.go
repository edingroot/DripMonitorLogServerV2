package dao

import (
	"database/sql"
	"tw.ntust.dripmonitor/logger/datamodels"
	log "github.com/sirupsen/logrus"
)

type EventLogDAO struct {
	db *sql.DB
}

func NewEventLogDAO(db *sql.DB) *EventLogDAO {
	return &EventLogDAO{
		db: db,
	}
}

func (d *EventLogDAO) InsertRecord(record *datamodels.EventRecord) bool {
	errorResult := false

	query := "insert into event_log (event_code, message, mac_adapter, mac_drip, src_ip, src_port) values (?,?,?,?,?,?)"
	stmtIns, err := d.db.Prepare(query)
	if err != nil {
		d.logError(err); return errorResult
	}
	defer stmtIns.Close()

	rs := record.SQLForm()
	_, err = stmtIns.Exec(rs.EventCode, rs.Message, rs.AdapterMAC, rs.DripMAC, rs.SrcIP, rs.SrcPort)
	if err != nil {
		d.logError(err); return errorResult
	}

	return true
}


func (d *EventLogDAO) logError(err error) {
	log.Errorf("[EventLogDAO] Error occurred: " + err.Error())
}
