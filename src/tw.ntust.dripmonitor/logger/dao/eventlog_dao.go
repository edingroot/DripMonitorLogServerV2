package dao

import (
	"database/sql"
	"tw.ntust.dripmonitor/logger/datamodels"
	log "github.com/sirupsen/logrus"
)

const LogTagELD = "[EventLogDAO]"

type EventLogDAO struct {
	db *sql.DB
}

func NewEventLogDAO(db *sql.DB) *EventLogDAO {
	return &EventLogDAO{
		db: db,
	}
}

func (dao *EventLogDAO) InsertRecord(record *datamodels.EventRecord) bool {
	var err error
	var stmtIns *sql.Stmt

	query := "insert into event_log (event_code, message, mac_adapter, mac_drip, src_ip, src_port) values (?,?,?,?,?,?)"
	stmtIns, err = dao.db.Prepare(query)
	if err != nil {
		log.Errorf("%s Error preparing insert: %s", LogTagELD, err.Error())
		return false
	}
	defer stmtIns.Close()

	rs := record.SQLForm()
	_, err = stmtIns.Exec(rs.EventCode, rs.Message, rs.AdapterMAC, rs.DripMAC, rs.SrcIP, rs.SrcPort)
	if err != nil {
		log.Errorf("%s Error inserting record: %s", LogTagELD, err.Error())
		return false
	}

	return true
}
