package dao

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

const LogTagSLD = "[StreamLogDAO]"

type StreamLogDAO struct {
	db *sql.DB
}

func NewStreamLogDAO(db *sql.DB) *StreamLogDAO {
	return &StreamLogDAO{
		db: db,
	}
}

func (dao *StreamLogDAO) InsertRecord(message string, srcIp string, srcPort int) bool {
	var err error
	var stmtIns *sql.Stmt

	query := "insert into tcp_log_stream (message, src_ip, src_port) values (?,?,?)"
	stmtIns, err = dao.db.Prepare(query)
	if err != nil {
		log.Errorf("%s Error preparing insert: %s", LogTagSLD, err.Error())
		return false
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(message, srcIp, srcPort)
	if err != nil {
		log.Errorf("%s Error inserting record: %s", LogTagSLD, err.Error())
		return false
	}

	return true
}
