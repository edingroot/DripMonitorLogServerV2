package helpers

import (
	"runtime"
	"path"
	"database/sql"
	"strings"
	"strconv"
)

const TimeFormat = "2006-01-02 15:04:05"

var ProjectPath string

func InitializePaths() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("InitializeCommon: No caller information")
	} else {
		ProjectPath = path.Join(path.Dir(filename), "../")
	}
}

func GetIpPortFromAddr(remoteAddr string) (string, int32) {
	s := strings.Split(remoteAddr, ":")
	srcIp := s[0]
	srcPort, _ := strconv.Atoi(s[1])
	return srcIp, int32(srcPort)
}

func StringToNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{"", false}
	} else {
		return sql.NullString{s, true}
	}
}

func Int64ToNullInt64(n int64) sql.NullInt64 {
	if n == 0 {
		return sql.NullInt64{0, false}
	} else {
		return sql.NullInt64{n, true}
	}
}
