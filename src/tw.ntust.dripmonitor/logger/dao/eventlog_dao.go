package dao

import (
	"database/sql"
	"tw.ntust.dripmonitor/logger/datamodels"
	log "github.com/sirupsen/logrus"
	"reflect"
	"time"
)

type EventLogDAO struct {
	db *sql.DB
}

func NewEventLogDAO(db *sql.DB) *EventLogDAO {
	return &EventLogDAO{
		db: db,
	}
}

func (d *EventLogDAO) InsertRecord(record *datamodels.EventLog) bool {
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

func (d *EventLogDAO) GetAdapterConnectsAfterRestart(adapterMAC string, recordCount int) *[]datamodels.EventLogSQL {
	errorResult := &[]datamodels.EventLogSQL{}

	// Find timestamp of the last boot_record
	var lastBootTime time.Time
	query := "select created_at from event_log where mac_adapter=? and event_code=52 order by created_at desc limit 1"
	err := d.db.QueryRow(query, adapterMAC).Scan(&lastBootTime)
	if err != nil {
		d.logError(err); return errorResult
	}

	// Query records related to bluetooth connects
	query = "select * from event_log where mac_adapter=? and event_code in (30, 31) " +
		"and created_at >= ? order by created_at desc limit ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		d.logError(err); return errorResult
	}
	defer stmt.Close()
	rows, err := stmt.Query(adapterMAC, lastBootTime, recordCount)
	if err != nil {
		d.logError(err); return errorResult
	}

	var results []datamodels.EventLogSQL
	for rows.Next() {
		record, err := d.scanAllColumns(rows)
		if err != nil {
			d.logError(err); return errorResult
		}
		results = append(results, *record)
	}

	return &results
}


func (d *EventLogDAO) scanAllColumns(rowOrRows interface{}) (*datamodels.EventLogSQL, error) {
	paramValue := reflect.ValueOf(rowOrRows)
	var r datamodels.EventLogSQL
	var err error

	// Fields of two scan operations are the same
	if paramValue.Type() == reflect.TypeOf(&sql.Rows{}) {
		err = rowOrRows.(*sql.Rows).Scan(&r.SN, &r.EventCode, &r.Message, &r.AdapterMAC, &r.DripMAC,
			&r.SrcIP, &r.SrcPort, &r.CreatedAt)
	} else {
		err = rowOrRows.(*sql.Row).Scan(&r.SN, &r.EventCode, &r.Message, &r.AdapterMAC, &r.DripMAC,
			&r.SrcIP, &r.SrcPort, &r.CreatedAt)
	}

	return &r, err
}

func (d *EventLogDAO) logError(err error) {
	log.Errorf("[EventLogDAO] Error occurred: " + err.Error())
}
