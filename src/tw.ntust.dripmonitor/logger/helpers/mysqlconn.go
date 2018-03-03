package helpers

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/kataras/iris/core/errors"
	"fmt"
)

// Instances should be created by calling NewMySQLConn(*Configuration)
type MySQLConn struct {
	DB *sql.DB
}

func NewMySQLConn(config *Configuration) (*MySQLConn, error) {
	var conn MySQLConn
	var err error

	conn.DB, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@%s/%s", config.DbUsername, config.DbPassword, config.DbHost, config.DbName))
	if err != nil {
		return nil, errors.New("Filed to open MySQL connection")
	}

	return &conn, nil
}

func (conn *MySQLConn) Close() {
	conn.DB.Close()
}
