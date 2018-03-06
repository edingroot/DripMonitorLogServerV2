package helpers

import (
	"runtime"
	"path"
	"database/sql"
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
